package snapshot

import (
	"math/rand"
)

type ProcIdx int
type Data int64

type Proc struct {
	id    ProcIdx
	myVal Data
	RNG   *rand.Rand
}

type Graph struct {
	numberProcs ProcIdx
	root        ProcIdx
	procs       []Proc
	neighbors   [][]HorizChanPair
	tops        []VertChanPair
	driverTops  []VertChanPair
}
