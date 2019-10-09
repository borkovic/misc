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
	tops := make([]snapshot.IoChans, nProc)

	neighbors := make([][]snapshot.IoChans, nProc)
	for i := 0; i < nProc-1; i++ {
		for j := i + 1; j < nProc; j++ {
			ijChan := make(snapshot.BiChan, 1)
			jiChan := make(snapshot.BiChan, 1)
			iIoChan := snapshot.IoChans{Out: ijChan, In: jiChan}
			jIoChan := snapshot.IoChans{Out: jiChan, In: ijChan}
			neighbors[i] = append(neighbors[i], iIoChan)
			neighbors[j] = append(neighbors[j], jIoChan)
		}
	}
	var rootBotUp snapshot.IChan
	var sum int = 0
	for i := 0; i < nProc; i++ {
		topDown := make(snapshot.BiChan, 1)
		tops[i].In = topDown
		botUp := make(snapshot.BiChan, 1)
		tops[i].Out = botUp

		go procs[i].Run(&tops[i], neighbors[i])

		if i != root {
			v := i + 5
			topDown <- snapshot.Data(v)
			sum += v
		} else {
			rootBotUp = botUp
			v := i + 4
			topDown <- snapshot.Data(-v)
			sum += v
		}
	}
	val := <-rootBotUp
	fmt.Println("Root returns ", val)
}
