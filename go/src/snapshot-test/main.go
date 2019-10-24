package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"strconv"
)
import (
	"snapshot"
)

/*************************************************************
*************************************************************/
func main() {
	var seed int64
	if len(os.Args) > 2 && os.Args[1] == "--seed" {
		var err error
		seed,err = strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			panic("Bad argument " + os.Args[2])
		}
	} else {
		seed = time.Now().UnixNano()
	}
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
