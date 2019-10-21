package snapshot

import (
	"math/rand"
)

type ProcIdx int
type Data int32

type Proc struct {
	id    ProcIdx
	myVal Data
	RNG   *rand.Rand
}

type Graph struct {
	NumberProcs ProcIdx
	Root        ProcIdx
	Procs       []Proc
	Neighbors   [][]HorizChanPair
	tops        []VertChanPair
	driverTops  []VertChanPair
}
