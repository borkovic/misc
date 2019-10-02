package snapshot

type ProcIdx int
type Data int

type BiChan = chan Data
type OChan  = chan<- Data
type IChan  = <-chan Data

type IoChans struct {
	Out OChan
	In  IChan
}

const (
	NumNeighbors = 3
	NumProc = 4
)

type NeighborChans = []IoChans

type Proc struct {
	m_MyVal Data
}
