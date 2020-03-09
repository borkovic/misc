package vc

//**********************************************
type EngineIdx int16
type OpIndexOnEng int32
type OpIndexExternal int32

//**********************************************
type ExtOp interface {
	NumPreds() int16
	NumSuccs() int16
	Pred(n int16) ExtOp
	Succ(n int16) ExtOp
	EngineIdx() EngineIdx
}

//**********************************************
type ExtGraph interface {
	NumOps() int32
	Op(n int32) ExtOp
}


//**********************************************
type Op struct {
	engIdx   EngineIdx
	idxOnEng OpIndexOnEng
	extIdx   OpIndexExternal
	extOp    ExtOp
	sameEngPred *Op
	sameEngSucc *Op
	crossPreds	[]*Op
	crossSuccs  []*Op
	ts       VC
}

//**********************************************
type SeqEng struct {
	engIdx EngineIdx
	numOps OpIndexOnEng
	ops []Op
}

//**********************************************
type Graph struct {
	engines []SeqEng
	ops []Op
}

