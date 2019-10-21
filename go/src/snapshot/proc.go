package snapshot

/*********************************************************
*********************************************************/
import (
	"fmt"
	"math/rand"
	"time"
)

const (
	//Debug bool = true
	Debug bool = false
)

/*********************************************************
*********************************************************/
func (proc *Proc) SLEEP(nn int) {
	var n int = proc.RNG.Intn(nn)
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
	r0 := time.Now().UnixNano()
	proc.RNG = rand.New(rand.NewSource(r0))

	proc.myVal, ok = <-topChan.in
	if !ok {
		s := fmt.Sprintf("ERROR: Proc %d - bad receive from top", proc.id)
		panic(s)
	}
	if proc.myVal < 0 { // root
		if Debug {
			fmt.Println("Run root, my val ", proc.myVal)
		}
		proc.myVal = -proc.myVal
		sum := proc.runRoot(neighbors)
		topChan.out <- sum
		close(topChan.out)
	} else { // slave
		if Debug {
			fmt.Println("Run child, my val ",
				proc.myVal)
		}
		proc.runChild(neighbors)
		topChan.out <- proc.myVal
		close(topChan.out)
	}
}
