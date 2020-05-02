package main

import "fmt"

var MAX int = 1000

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
		ch <- -1  // end marker
	}()

	return c
}

func filter_multiples(in <-chan int, prime int) <-chan int {
	o := make(chan int)

	var out chan<- int = o
	go func() {
		for {
			i := <-in
			if i < 0 {
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
			if prime < 0 {
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
	for {
		p := <-primes
		if p < 0 {
			break
		}
		fmt.Print(p, " ")
		i++
		if i == N {
			i = 0
			fmt.Println("")
		}
	}
	fmt.Println("")
}
