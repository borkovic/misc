package main

import "fmt"

import "gitlet"

func main() {
	s := "Foo"
	h := gitlet.StringSha(s)
	fmt.Println("Sha1 of", s, "is", h)
}
