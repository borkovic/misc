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

/*************************************************************
*************************************************************/
func makeNeighborChans(nProc int) [][]snapshot.HorizChanPair {
	const (
		HorizChanCap = 1
	)

	r0 := time.Now().UnixNano()
	RNG := rand.New(rand.NewSource(r0))

	// make neighbor channels
	neighbors := make([][]snapshot.HorizChanPair, nProc)
	for i := 0; i < nProc-1; i++ {
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

			var jIoChan snapshot.HorizChanPair
			jIoChan.In = snapshot.HorizBidir2InChan(ijChan)
			jIoChan.Out = snapshot.HorizBidir2OutChan(jiChan)

			neighbors[i] = append(neighbors[i], iIoChan)
			neighbors[j] = append(neighbors[j], jIoChan)
		}
	}
	return neighbors
}

/*************************************************************
*************************************************************/
func startProcs(procs []snapshot.Proc,
	tops []snapshot.VertChanPair,
	neighbors [][]snapshot.HorizChanPair) {

	nProc := len(procs)
	for i := 0; i < nProc; i++ {
		go procs[i].Run(&tops[i], neighbors[i])
	}
}


/*************************************************************
*************************************************************/
func main() {
	const (
		VertChanCap  = 0
	)

	r0 := time.Now().UnixNano()
	RNG := rand.New(rand.NewSource(r0))

	bias := RNG.Intn(5)
	nProc := 100 + RNG.Intn(20)
	root := RNG.Intn(nProc)
	fmt.Println("Num proc ", nProc, ", Bias is ", bias, ", root is ", root)

	neighbors := makeNeighborChans(nProc)

	procs := make([]snapshot.Proc, nProc)
	tops := make([]snapshot.VertChanPair, nProc)
	driverTops := make([]snapshot.VertChanPair, nProc)


	// make vert channels, start goroutines and send data down
	var localSum snapshot.Data = 0

	for i := 0; i < nProc; i++ {
		{
			topDown := make(snapshot.VertBidirChan, VertChanCap)
			botUp := make(snapshot.VertBidirChan, VertChanCap)

			tops[i].In = snapshot.VertBidir2InChan(topDown)
			tops[i].Out = snapshot.VertBidir2OutChan(botUp)

			driverTops[i].In = snapshot.VertBidir2InChan(botUp)
			driverTops[i].Out = snapshot.VertBidir2OutChan(topDown)
		}
	}

	startProcs(procs, tops, neighbors)

	for i := 0; i < nProc; i++ {
		var v, sendv snapshot.Data
		if i != root {
			v = snapshot.Data(i + bias + 10)
			sendv = v
		} else {
			v = snapshot.Data(i + bias + 1000)
			sendv = -v
		}
		localSum += v
		downChanOut := driverTops[i].Out
		downChanOut <- sendv
		close(downChanOut)
	}
	fmt.Println("Local sum: ", localSum)

	// receive value from root first
	val, ok := <-driverTops[root].In
	if !ok {
		panic("Bad receive 1")
	}
	fmt.Println("Root returns: ", val)
	if val != localSum {
		fmt.Println("Local sum (", localSum, ") != received sum (", val, ")")
	}

	// receive on remaining vert channel
	for i := 0; i < nProc; i++ {
		if i != root {
			val, ok = <-driverTops[i].In
			if !ok {
				panic("Bad receive 2")
			}
		}
	}
}
