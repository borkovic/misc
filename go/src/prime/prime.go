package main

import "fmt"

func generate_integers() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter_multiples(in <-chan int, prime int) <-chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() <-chan int {
	out := make(chan int)
	go func() {
		ch := generate_integers()
		for {
			prime := <-ch
			out <- prime
			ch = filter_multiples(ch, prime)
		}
	}()
	return out
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
