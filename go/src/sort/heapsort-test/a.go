package main

import "fmt"

type Index int

func main() {
	var v [20]int16
	v[3] = 17
	var i Index = 3

	fmt.Println(v[i])
}
