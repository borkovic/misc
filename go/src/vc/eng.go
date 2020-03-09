package vc

func (eng *SeqEng) MakeEdgesOnEng() {
    numOps := OpIndexOnEng(len(eng.ops))
	eng.ops[0].sameEngSucc = &eng.ops[1]
	for x := numOps*0+1; x < numOps-1; x++ {
		eng.ops[x].sameEngPred = &eng.ops[x-1]
		eng.ops[x].sameEngSucc = &eng.ops[x+1]
	}
	eng.ops[numOps-1].sameEngPred = &eng.ops[numOps-2]
}

