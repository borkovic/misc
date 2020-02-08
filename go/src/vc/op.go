package vc

type Op struct {
	engIdx   EngineIdx
	idxOnEng OpIndexOnEng
	glbIdx 	 OpIndexGlobal
	ts       VC
}

func (op1 *Op) before(op2 *Op) bool {
	return op1.ts.val(op1.engIdx) <= op2.ts.val(op1.engIdx)
}

func (op1 *Op) after(op2 *Op) bool {
	return op2.before(op1)
}
