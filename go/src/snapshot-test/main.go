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
	percChan := 5

	fmt.Println("Num proc ", nProc, ", %chan ", percChan,
			", Bias is ", bias, ", root is ", root)

	snapshot.Driver(nProc, root, bias, percChan)
}