package vc

//**********************************************
type OpIndexOnEng int32
type OpIndexExternal int32

//**********************************************
type Op struct {
	engIdx      EngineIdx
	idxOnEng    OpIndexOnEng
	extIdx      OpIndexExternal
	extOp       ExtOp
	sameEngPred *Op
	sameEngSucc *Op
	crossPreds  []*Op
	crossSuccs  []*Op
	ts          VC
}

//**********************************************
type SeqEng struct {
	engIdx EngineIdx
	numOps OpIndexOnEng
	ops    []Op
}

//**********************************************
type Graph struct {
	engines []SeqEng
	ops     []Op
}
