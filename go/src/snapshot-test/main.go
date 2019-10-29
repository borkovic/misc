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
	var mm snapshot.ProbInt = -1
	var nn snapshot.ProbInt = -1
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
				panic("Error processing --seed argument")
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
				panic("Error processing --nproc argument")
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
				panic("Error processing --root argument")
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
				panic("Error processing --bias argument")
			}
			bias = int(n)
			if bias < 1 || bias >= 29 {
				panic("Bias must be in [1,29]")
			}
			i++
		case "--chan-prob":
			if i+2 >= N {
				panic("Not enough args for --chan-prob")
			}
			m, err := strconv.ParseInt(os.Args[i+1], 10, 32)
			if err != nil {
				panic("Error processing --chan-prob argument")
			}
			mm = snapshot.ProbInt(m)
			n, err := strconv.ParseInt(os.Args[i+2], 10, 32)
			if err != nil {
				panic("Error processing --chan-prob argument")
			}
			nn = snapshot.ProbInt(n)
			if nn < 2 || mm < 1 || mm > nn-1 {
				s := fmt.Sprintf("Bias must be in [1,%d]", nn-1)
				panic(s)
			}
			i += 2
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
	} else {
		nProc = snapshot.ProcIdx(int(nProc) + rng.Intn(40))
	}
	if root < 0 {
		root = snapshot.ProcIdx(rng.Intn(int(nProc)))
	}
	if mm < 1 || nn < 2 {
		//fmt.Println("Creating chan probs mm:", mm, ", nn:", nn)
		mm = 2 + snapshot.ProbInt(rng.Intn(20))
		nn = snapshot.ProbInt(100)
	}
	if bias < 1 {
		bias = rng.Intn(2)
	}

	fmt.Print("Num proc: ", nProc,
		", Root: ", root,
		", Chan prob: ", mm, "/", nn,
		", Bias: ", bias,
		", Seed: ", seed,
		"\n")

	graph := new(snapshot.Graph)
	graph.BuildAndCollectData(nProc, root, mm, nn, bias, rng.Int63())
	fmt.Println("Done")
}
