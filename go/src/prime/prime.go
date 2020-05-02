package main

import "fmt"

//var MAX int = 1000
var MAX int = 941

var END_MARKER int = -1
func is_end(k int) bool {
	return k <= END_MARKER
}

func generate_integers() <-chan int {
	c := make(chan int)

	var ch chan<- int = c
	go func() {
		for i := 2; ; i++ {
			ch <- i
			if i > MAX {
				break
			}
		}
		ch <- END_MARKER  // end marker
	}()

	return c
}

func filter_multiples(in <-chan int, prime int) <-chan int {
	o := make(chan int)

	var out chan<- int = o
	go func() {
		for {
			i := <-in
			if is_end(i) {
				out <- i
				break
			}
			if i%prime != 0 {
				out <- i
			}
		}
	}()

	return o
}

func sieve() <-chan int {
	o := make(chan int)

	var out chan<- int = o
	go func() {
		ch := generate_integers()
		for {
			prime := <-ch
			out <- prime
			if is_end(prime) {
				break
			}
			ch = filter_multiples(ch, prime)
		}
	}()

	return o
}

func main() {
	primes := sieve()
	N := 10
	i := 0
	need_nl := true
	for {
		p := <-primes
		if is_end(p) {
			break
		}
		fmt.Print(p, " ")
		i++
		need_nl = true
		if i == N {
			i = 0
			fmt.Println("")
			need_nl = false
		}
	}
	if need_nl {
		fmt.Println("")
	}
}
