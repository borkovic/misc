package vc

type ExtGraph interface {
	NumOps() int
	Op(n int) ExtOp
}

type Graph struct {
	engines []SeqEng
	ops []Op
}

func (graph *Graph) MkEngines() {
	var maxEngIdx EngineIdx = -1
	for _, op := range graph.ops {
		engIdx := op.engIdx
		if engIdx > maxEngIdx {
			maxEngIdx = engIdx
		}
	}
	numEngines := maxEngIdx + 1
	graph.engines = make([]SeqEng, numEngines)

	N := len(graph.ops)

	i := 0
	for i < N {
		s := i
		engIdx := graph.ops[s].engIdx
		for i < N && graph.ops[i].engIdx == engIdx {
			i++
		}
		if s < i {
			graph.engines[engIdx].ops = graph.ops[s:i]
		}
	}
}

func (graph *Graph) MkGraph(extGraph ExtGraph) {
	var maxEngIdx EngineIdx = -1
	numExtOps := extGraph.NumOps()
	for i := 0; i < numExtOps; i++ {
		extOp := extGraph.Op(i)
		engIdx := extOp.Engine()
		if engIdx > maxEngIdx {
			maxEngIdx = engIdx 
		}
	}
}

