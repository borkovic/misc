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
func (graph *Graph) verifyConnectivity() bool {
	//dbg := true
	dbg := false
	nProc := ProcIdx(len(graph.Neighbors))
	if dbg {
		for i := ProcIdx(0); i < nProc; i++ {
			fmt.Print(i, " : ")
			myNeigh := graph.Neighbors[i]
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

		myNeigh := graph.Neighbors[v]
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
func (graph *Graph) makeOneHorizChan(i, j ProcIdx) {

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

	graph.Neighbors[i] = append(graph.Neighbors[i], iIoChan)
	graph.Neighbors[j] = append(graph.Neighbors[j], jIoChan)
}

/*************************************************************
 * Add connections i->i+1 if it does not exist
*************************************************************/
func (graph *Graph) addConnections() {
	var numAdded int = 0
	nProc := ProcIdx(len(graph.Neighbors))
	for p := ProcIdx(0); p < nProc-1; p++ {
		myNeigh := (graph.Neighbors)[p]
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
			graph.makeOneHorizChan(p, p+1)
			numAdded++
		}
	}
	if numAdded > 0 {
		fmt.Println("Added", numAdded, "chans p->p+1")
	}
}

/*************************************************************
*************************************************************/
func (graph *Graph) makeNeighborChans(percChans int) {
	nProc := graph.NumberProcs

	r0 := time.Now().UnixNano()
	RNG := rand.New(rand.NewSource(r0))

	graph.Neighbors = make([][]HorizChanPair, nProc)
	percNoChan := 100 - percChans

	for i := ProcIdx(0); i < nProc-1; i++ {
		for j := i + 1; j < nProc; j++ {
			n := RNG.Intn(100)
			if n < percNoChan {
				continue
			}
			graph.makeOneHorizChan(i, j)
		}
	}
	if ! graph.verifyConnectivity() {
		graph.addConnections()
	}
}

/*************************************************************
*************************************************************/
func (graph *Graph) makeVertChans() {
	nProc := graph.NumberProcs
	graph.Tops = make([]VertChanPair, nProc)
	graph.DriverTops = make([]VertChanPair, nProc)

	for i := ProcIdx(0); i < nProc; i++ {
		topDown := make(VertBidirChan, VertChanCap)
		botUp := make(VertBidirChan, VertChanCap)

		graph.Tops[i].In = VertBidir2InChan(topDown)
		graph.Tops[i].Out = VertBidir2OutChan(botUp)

		graph.DriverTops[i].In = VertBidir2InChan(botUp)
		graph.DriverTops[i].Out = VertBidir2OutChan(topDown)
	}
}

/*************************************************************
*************************************************************/
func (graph *Graph) startProcs() {
	nProcs := graph.NumberProcs
	graph.Procs = make([]Proc, nProcs)
	for i := ProcIdx(0); i < nProcs; i++ {
		go graph.Procs[i].Run(&graph.Tops[i], graph.Neighbors[i])
	}
}

/*************************************************************
*************************************************************/
func (graph *Graph) sendDataDown(
		bias int) Data {

	nProc := graph.NumberProcs
	root := graph.Root
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
		downChanOut := graph.DriverTops[i].Out
		downChanOut <- sendv
		close(downChanOut)
	}
	fmt.Println("Local sum: ", localSum)
	return localSum
}

/*************************************************************
*************************************************************/
func (graph *Graph) receiveFromNonRoots() {
	nProc := graph.NumberProcs
	root := graph.Root
	for i := ProcIdx(0); i < nProc; i++ {
		if i != root {
			_, ok := <-graph.DriverTops[i].In
			if !ok {
				panic("Bad receive 2")
			}
		}
	}
}

/*************************************************************
*************************************************************/
func (graph *Graph) receiveFromRoot() Data {
	fromRoot := graph.DriverTops[graph.Root].In
	val, ok := <-fromRoot
	if !ok {
		panic("Bad receive 1")
	}
	fmt.Println("Root returns: ", val)
	return val
}

/*************************************************************
*************************************************************/
func (graph *Graph) buildGraph(nProc ProcIdx, root ProcIdx, percChans int) {

	graph.Root = root
	graph.NumberProcs = nProc
	graph.makeNeighborChans(percChans)
	graph.makeVertChans()
}

/*************************************************************
*************************************************************/
func (graph *Graph) Driver(nProc ProcIdx, root ProcIdx, percChans int, bias int) {
	graph.buildGraph(nProc, root, percChans)
	graph.startProcs()
	localSum := graph.sendDataDown(bias)

	// receive value from root first
	val := graph.receiveFromRoot()
	if val != localSum {
		fmt.Println("Local sum (", localSum, ") != received sum (", val, ")")
	}

	graph.receiveFromNonRoots()
}
