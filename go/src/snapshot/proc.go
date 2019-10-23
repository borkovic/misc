package snapshot

/*********************************************************
*********************************************************/
import (
	"fmt"
	//"math/rand"
	"time"
)

/*********************************************************
*********************************************************/
func (proc *Proc) SLEEP(nn int) {
	var n int = proc.rng.Intn(nn)
	m := 1000 * n
	//t := time.Duration(m) * time.Millisecond
	t := time.Duration(m) * time.Microsecond
	time.Sleep(t)
}

/*********************************************************
 * This process recieves value from top.
 * If negative, that's root, otherwise it's child
*********************************************************/
func (proc *Proc) Run(topChan *VertChanPair,
	neighbors NeighborChans) {
	var ok bool

	proc.myVal, ok = <-topChan.in
	if !ok {
		s := fmt.Sprintf("ERROR: Proc %d - bad receive from top", proc.id)
		panic(s)
	}
	if proc.myVal < 0 { // root
		if dbg() {
			fmt.Println("Run root, my val ", proc.myVal)
		}
		proc.myVal = -proc.myVal
		sum := proc.runRoot(neighbors)
		topChan.out <- sum
		close(topChan.out)
	} else { // slave
		if dbg() {
			fmt.Println("Run child, my val ",
				proc.myVal)
		}
		proc.runChild(neighbors)
		topChan.out <- proc.myVal
		close(topChan.out)
	}
}
