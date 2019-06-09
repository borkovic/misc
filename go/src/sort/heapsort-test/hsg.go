package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)
import (
	hs "../heapsort"
	"../utils"
)

type (
	Vec     = hs.Vec
	CmpFunc = hs.CmpFunc
	Index   = hs.Index
	Value   = hs.Value
)

/***********************************************************/
func checkSorted(v Vec, cmp CmpFunc) {
	last := hs.Len(v) - 1
	ok := true
	for k := Index(0); k < last-1; k++ {
		if cmp(v[k], v[k+1]) < 0 {
			fmt.Println("Error: v[", k, "]=", v[k], "v[", k+1, "]=", v[k+1])
			ok = false
		}
	}
	if ok {
		fmt.Println("hs1: OK")
	} else {
		fmt.Println("hs1: BAD")
	}
	//fmt.Println("C")
	//prHeap(v[:], 0, "")
	//fmt.Println(v)
}

/***********************************************************/
func prHeap(v Vec, k Index, ident string) {
	fmt.Println(ident, v[k])
	last := hs.Len(v) - 1
	lCld := hs.LeftCld(k)
	rCld := hs.RightCld(k)
	if lCld <= last {
		prHeap(v, lCld, ident+"  ")
	}
	if rCld <= last {
		prHeap(v, rCld, ident+"  ")
	}
}

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

	v := make([]Value, N) // [N]Value
	for i := 0; i < len(v); i++ {
		v[i] = Value(rand.Int31n(int32(N)))
	}
	cmp := hs.CmpGT

	start := time.Now()
	hs.Heapsort(v[:], cmp)
	elapsed := time.Since(start)
	//elapsed *= 1000.0
	fmt.Printf("GO: Sorting [%s]%T: %T  %v seconds\n",
		utils.PrintLong(N), v[0], elapsed, elapsed.Seconds())

	checkSorted(v[:], cmp)
}
