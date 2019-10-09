package snapshot

/*************************************************************
*************************************************************/
import (
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
	for _, n := range neighbors {
		n.Out <- (-proc.m_MyVal)
	}
	proc.SLEEP(4)
	sum := proc.m_MyVal
	for i := 0; i < NumNeighbors; i++ {
		x := <-neighbors[i].In
		proc.SLEEP(4)
		sum += x
	}
	return sum
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
	for i := 0; i < NumNeighbors; i++ {
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
	sum := proc.m_MyVal
	for i := 0; i < NumNeighbors; i++ {
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
	var parIdx int = -1

	proc.SLEEP(4)
	// 1. read from parent In tree

	numNeighbors := len(neighbors)

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
	parIdx, _, _ = reflect.Select(cases)

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
	for i := 0; i < NumNeighbors; i++ {
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
	sum := proc.m_MyVal
	for i := 0; i < NumNeighbors; i++ {
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
 * This process recieves value from top.
 * If negative, that's root, otherwise it's child
*************************************************************/
func (proc *Proc) Run(topChan *IoChans, neighbors NeighborChans) {
	r0 := time.Now().UnixNano()
	proc.RNG = rand.New(rand.NewSource(r0))

	proc.m_MyVal = <-topChan.In
	if proc.m_MyVal < 0 { // master
		proc.m_MyVal = -proc.m_MyVal
		sum := proc.runRoot(neighbors)
		topChan.Out <- sum
	} else { // slave
		proc.runChild2(neighbors)
	}
}
