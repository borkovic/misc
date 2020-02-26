package vc

type Graph struct {
	engines []SeqEng
	ops []Op
}

func (Graph *graph) MkEngines {
	var maxEngIdx EngineIdx = -1
	for _, op := range graph.ops {
		engIdx := op.engIdx
		if engIdx > maxEngIdx {
			maxEngIdx = engIdx
		}
	}
	numEngines := maxEngIdx + 1
	graph.engines = make([]SeqEng, numEngines)

	N := len(graph.op)

	i := 0
	for i < N {
		s := i
		engIdx := graph.op[s].engIdx
		for i < N && graph.op[i].engIdx == engIdx {
			i++
		}
		if s < i {
			graph.engines[engIdx].ops = graph.op[s:i]
		}
	}
}

