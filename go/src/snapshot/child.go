package snapshot

import (
	"fmt"
	"reflect"
)

/*********************************************************
 * Child
 * 1. receive value from parent
 * 2. send its neg. value to others (children/siblings)
 * 3. receive values from children(+) and siblings(-)
 * 4. sums values from children (+)
 * 5. send sum to parent
*********************************************************/
/***************************************************
 * var parIdx int = -1
 * var parVal Data
 * select {
 * case parVal = <-neighbors[0].in:
 *     parIdx = 0
 * case parVal <-neighbors[1].in:
 *     parIdx = 1
 * case parVal <-neighbors[2].in:
 *     parIdx = 2
 * }
 *
 *=================================================
 * package reflect
 * type SelectCase struct {
 *     Dir  SelectDir // direction of case
 *     Chan Value     // chan type (send or receive)
 *     Send Value     // value to send (for send)
 * }
 * type SelectDir 1.1
 * A SelectDir describes the communication direction
 * of a select case.
 *
 * type SelectDir int
 * const (
 *     SelectSend    SelectDir // case Chan<- Send
 *     SelectRecv              // case <-Chan:
 *     SelectDefault           // default
 * )
 *
 * func Select(cases []SelectCase) (chosen int,
 * 			   recv Value, recvOK bool)
 *
***************************************************/
func (proc *Proc) runChild(neighbors NeighborChans) {
	numNeighbors := len(neighbors)

	proc.SLEEP(2)
	// 1. read from parent in tree
	cases := make([]reflect.SelectCase, numNeighbors)
	for i := 0; i < numNeighbors; i++ {
		cases[i].Dir = reflect.SelectRecv
		cases[i].Chan = reflect.ValueOf(neighbors[i].in)
		cases[i].Send = reflect.ValueOf(nil)
	}
	parIdx, parVal, recvOk := reflect.Select(cases)
	if !recvOk {
		s := fmt.Sprintf("ERROR: Child %d - bad select", proc.id)
		panic(s)
	}

	/*-------------------------------------------------*/
	if dbg() {
		fmt.Println("Child with val ", proc.myVal,
			", par idx ", parIdx,
			", par value ", parVal)
	}

	/*-------------------------------------------------*/

	// 2. send neg to all but parent (children and
	//    siblings)
	for i := int(0); i < numNeighbors; i++ {
		if i == parIdx {
			proc.SLEEP(3)
			continue
		}
		proc.SLEEP(2)
		neighbors[i].out <- (-proc.myVal)
		close(neighbors[i].out)
	}

	// 3. receive from children (+) and siblings (-)
	// 4. sum from children(+)
	sum := (proc.myVal)
	for i := 0; i < numNeighbors; i++ {
		if i == parIdx {
			continue
		}
		proc.SLEEP(4)
		v, ok := <-neighbors[i].in
		if !ok {
			s := fmt.Sprintf("ERROR: Child %d - bad receive from neighbor", proc.id)
			panic(s)
		}
		if v > 0 { // this is child, siblings negative
			sum += v
		}
	}

	// 5. send sum to parent
	if dbg() {
		fmt.Println("Child back to parent:",
			" my val ",
			proc.myVal,
			", sum ", sum, ", par val ", parVal)
	}
	neighbors[parIdx].out <- sum
	close(neighbors[parIdx].out)
}
