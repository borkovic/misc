package main

import (
	"fmt"
	"math/rand"
	"time"
)
import (
	"snapshot"
)

/*************************************************************
*************************************************************/
func main() {

	r0 := time.Now().UnixNano()
	RNG := rand.New(rand.NewSource(r0))

	bias := RNG.Intn(10)
	nProc := snapshot.ProcIdx(100 + RNG.Intn(40))
	root := snapshot.ProcIdx(RNG.Intn(int(nProc)))
	percChan := 2 + RNG.Intn(20)

	fmt.Print("Num proc: ", nProc,
		", Root: ", root,
		", Chan prob: ", percChan, "/100",
		", Bias: ", bias,
		"\n")

	graph := new(snapshot.Graph)
	graph.BuildAndCollectData(nProc, root, bias, percChan)
}
