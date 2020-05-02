package main

import "fmt"

func generate_integers() <-chan int {
	c := make(chan int)

	var ch chan<- int = c
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()

	return c
}

func filter_multiples(in <-chan int, prime int) <-chan int {
	o := make(chan int)

	var out chan<- int = o
	go func() {
		for {
			if i := <-in; i%prime != 0 {
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
		fmt.Print(<-primes, " ")
		i++
		if i == N {
			i = 0
			fmt.Println("")
		}
	}
}
