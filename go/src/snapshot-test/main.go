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

	seed := time.Now().UnixNano()
	if seed < 0 {
		seed = -seed // always positive seed
	}
	rng := rand.New(rand.NewSource(seed))

	bias := rng.Intn(10)
	nProc := snapshot.ProcIdx(300 + rng.Intn(40))
	root := snapshot.ProcIdx(rng.Intn(int(nProc)))
	percChan := 2 + rng.Intn(20)

	fmt.Print("Num proc: ", nProc,
		", Root: ", root,
		", Chan prob: ", percChan, "/100",
		", Bias: ", bias,
		", Seed: ", seed,
		"\n")

	graph := new(snapshot.Graph)
	graph.BuildAndCollectData(nProc, root, bias, percChan, rng.Int63())
	fmt.Println("Done")
}
