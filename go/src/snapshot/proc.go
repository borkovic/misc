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

	proc.m_MyVal,ok  = <-topChan.In
	if ! ok {
		panic("Recieve from closed top chan")
	}
	if proc.m_MyVal < 0 { // root
		if Debug {
			fmt.Println("Run root, my val ", proc.m_MyVal)
		}
		proc.m_MyVal = -proc.m_MyVal
		sum := proc.runRoot(neighbors)
		topChan.Out<- sum
		close(topChan.Out)
	} else { // slave
		if Debug {
			fmt.Println("Run child, my val ",
						proc.m_MyVal)
		}
		proc.runChild2(neighbors)
		topChan.Out<- proc.m_MyVal
		close(topChan.Out)
	}
}
