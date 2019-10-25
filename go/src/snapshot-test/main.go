package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)
import (
	"snapshot"
)

/*************************************************************
*************************************************************/
func main() {
	var seed int64 = -1
	var nProc snapshot.ProcIdx = -1
	var root snapshot.ProcIdx = -1
	var percChan int = -1
	var bias int = -1

	var err error
	var n int64

	N := len(os.Args)
	for i := 1; i < N; i++ {
		switch os.Args[i] {
		case "--seed":
			if i+1 >= N {
				panic("Not enough args for --seed")
			}
			seed, err = strconv.ParseInt(os.Args[i+1], 10, 64)
			if err != nil {
				panic("Processing --seed argument")
			}
			if seed < 0 {
				seed = -seed // always positive seed
			}
			i++
		case "--nproc":
			if i+1 >= N {
				panic("Not enough args for --nproc")
			}
			n, err = strconv.ParseInt(os.Args[i+1], 10, 32)
			if err != nil {
				panic("Processing --nproc argument")
			}
			nProc = snapshot.ProcIdx(n)
			if nProc <= 0 {
				panic("nproc must be positive")
			}
			i++
		case "--root":
			if nProc <= 0 {
				panic("Cannot choose root before nproc")
			}
			if i+1 >= N {
				panic("Not enough args for --root")
			}
			n, err = strconv.ParseInt(os.Args[i+1], 10, 32)
			if err != nil {
				panic("Processing --root argument")
			}
			root = snapshot.ProcIdx(n)
			if root < 0 || root >= nProc {
				panic("Root must be in [0,nproc)")
			}
			i++
		case "--bias":
			if i+1 >= N {
				panic("Not enough args for --bias")
			}
			n, err = strconv.ParseInt(os.Args[i+1], 10, 32)
			if err != nil {
				panic("Processing --bias argument")
			}
			bias = int(n)
			if bias < 1 || bias >= 99 {
				panic("Bias must be in [1,99]")
			}
			i++
		default:
			panic("Unknown option")
		}
	}

	if seed < 0 {
		seed = time.Now().UnixNano()
		if seed < 0 {
			seed = -seed // always positive seed
		}
	}
	rng := rand.New(rand.NewSource(seed))

	if nProc <= 0 {
		nProc = snapshot.ProcIdx(300 + rng.Intn(40))
	}
	if root < 0 {
		root = snapshot.ProcIdx(rng.Intn(int(nProc)))
	}
	if percChan <= 0 {
		percChan = 2 + rng.Intn(20)
	}

	if bias < 1 {
		bias = rng.Intn(10)
	}

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
