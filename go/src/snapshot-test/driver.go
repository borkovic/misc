package main

import (
	"fmt"
)
import (
	"snapshot"
)

/*************************************************************
*************************************************************/
func main() {
	root := 2
	nProc := snapshot.NumProc
	procs := make([]snapshot.Proc, nProc)
	tops := make([]snapshot.VertChanPair, nProc)

	neighbors := make([][]snapshot.HorizChanPair, nProc)
	for i := 0; i < nProc-1; i++ {
		for j := i + 1; j < nProc; j++ {
			ijChan := make(snapshot.HorizBidirChan, 1)
			jiChan := make(snapshot.HorizBidirChan, 1)

			var iIoChan snapshot.HorizChanPair
			iIoChan.In = snapshot.HorizInChan(jiChan)
			iIoChan.Out = snapshot.HorizOutChan(ijChan)

			var jIoChan snapshot.HorizChanPair
			jIoChan.In = snapshot.HorizInChan(ijChan)
			jIoChan.Out = snapshot.HorizOutChan(jiChan)

			neighbors[i] = append(neighbors[i], iIoChan)
			neighbors[j] = append(neighbors[j], jIoChan)
		}
	}

	var rootBotUp snapshot.VertBidirChan
	var localSum int = 0
	var sum int = 0

	for i := 0; i < nProc; i++ {
		topDown := make(snapshot.VertBidirChan, 1)
		tops[i].In = topDown
		botUp := make(snapshot.VertBidirChan, 1)
		tops[i].Out = botUp

		go procs[i].Run(&tops[i], neighbors[i])

		if i != root {
			v := i + 5
			localSum += v
			topDown <- snapshot.Data(v)
			sum += v
		} else {
			rootBotUp = botUp
			v := i + 25
			localSum += v
			topDown <- snapshot.Data(-v)
			sum += v
		}
	}
	fmt.Println("Local sum ", localSum)
	val := <-rootBotUp
	fmt.Println("Root returns ", val)
}
