package main

import (
	"fmt"
	"math/rand"
	"time"
)
import (
	"snapshot"
)

const (
	Debug bool = snapshot.Debug
)
const (
	VertChanCap  = 0
	HorizChanCap = 1
)

type ProcIdx = snapshot.ProcIdx

/*************************************************************
*************************************************************/
func makeNeighborChans(nProc ProcIdx) [][]snapshot.HorizChanPair {

	r0 := time.Now().UnixNano()
	RNG := rand.New(rand.NewSource(r0))

	neighbors := make([][]snapshot.HorizChanPair, nProc)
	for i := ProcIdx(0); i < nProc-1; i++ {
		for j := i + 1; j < nProc; j++ {
			if i+1 < j { // always connect (i -> j) with j==i+1 for full connectedness
				n := RNG.Intn(11)
				b := n < 6 // with 60% probability do not have a connection
				if b {
					continue
				}
			}
			ijChan := make(snapshot.HorizBidirChan, HorizChanCap)
			jiChan := make(snapshot.HorizBidirChan, HorizChanCap)

			var iIoChan snapshot.HorizChanPair
			iIoChan.In = snapshot.HorizBidir2InChan(jiChan)
			iIoChan.Out = snapshot.HorizBidir2OutChan(ijChan)
			iIoChan.From = i
			iIoChan.To = j

			var jIoChan snapshot.HorizChanPair
			jIoChan.In = snapshot.HorizBidir2InChan(ijChan)
			jIoChan.Out = snapshot.HorizBidir2OutChan(jiChan)
			jIoChan.From = j
			jIoChan.To = i

			neighbors[i] = append(neighbors[i], iIoChan)
			neighbors[j] = append(neighbors[j], jIoChan)
		}
	}
	return neighbors
}

/*************************************************************
*************************************************************/
func makeVertChans(
	tops []snapshot.VertChanPair,
	driverTops []snapshot.VertChanPair) {

	nProc := ProcIdx(len(tops))

	for i := ProcIdx(0); i < nProc; i++ {
		topDown := make(snapshot.VertBidirChan, VertChanCap)
		botUp := make(snapshot.VertBidirChan, VertChanCap)

		tops[i].In = snapshot.VertBidir2InChan(topDown)
		tops[i].Out = snapshot.VertBidir2OutChan(botUp)

		driverTops[i].In = snapshot.VertBidir2InChan(botUp)
		driverTops[i].Out = snapshot.VertBidir2OutChan(topDown)
	}
}

/*************************************************************
*************************************************************/
func startProcs(procs []snapshot.Proc,
	tops []snapshot.VertChanPair,
	neighbors [][]snapshot.HorizChanPair) {

	nProc := ProcIdx(len(procs))
	for i := ProcIdx(0); i < nProc; i++ {
		go procs[i].Run(&tops[i], neighbors[i])
	}
}

/*************************************************************
*************************************************************/
func sendDataDown(
	driverTops []snapshot.VertChanPair,
	root ProcIdx,
	bias int) snapshot.Data {

	nProc := ProcIdx(len(driverTops))
	var localSum snapshot.Data = 0
	for i := ProcIdx(0); i < nProc; i++ {
		var sendv snapshot.Data
		v := snapshot.Data(int(i) + bias)
		if i != root {
			v += 10
			sendv = v
		} else {
			v += 1000
			sendv = -v
		}
		localSum += v
		downChanOut := driverTops[i].Out
		downChanOut <- sendv
		close(downChanOut)
	}
	fmt.Println("Local sum: ", localSum)
	return localSum
}

/*************************************************************
*************************************************************/
func receiveFromNonRoots(
	driverTops []snapshot.VertChanPair,
	root ProcIdx) {

	nProc := ProcIdx(len(driverTops))
	for i := ProcIdx(0); i < nProc; i++ {
		if i != root {
			_, ok := <-driverTops[i].In
			if !ok {
				panic("Bad receive 2")
			}
		}
	}
}

/*************************************************************
*************************************************************/
func receiveFromRoot(fromRoot snapshot.VertInChan) snapshot.Data {
	val, ok := <-fromRoot
	if !ok {
		panic("Bad receive 1")
	}
	fmt.Println("Root returns: ", val)
	return val
}

/*************************************************************
*************************************************************/
func Driver(nProc ProcIdx, root ProcIdx, bias int) {

	neighbors := makeNeighborChans(nProc)

	tops := make([]snapshot.VertChanPair, nProc)
	driverTops := make([]snapshot.VertChanPair, nProc)

	makeVertChans(tops, driverTops)

	procs := make([]snapshot.Proc, nProc)
	startProcs(procs, tops, neighbors)

	localSum := sendDataDown(driverTops, root, bias)

	// receive value from root first
	val := receiveFromRoot(driverTops[root].In)
	if val != localSum {
		fmt.Println("Local sum (", localSum, ") != received sum (", val, ")")
	}

	receiveFromNonRoots(driverTops, root)
}

/*************************************************************
*************************************************************/
func main() {

	r0 := time.Now().UnixNano()
	RNG := rand.New(rand.NewSource(r0))

	bias := RNG.Intn(5)
	nProc := ProcIdx(100 + RNG.Intn(20))
	root := ProcIdx(RNG.Intn(int(nProc)))
	fmt.Println("Num proc ", nProc, ", Bias is ", bias, ", root is ", root)

	Driver(nProc, root, bias)
}
