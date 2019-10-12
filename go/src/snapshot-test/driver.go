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
	)

	r0 := time.Now().UnixNano()
	RNG := rand.New(rand.NewSource(r0))

	nProc := NumProc
	procs := make([]snapshot.Proc, nProc)
	tops := make([]snapshot.VertChanPair, nProc)

	neighbors := make([][]snapshot.HorizChanPair, nProc)
	for i := 0; i < nProc-1; i++ {
		for j := i + 1; j < nProc; j++ {
			if i+1 < j { // connect i at least with j==i+1 to have fully connected graph
				var n int = RNG.Intn(11)
				var b bool = n < 3 // with 30% probability do not have a connection
				if b {
					continue
				}
			}
			ijChan := make(snapshot.HorizBidirChan, 1)
			jiChan := make(snapshot.HorizBidirChan, 1)

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

	var rootBotUpIn snapshot.VertInChan
	var localSum int = 0
	var sum int = 0

	for i := 0; i < nProc; i++ {
		var topDownOut snapshot.VertOutChan
		var botUpIn snapshot.VertInChan
		{
			topDown := make(snapshot.VertBidirChan, 1)
			topDownOut = snapshot.VertBidir2OutChan(topDown)
			tops[i].In = snapshot.VertBidir2InChan(topDown)
			botUp := make(snapshot.VertBidirChan, 1)
			botUpIn = snapshot.VertBidir2InChan(botUp)
			tops[i].Out = snapshot.VertBidir2OutChan(botUp)
		}

		go procs[i].Run(&tops[i], neighbors[i])

		if i != root {
			v := i + 10
			localSum += v
			topDownOut <- snapshot.Data(v)
			sum += v
		} else {
			rootBotUpIn = botUpIn
			v := i + 100
			localSum += v
			topDownOut <- snapshot.Data(-v)
			sum += v
		}
	}
	fmt.Println("Local sum ", localSum)
	val := <-rootBotUpIn
	fmt.Println("Root returns ", val)
}
