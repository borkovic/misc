package vc

//**********************************************
type EngineIdx int16

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
