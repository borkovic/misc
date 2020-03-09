package vc

//**********************************************
func (graph *Graph) MkGraph(extGraph ExtGraph) {
	extOp2Op := make(map[ExtOp]*Op)
	graph.makeEngines(extGraph, extOp2Op)
	numEngines := EngineIdx(len(graph.engines))
	for engIdx := EngineIdx(0); engIdx < numEngines; engIdx++ {
		eng := &graph.engines[engIdx]
		eng.makeEdgesOnEng()
	}
	graph.makeCrossEdges(extOp2Op)
	graph.updateVcs()
}

func (graph *Graph) updateVcs() {
	numEngines := EngineIdx(len(graph.engines))
	for engIdx := numEngines*0; engIdx < numEngines; engIdx++ {
		eng := &graph.engines[engIdx]
		numOps := OpIndexOnEng(len(eng.ops))
		for opIdx := numOps*0; opIdx < numOps; opIdx++ {
			op := &eng.ops[opIdx]
			op.updateVc()
		}
	}
}

//**********************************************
func (graph *Graph) makeCrossEdges(extOp2Op map[ExtOp]*Op) {
	numEngines := EngineIdx(len(graph.engines))
	for engIdx := numEngines*0; engIdx < numEngines; engIdx++ {
		eng := &graph.engines[engIdx]
		numOps := OpIndexOnEng(len(eng.ops))
		for opIdx := numOps*0; opIdx < numOps; opIdx++ {
			op := &eng.ops[opIdx]
			op.makeCrossEdges(extOp2Op)

		}
	}
}

//**********************************************
func (graph *Graph)makeEngines(extGraph ExtGraph,
		extOp2Op map[ExtOp]*Op) {
	numOps := extGraph.NumOps()
	var maxEngIdx EngineIdx = -1
	for i := numOps*0; i < numOps; i++ {
		engIdx := extGraph.Op(i).EngineIdx()
		if engIdx > maxEngIdx {
			maxEngIdx = engIdx 
		}
	}
	numEngs := maxEngIdx + 1
	graph.engines = make([]SeqEng, numEngs)


	// calc num ops per each engine
	numOpsPerEngine := make([]OpIndexOnEng, numEngs)
	for i := numOps*0; i < numOps; i++ {
		numOpsPerEngine[extGraph.Op(i).EngineIdx()]++
	}

	// allocate ops on engines
	for engIdx := EngineIdx(0); engIdx < maxEngIdx; engIdx++ {
		eng := &graph.engines[engIdx]
		eng.numOps = 0
		eng.ops = make([]Op, numOpsPerEngine[engIdx])
	}
	// intialize ops on engines
	for i := numOps*0; i < numOps; i++ {
		extOp := extGraph.Op(i)
		engIdx := extOp.EngineIdx()
		eng := &graph.engines[engIdx]
		op := &eng.ops[eng.numOps]
		op.engIdx = engIdx
		op.idxOnEng = eng.numOps
		op.extIdx = OpIndexExternal(i)
		op.extOp = extOp
		extOp2Op[extOp] = op
		eng.numOps++
	}
	// check num ops per engine
	for engIdx := EngineIdx(0); engIdx < maxEngIdx; engIdx++ {
		eng := &graph.engines[engIdx]
		if eng.numOps != numOpsPerEngine[engIdx] {
			panic("Wrong number of ops on engine")
		}
	}
}

