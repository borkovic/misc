package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

import (
	hs "../heapsort"
	"../utils"
)

type (
	Vec   = sort.IntSlice
	Index = hs.Index
)

/***********************************************************/
func main() {
	/*
	   const N = 1 * 1000 * 1000
	   const N = 10
	   const N = 100*1000*1000
	   const N = 10 * 1000 * 1000
	*/
	ts := os.Args[1]
	N, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		panic(err)
	}

	v := make(Vec, N)
	for i := 0; i < len(v); i++ {
		v[i] = int(rand.Int31n(int32(N)))
	}

	start := time.Now()
	v.Sort() //sort.Ints(v)
	elapsed := time.Since(start)
	//elapsed *= 1000.0
	fmt.Printf("GOS: Sorting [%s]%T: %T  %v seconds\n",
		utils.PrintLong(N), v[0], elapsed, elapsed.Seconds())

	utils.CheckSliceSorted(v[:])
}
