package snapshot

import (
	"fmt"
)

/*********************************************************
 * This process is root.
 * It sends its neg. value to all children.
 * Collects sums from the children
*********************************************************/
func (proc *Proc) runRoot(neighbors NeighborChans) Data {
	//numNeighbors := len(neighbors)
	for _, n := range neighbors {
		proc.SLEEP(88)
		n.out <- (-proc.myVal)
		close(n.out)
	}
	proc.SLEEP(4)
	var sum Data = (proc.myVal)
	for _, n := range neighbors {
		x, ok := <-n.in
		if !ok {
			s := fmt.Sprintf("ERROR: Root %d - bad receive from neighbor", proc.id)
			panic(s)
		}
		proc.SLEEP(8)
		if x > 0 {
			sum += x
		}
	}
	return Data(sum)
}
