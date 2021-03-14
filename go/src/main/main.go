package main

import "io"
import "fmt"
import "encoding/hex"
import "crypto/sha1"

import "gitlet"

func main() {
	//numa := [3]int{78, 79 ,80}
	ss := [2]string{
		"Foo",
		"His money is twice tainted: 'taint yours and 'taint mine."}

	//	for i, v := range pow {
	//		fmt.Printf("2**%d = %d\n", i, v)
	//	}
	for _, s := range ss {
		sign := gitlet.StringSha(s)
		fmt.Println("Sha1 of string", s, "is")
		fmt.Println("\t", sign, "\n")
	}

	{
		fname := "File"
		sign, _ := gitlet.FileSha(fname)
		fmt.Println("Sha1 of file", fname, "is")
		fmt.Println("\t", sign, "\n")
	}
	{
		hasher := sha1.New()
		io.WriteString(hasher, "F")
		io.WriteString(hasher, "o")
		io.WriteString(hasher, "o")
		fmt.Println("Sha1 of chars F,o,o", "", "is")
		fmt.Println("\t", hex.EncodeToString(hasher.Sum(nil)), "\n")
	}
}
