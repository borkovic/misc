package snapshot

type ProcIdx int
type Data int

type IoChan struct {
	out chan<- Data
	in  <-chan Data
}

const (
	NumNeighbors = 3
)

type NeighborChans *[NumNeighbors]IoChan

type Proc struct {
	m_MyVal Data
}
