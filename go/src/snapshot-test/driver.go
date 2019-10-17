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
func main() {
	const (
		NumNeighbors = 6
		NumProc      = 100
		root         = 2
		VertChanCap  = 0
		HorizChanCap = 1
	)

	r0 := time.Now().UnixNano()
	RNG := rand.New(rand.NewSource(r0))

	nProc := NumProc
	procs := make([]snapshot.Proc, nProc)
	tops := make([]snapshot.VertChanPair, nProc)
	driverTops := make([]snapshot.VertChanPair, nProc)

	// make neighbor channels
	neighbors := make([][]snapshot.HorizChanPair, nProc)
	for i := 0; i < nProc-1; i++ {
		for j := i + 1; j < nProc; j++ {
			if i+1 < j { // connect i at least with j==i+1 to have fully connected graph
				n := RNG.Intn(11)
				b := n < 3 // with 30% probability do not have a connection
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

	// make vert channels, start goroutines and send data down
	var rootBotUpIn snapshot.VertInChan
	var localSum snapshot.Data = 0

	for i := 0; i < nProc; i++ {
		var topDownOut snapshot.VertOutChan
		var botUpIn snapshot.VertInChan
		{
			topDown := make(snapshot.VertBidirChan, VertChanCap)
			botUp := make(snapshot.VertBidirChan, VertChanCap)

			topDownOut = snapshot.VertBidir2OutChan(topDown)
			botUpIn = snapshot.VertBidir2InChan(botUp)

			tops[i].In = snapshot.VertBidir2InChan(topDown)
			tops[i].Out = snapshot.VertBidir2OutChan(botUp)

			driverTops[i].In = snapshot.VertBidir2InChan(botUp)
			driverTops[i].Out = snapshot.VertBidir2OutChan(topDown)
		}

		go procs[i].Run(&tops[i], neighbors[i])

		if i != root {
			v := snapshot.Data(i + 10)
			localSum += v
			topDownOut <- v
			close(topDownOut)
		} else {
			rootBotUpIn = botUpIn
			v := snapshot.Data(i + 1000)
			localSum += v
			topDownOut <- (-v)
			close(topDownOut)
		}
	}
	fmt.Println("Local sum: ", localSum)

	// receive value from root
	val, ok := <-rootBotUpIn
	if !ok {
		panic("Bad receive 1")
	}
	fmt.Println("Root returns: ", val)
	if val != localSum {
		fmt.Println("Local sum (", localSum, ") != received sum (", val, ")")
	}
	for i := 0; i < nProc; i++ {
		if i != root {
			val, ok = <-driverTops[i].In
			if !ok {
				panic("Bad receive 2")
			}
		}
	}
}
