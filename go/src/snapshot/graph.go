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
 * 	Check whether the graph is connected
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
				j := myNeigh[n].to
				fmt.Print(" ", j)
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
	}

	visited := make(map[ProcIdx]bool)

	stack := make([]ProcIdx, nProc)
	stackSize := 0

	// push v
	v := ProcIdx(1)
	stack[stackSize] = 1
	stackSize++
	visited[v] = true

	for stackSize > 0 {
		stackSize--
		v = stack[stackSize]

		if dbg {
			fmt.Print("Popping ", v, " sz ", stackSize, ": ")
		}
		if !visited[v] {
			s := fmt.Sprintf("ERROR: Popped non-visited proc %d from stack\n", v)
			panic(s)
		}

		myNeigh := graph.Neighbors[v]
		numNeigh := len(myNeigh)
		if dbg {
			fmt.Print("  Pushing out of ", numNeigh, " neighbors:")
		}
		for n := 0; n < numNeigh; n++ {
			neigh := myNeigh[n].to
			if visited[neigh] {
				continue
			}
			if dbg {
				fmt.Print(" [", stackSize, "]", neigh)
			}

			stack[stackSize] = neigh
			stackSize++
			visited[neigh] = true
		}
		if dbg {
			fmt.Println("(", stackSize, ")")
		}
	}
	return len(visited) == int(nProc)
}

/*************************************************************
 * Make one horizontal (real graph) connection
*************************************************************/
func (graph *Graph) makeOneHorizChan(i, j ProcIdx) {
	ijChan := make(HorizBidirChan, HorizChanCap)
	jiChan := make(HorizBidirChan, HorizChanCap)

	var iIoChan HorizChanPair
	iIoChan.in = HorizBidir2InChan(jiChan)
	iIoChan.out = HorizBidir2OutChan(ijChan)
	iIoChan.from = i
	iIoChan.to = j

	var jIoChan HorizChanPair
	jIoChan.in = HorizBidir2InChan(ijChan)
	jIoChan.out = HorizBidir2OutChan(jiChan)
	jIoChan.from = j
	jIoChan.to = i

	graph.Neighbors[i] = append(graph.Neighbors[i], iIoChan)
	graph.Neighbors[j] = append(graph.Neighbors[j], jIoChan)
}

/*************************************************************
 * For the graph that is not fully connected, add connections
 * (i,i+1) that do not already exist
*************************************************************/
func (graph *Graph) addConnectionsToDisconnected() {
	var numAdded int = 0
	nProc := graph.NumberProcs
	for p := ProcIdx(0); p < nProc-1; p++ {
		myNeigh := (graph.Neighbors)[p]
		numNeigh := len(myNeigh)

		found := false
		for n := 0; n < numNeigh; n++ {
			neigh := myNeigh[n].to
			if neigh == p+1 {
				found = true
				break
			}
		}
		if !found {
			graph.makeOneHorizChan(p, p+1)
			numAdded++
		}
	}
	if numAdded > 0 {
		fmt.Println("Added", numAdded, "new chans p->p+1")
	}
}

/*************************************************************
 * Make real graph channels (not top down)
 * percChans is probability (as percentage) of channel being
 * present.
 * See: Erdős–Rényi model at
 *      https://en.wikipedia.org/wiki/Erdős–Rényi_model
*************************************************************/
func (graph *Graph) makeNeighborChans(percChans int) {
	nProc := graph.NumberProcs
	numChans := 0

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
			numChans++
		}
	}
	fmt.Println("First created", numChans, "channels")
	if !graph.verifyConnectivity() {
		graph.addConnectionsToDisconnected()
	} else {
		fmt.Println("Added zero channels")
	}
}

/*************************************************************
 * Make channels to communicate values to individual Procs
*************************************************************/
func (graph *Graph) makeVertChans() {
	nProc := graph.NumberProcs
	graph.tops = make([]VertChanPair, nProc)
	graph.driverTops = make([]VertChanPair, nProc)

	for i := ProcIdx(0); i < nProc; i++ {
		topDown := make(VertBidirChan, VertChanCap)
		botUp := make(VertBidirChan, VertChanCap)

		graph.tops[i].in = VertBidir2InChan(topDown)
		graph.tops[i].out = VertBidir2OutChan(botUp)

		graph.driverTops[i].in = VertBidir2InChan(botUp)
		graph.driverTops[i].out = VertBidir2OutChan(topDown)
	}
}

/*************************************************************
 * Start proc goroutines
*************************************************************/
func (graph *Graph) startProcs() {
	nProcs := graph.NumberProcs
	graph.Procs = make([]Proc, nProcs)
	for i := ProcIdx(0); i < nProcs; i++ {
		graph.Procs[i].id = i
	}
	for i := ProcIdx(0); i < nProcs; i++ {
		go graph.Procs[i].Run(&graph.tops[i], graph.Neighbors[i])
	}
}

/*************************************************************
 * Send data to each proc. Root value is negative.
 * Calculate the sum of all positive values for comparison.
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
		downChanOut := graph.driverTops[i].out
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
			_, ok := <-graph.driverTops[i].in
			if !ok {
				s := fmt.Sprintf("ERROR: Graph - bad receive from non-root %d", i)
				panic(s)
			}
		}
	}
}

/*************************************************************
 * Receive data from root (that is sum of all data sent in
 * sendDataDown(), except positive root data is used
*************************************************************/
func (graph *Graph) receiveFromRoot() Data {
	fromRoot := graph.driverTops[graph.Root].in
	val, ok := <-fromRoot
	if !ok {
		s := fmt.Sprintf("ERROR: Graph - bad receive from root %d", graph.Root)
		panic(s)
	}
	fmt.Println("Root returns: ", val)
	return val
}

/*************************************************************
 * Build horizontal and vertical channels
*************************************************************/
func (graph *Graph) buildChans(nProc ProcIdx, root ProcIdx, percChans int) {

	graph.Root = root
	graph.NumberProcs = nProc
	graph.makeNeighborChans(percChans)
	graph.makeVertChans()
}

/*************************************************************
 * Build chans and procs
*************************************************************/
func (graph *Graph) buildGraph(nProc ProcIdx, root ProcIdx, percChans int) {
	graph.buildChans(nProc, root, percChans)
	graph.startProcs()
}

/*************************************************************
*************************************************************/
func (graph *Graph) BuildAndCollectData(nProc ProcIdx, root ProcIdx, percChans int, bias int) {
	graph.buildGraph(nProc, root, percChans)

	localSum := graph.sendDataDown(bias)

	// receive value from root first
	val := graph.receiveFromRoot()
	if val != localSum {
		fmt.Println("ERROR: Local sum (", localSum, ") != received sum (", val, ")")
	}

	graph.receiveFromNonRoots()
}
