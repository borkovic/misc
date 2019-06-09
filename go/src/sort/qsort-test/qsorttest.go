package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

import (
	"../qsort"
	"../utils"
)

type ValType = utils.ValType
type IndexType = utils.IndexType

func main() {
	const debug bool = false

	ts := os.Args[1]
	N, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		panic(err)
	}
	v := make([]ValType, N)
	for i := 0; i < len(v); i++ {
		v[i] = ValType(rand.Int31n(int32(N)))
	}
	if debug {
		fmt.Println(v)
	}
	start := time.Now()

	qsort.QSort2(v[:])

	elapsed := time.Since(start)
	//elapsed *= 1000.0
	fmt.Printf("GO: Sorting [%s]%T: %T  %v seconds\n",
		utils.PrintLong(N), v[0], elapsed, elapsed.Seconds())
	if debug {
		fmt.Println(v)
	}

	utils.CheckSorted(v)
}
