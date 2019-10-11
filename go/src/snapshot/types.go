package snapshot

import (
	"math/rand"
)

type ProcIdx int
type Data int32
type HorizData int32
type VertData int32

type HorizBidirChan = chan HorizData
type HorizInChan = <-chan HorizData
type HorizOutChan = chan<- HorizData

type HorizChanPair struct {
	In  HorizInChan
	Out HorizOutChan
}

type VertBidirChan = chan Data
type VertInChan = <-chan Data
type VertOutChan = chan<- Data
type VertChanPair struct {
	In  VertInChan
	Out VertOutChan
}

const (
	NumNeighbors = 6
	NumProc      = 10
)

type NeighborChans = []HorizChanPair

type Proc struct {
	m_MyVal Data
	RNG     *rand.Rand
}
