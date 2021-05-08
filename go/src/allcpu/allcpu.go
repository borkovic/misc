
package main

import (
	"fmt"
	"runtime"
)

type bidirchan = chan int
type readchan = <-chan int
type writechan = chan<- int

//func loopfunc(c writechan) {
func loopfunc(c chan<- int) {
	for {
	}
	c<- 0
}

func main() {
	var ncpu int = runtime.NumCPU()
	fmt.Println("Num CPU:", ncpu)

	//cb := make(bidirchan, ncpu)
	cb := make(chan int, ncpu)

	//cr := readchan(cb)
	cr := <-chan int(cb)

	for i := ncpu-1; i >= 0; i-- {
		go loopfunc(cb)
	}
	var sum int = 0
	for i := ncpu-1; i >= 0; i-- {
		a := <-cr
		sum += a
	}
	fmt.Println("Sum:", sum)
}
