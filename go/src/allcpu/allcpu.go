package main

import (
	"fmt"
	"runtime"
)

type bidirChan = chan int
type readChan = <-chan int
type writeChan = chan<- int

func loopfunc(c writeChan, i int) { // func loopfunc(c chan<- int, i int)
	for { // forever
	}
	c <- i
}

func main() {
	var nCpu int = runtime.NumCPU()
	fmt.Println("Num CPUs:", nCpu)

	cb := make(bidirChan, nCpu) //cb := make(chan int, nCpu)

	cr := readChan(cb) //cr := <-chan int(cb)

	for i := int(0); i < nCpu; i++ {
		go loopfunc(cb, i)
	}
	var sum int = 0
	for i := int(0); i < nCpu; i++ {
		a := <-cr
		sum += a
	}
	fmt.Println("Sum:", sum)
}
