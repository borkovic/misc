package snapshot

/*************************************************************
*************************************************************/
import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

/*************************************************************
*************************************************************/
func (proc *Proc) SLEEP(nn int) {
	var n int = proc.RNG.Intn(nn)
	m := 100 * n
	time.Sleep(time.Duration(m) * time.Millisecond)
}

/*************************************************************
 * This process is root.
 * It sends its neg. value to all children.
 * Collects sums from the children
*************************************************************/
func (proc *Proc) runRoot(neighbors NeighborChans) Data {
	numNeighbors := len(neighbors)
	for _, n := range neighbors {
		proc.SLEEP(88)
		n.Out <- (-proc.m_MyVal)
	}
	proc.SLEEP(4)
	var sum Data = (proc.m_MyVal)
	for i := 0; i < numNeighbors; i++ {
		x := <-neighbors[i].In
		proc.SLEEP(8)
		if x > 0 {
			sum += x
		}
	}
	return Data(sum)
}

/*************************************************************
 * Child
 * 1. receive value from parent
 * 2. send its neg. value to others (children and siblings)
 * 3. receive values from children(positive) and siblings(negative)
 * 4. sums values from children (positive)`
 * 5. send sum to parent
*************************************************************/
func (proc *Proc) runChild(neighbors NeighborChans) {
	numNeighbors := len(neighbors)
	var parIdx int = -1

	proc.SLEEP(4)
	// 1. read from parent In tree
	select {
	case <-neighbors[0].In:
		parIdx = 0
	case <-neighbors[1].In:
		parIdx = 1
	case <-neighbors[2].In:
		parIdx = 2
	}

	// 2. send to all but parent (children and siblings)
	for i := 0; i < numNeighbors; i++ {
		if i == parIdx {
			proc.SLEEP(4)
			continue
		}
		proc.SLEEP(4)
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

	/***********************************************************/
	/*
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
	 */
	cases := make([]reflect.SelectCase, numNeighbors)
	for i := 0; i < numNeighbors; i++ {
		cases[i].Dir = reflect.SelectRecv
		cases[i].Chan = reflect.ValueOf(neighbors[i].In)
		cases[i].Send = reflect.ValueOf(nil)
	}
	var parIdx int = -1
	var parVal reflect.Value
	parIdx, parVal, _ = reflect.Select(cases)
	fmt.Println("Child with val ", proc.m_MyVal, ", par idx ", parIdx, ", par value ", parVal)

	/*
		select {
		case <-neighbors[0].In:
			parIdx = 0
		case <-neighbors[1].In:
			parIdx = 1
		case <-neighbors[2].In:
			parIdx = 2
		}
	*/
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
	fmt.Println("Child back to parent: my val ", proc.m_MyVal,
		", sum ", sum, ", par val ", parVal)
	neighbors[parIdx].Out <- sum
}

/*************************************************************
 * This process recieves value from top.
 * If negative, that's root, otherwise it's child
*************************************************************/
func (proc *Proc) Run(topChan *VertChanPair, neighbors NeighborChans) {
	r0 := time.Now().UnixNano()
	proc.RNG = rand.New(rand.NewSource(r0))

	proc.m_MyVal = <-topChan.In
	if proc.m_MyVal < 0 { // root
		fmt.Println("Run root, my val ", proc.m_MyVal)
		proc.m_MyVal = -proc.m_MyVal
		sum := proc.runRoot(neighbors)
		topChan.Out <- sum
	} else { // slave
		fmt.Println("Run child, my val ", proc.m_MyVal)
		proc.runChild2(neighbors)
	}
}
