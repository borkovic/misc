package snapshot

import (
	"fmt"
	"reflect"
)

/*************************************************************
 * Child
 * 1. receive value from parent
 * 2. send its neg. value to others (children and siblings)
 * 3. receive values from children(positive) and siblings(negative)
 * 4. sums values from children (positive)`
 * 5. send sum to parent
*************************************************************/
func (proc *Proc) runChild(neighbors NeighborChans) {
	//numNeighbors := len(neighbors)
	var parIdx int = -1
	var parVal Data

	proc.SLEEP(4)
	// 1. read from parent In tree
	select {
	case parVal = <-neighbors[0].In:
		parIdx = 0
	case parVal = <-neighbors[1].In:
		parIdx = 1
	case parVal = <-neighbors[2].In:
		parIdx = 2
	}
	if Debug {
		fmt.Println("Child with val ", proc.m_MyVal, ", par idx ", parIdx,
			", par value ", parVal)
	}

	// 2. send to all but parent (children and siblings)
	for i, n := range neighbors {
		if i == parIdx {
			proc.SLEEP(4)
			continue
		}
		proc.SLEEP(4)
		n.Out <- (-proc.m_MyVal)
	}

	// 3. receive from children and siblings
	// from children positive, from siblings negative
	// 4. sum from children
	sum := (proc.m_MyVal)
	for i, n := range neighbors {
		if i == parIdx {
			continue
		}
		proc.SLEEP(4)
		v := <-n.In
		if v > 0 { // this is child, siblings send negative
			sum += v
		}
	}

	// 5. send sum to parent
	neighbors[parIdx].Out <- sum
}

/*************************************************************
 * Child
 * 1. receive value from parent
 * 2. send its neg. value to others (children and siblings)
 * 3. receive values from children(positive) and siblings(negative)
 * 4. sums values from children (positive)`
 * 5. send sum to parent
*************************************************************/
func (proc *Proc) runChild2(neighbors NeighborChans) {
	numNeighbors := len(neighbors)

	proc.SLEEP(2)
	// 1. read from parent In tree

	/***********************************************************
	 * select {
	 * case <-neighbors[0].In:
	 *     parIdx = 0
	 * case <-neighbors[1].In:
	 *     parIdx = 1
	 * case <-neighbors[2].In:
	 *     parIdx = 2
	 * }
	 *
	 *
	 * package reflect
	 * type SelectCase struct {
	 *     Dir  SelectDir // direction of case
	 *     Chan Value     // channel to use (for send or receive)
	 *     Send Value     // value to send (for send)
	 * }
	 * type SelectDir 1.1
	 * A SelectDir describes the communication direction of a select case.
	 *
	 * type SelectDir int
	 * const (
	 *     SelectSend    SelectDir // case Chan <- Send
	 *     SelectRecv              // case <-Chan:
	 *     SelectDefault           // default
	 * )
	 *
	 * func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)
	 *
	***********************************************************/
	cases := make([]reflect.SelectCase, numNeighbors)
	for i := 0; i < numNeighbors; i++ {
		cases[i].Dir = reflect.SelectRecv
		cases[i].Chan = reflect.ValueOf(neighbors[i].In)
		cases[i].Send = reflect.ValueOf(nil)
	}
	var parIdx int = -1
	var parVal reflect.Value
	parIdx, parVal, _ = reflect.Select(cases)

	if Debug {
		fmt.Println("Child with val ", proc.m_MyVal, ", par idx ", parIdx,
			", par value ", parVal)
	}

	/***********************************************************/

	// 2. send to all but parent (children and siblings)
	for i := 0; i < numNeighbors; i++ {
		if i == parIdx {
			proc.SLEEP(2)
			continue
		}
		proc.SLEEP(2)
		neighbors[i].Out <- (-proc.m_MyVal)
	}

	// 3. receive from children and siblings
	// from children positive, from siblings negative
	// 4. sum from children
	sum := (proc.m_MyVal)
	for i := 0; i < numNeighbors; i++ {
		if i == parIdx {
			continue
		}
		proc.SLEEP(4)
		v := <-neighbors[i].In
		if v > 0 { // this is child, siblings send negative
			sum += v
		}
	}

	// 5. send sum to parent
	if Debug {
		fmt.Println("Child back to parent: my val ", proc.m_MyVal,
			", sum ", sum, ", par val ", parVal)
	}
	neighbors[parIdx].Out <- sum
}
