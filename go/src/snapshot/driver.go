package snapshot

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	VertChanCap  = 0
	HorizChanCap = 1
)


/*************************************************************
*************************************************************/
func makeNeighborChans(nProc ProcIdx) [][]HorizChanPair {

	r0 := time.Now().UnixNano()
	RNG := rand.New(rand.NewSource(r0))

	neighbors := make([][]HorizChanPair, nProc)
	for i := ProcIdx(0); i < nProc-1; i++ {
		for j := i + 1; j < nProc; j++ {
			if i+1 < j { // always connect (i -> j) with j==i+1 for full connectedness
				n := RNG.Intn(11)
				b := n < 6 // with 60% probability do not have a connection
				if b {
					continue
				}
			}
			ijChan := make(HorizBidirChan, HorizChanCap)
			jiChan := make(HorizBidirChan, HorizChanCap)

			var iIoChan HorizChanPair
			iIoChan.In = HorizBidir2InChan(jiChan)
			iIoChan.Out = HorizBidir2OutChan(ijChan)
			iIoChan.From = i
			iIoChan.To = j

			var jIoChan HorizChanPair
			jIoChan.In = HorizBidir2InChan(ijChan)
			jIoChan.Out = HorizBidir2OutChan(jiChan)
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
	tops []VertChanPair,
	driverTops []VertChanPair) {

	nProc := ProcIdx(len(tops))

	for i := ProcIdx(0); i < nProc; i++ {
		topDown := make(VertBidirChan, VertChanCap)
		botUp := make(VertBidirChan, VertChanCap)

		tops[i].In = VertBidir2InChan(topDown)
		tops[i].Out = VertBidir2OutChan(botUp)

		driverTops[i].In = VertBidir2InChan(botUp)
		driverTops[i].Out = VertBidir2OutChan(topDown)
	}
}

/*************************************************************
*************************************************************/
func startProcs(procs []Proc,
	tops []VertChanPair,
	neighbors [][]HorizChanPair) {

	nProc := ProcIdx(len(procs))
	for i := ProcIdx(0); i < nProc; i++ {
		go procs[i].Run(&tops[i], neighbors[i])
	}
}

/*************************************************************
*************************************************************/
func sendDataDown(
	driverTops []VertChanPair,
	root ProcIdx,
	bias int) Data {

	nProc := ProcIdx(len(driverTops))
	var localSum Data = 0
	for i := ProcIdx(0); i < nProc; i++ {
		var sendv Data
		v := Data(int(i) + bias)
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
	driverTops []VertChanPair,
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
func receiveFromRoot(fromRoot VertInChan) Data {
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

	tops := make([]VertChanPair, nProc)
	driverTops := make([]VertChanPair, nProc)

	makeVertChans(tops, driverTops)

	procs := make([]Proc, nProc)
	startProcs(procs, tops, neighbors)

	localSum := sendDataDown(driverTops, root, bias)

	// receive value from root first
	val := receiveFromRoot(driverTops[root].In)
	if val != localSum {
		fmt.Println("Local sum (", localSum, ") != received sum (", val, ")")
	}

	receiveFromNonRoots(driverTops, root)
}

