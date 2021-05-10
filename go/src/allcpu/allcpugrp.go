package main

import (
	"fmt"
	"runtime"
	"sync"
)

func loopfunc(grp *sync.WaitGroup, i int) {
	for { // forever
	}
	grp.Done()
}

func main() {
	var nCpu int = runtime.NumCPU()
	fmt.Println("Num CPUs:", nCpu)

	var grp sync.WaitGroup

	grp.Add(nCpu)
	for i := int(0); i < nCpu; i++ {
		go loopfunc(&grp, i)
	}
	grp.Wait()
}
