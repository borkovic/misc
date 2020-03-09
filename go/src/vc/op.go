package vc

//**********************************************
func (op1 *Op) Before(op2 *Op) bool {
	return op1.ts.val(op1.engIdx) <= op2.ts.val(op1.engIdx)
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

func (op *Op) updateVc() {
	vc := &op.ts
	for _, pred := range op.crossPreds {
		vc2 := &pred.ts
		vc.Maximize(vc2)
	}
	vc.timestamps[op.engIdx]++
}
