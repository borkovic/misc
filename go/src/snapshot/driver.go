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
func verifyConnectivity(neighbors [][]HorizChanPair) bool {
	//dbg := true
	dbg := false
	nProc := ProcIdx(len(neighbors))
	if dbg {
		for i := ProcIdx(0); i < nProc; i++ {
			fmt.Print(i, " : ")
			myNeigh := neighbors[i]
			numNeigh := len(myNeigh)
			for n := 0; n < numNeigh; n++ {
				j := myNeigh[n].To
				fmt.Print(" ", j)
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
	}

	visited := make(map[ProcIdx] bool)

	stack := make([]ProcIdx, nProc)
	stackSize := 0

	// push v
	v := ProcIdx(1)
	stack[stackSize] = 1;
	stackSize++
	visited[v] = true

	for stackSize > 0 {
		stackSize--;
		v = stack[stackSize]

		if dbg { fmt.Print("Popping ", v, " sz ", stackSize, ": ") }
		if ! visited[v] {
			fmt.Println("Not Visited proc ", v, "popped from stack")
			panic("Panic")
		}

		myNeigh := neighbors[v]
		numNeigh := len(myNeigh)
		if dbg { fmt.Print("  Pushing out of ", numNeigh, " neighbors:") }
		for n := 0; n < numNeigh; n++ {
			neigh := myNeigh[n].To
			if visited[neigh] {
				continue
			}
			if dbg { fmt.Print(" [", stackSize, "]", neigh) }

			stack[stackSize] = neigh
			stackSize++
			visited[neigh] = true
		}
		if dbg { fmt.Println("(", stackSize, ")") }
	}
	return len(visited) == int(nProc)
}

/*************************************************************
*************************************************************/
func makeOneHorizChan(neighbors *[][]HorizChanPair, i, j ProcIdx) {

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

	(*neighbors)[i] = append((*neighbors)[i], iIoChan)
	(*neighbors)[j] = append((*neighbors)[j], jIoChan)
}

/*************************************************************
 * Add connections i->i+1 if it does not exist
*************************************************************/
func addConnections(neighbors *[][]HorizChanPair) {
	var numAdded int = 0
	nProc := ProcIdx(len(*neighbors))
	for p := ProcIdx(0); p < nProc-1; p++ {
		myNeigh := (*neighbors)[p]
		numNeigh := len(myNeigh)
		found := false
		for n := 0; n < numNeigh; n++ {
			neigh := myNeigh[n].To
			if neigh == p+1 {
				found = true
				break
			}
		}
		if ! found {
			makeOneHorizChan(neighbors, p, p+1)
			numAdded++
		}
	}
	if numAdded > 0 {
		fmt.Println("Added", numAdded, "chans p->p+1")
	}
}

/*************************************************************
*************************************************************/
func makeNeighborChans(nProc ProcIdx, percChans int) [][]HorizChanPair {

	r0 := time.Now().UnixNano()
	RNG := rand.New(rand.NewSource(r0))

	neighbors := make([][]HorizChanPair, nProc)

	for i := ProcIdx(0); i < nProc-1; i++ {
		for j := i + 1; j < nProc; j++ {
			n := RNG.Intn(100)
			b := n < (100 - percChans)
			if b {
				continue
			}
			makeOneHorizChan(&neighbors, i, j)
		}
	}
	if ! verifyConnectivity(neighbors) {
		fmt.Println("Not connected, connecting p->p+1")
		addConnections(&neighbors)
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
func Driver(nProc ProcIdx, root ProcIdx, percChans int, bias int) {

	neighbors := makeNeighborChans(nProc, percChans)

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
