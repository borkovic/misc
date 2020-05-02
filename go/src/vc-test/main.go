package main

import (
	"fmt"
)
import (
	"vc"
)


//**********************************************
type MyOp struct {
	succs []*MyOp
	preds []*MyOp
}

type MyGraph struct {
	ops []MyOp
}

//**********************************************
func (op *MyOp) EngineIdx() vc.EngineIdx {
	return 0
}

func (op *MyOp) NumPreds() int16 {
	return int16(len(op.preds))
}

func (op *MyOp) NumSuccs() int16 {
	return int16(len(op.succs))
}



func (op *MyOp) Pred(n int16) vc.ExtOp {
	return op.preds[n]
}

func (op *MyOp) Succ(n int16) vc.ExtOp {
	return op.succs[n]
}


//**********************************************
func (g *MyGraph) NumOps() int32 {
	return int32(len(g.ops))
}

func (g *MyGraph) Op(n int32) vc.ExtOp {
	return &g.ops[n]
}

//**********************************************
func main() {
	var mygraph MyGraph
	var vcgraph vc.Graph
	fmt.Println("Making graph")
	vcgraph.MkGraph(&mygraph)
}
