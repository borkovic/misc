package vc

//**********************************************
func (op1 *Op) Before(op2 *Op) bool {
	engIdx := op1.engIdx
	return op1.vc.val(engIdx) <= op2.vc.val(engIdx)
}

//**********************************************
func (op1 *Op) After(op2 *Op) bool {
	return op2.Before(op1)
}

//**********************************************
func (op *Op) makeCrossEdges(extOp2Op map[ExtOp]*Op) {
	engIdx := op.engIdx
	extOp := op.extOp
	numPreds := extOp.NumPreds()
	for p := numPreds * 0; p < numPreds; p++ {
		extPred := extOp.Pred(p)
		if extPred.EngineIdx() == engIdx {
			continue
		}
		pred := extOp2Op[extPred]
		op.crossPreds = append(op.crossPreds, pred)
	}
	numSuccs := extOp.NumSuccs()
	for s := numSuccs * 0; s < numSuccs; s++ {
		extSucc := extOp.Succ(s)
		if extSucc.EngineIdx() == engIdx {
			continue
		}
		succ := extOp2Op[extSucc]
		op.crossSuccs = append(op.crossSuccs, succ)
	}
}

//**********************************************
func (op *Op) updateVc() {
	vc := &op.vc
	for _, pred := range op.crossPreds {
		vc2 := &pred.vc
		vc.Maximize(vc2)
	}
	vc.Incr(op.engIdx, 1)
}
