package main

import "io"
import "fmt"
import "encoding/hex"
import "crypto/sha1"

import "gitlet"

func main() {
	//numa := [3]int{78, 79 ,80}
	type StringSha struct {
		s string
		expected_sha string
	}
	ss2 := [2]StringSha{
			{"Foo", "201a6b3053cc1422d2c3670b62616221d2290929"},
			{"His money is twice tainted: 'taint yours and 'taint mine.",
	 		 "597f6a540010f94c15d71806a99a2c8710e747bd"},
	}
		

	//	for i, v := range pow {
	//		fmt.Printf("2**%d = %d\n", i, v)
	//	}
	for _, ss := range ss2 {
		var sha gitlet.ShaId 
		sha.ShaOfString(ss.s)
		fmt.Println("Sha1 of string", ("<" + ss.s + ">"), "is")
		fmt.Println("\t", sha.AsString())
		fmt.Println("Expected:", ss.expected_sha, "\n")
	}

	{ 
		expected := "201a6b3053cc1422d2c3670b62616221d2290929"
		fname := "File"
		var sha gitlet.ShaId 
		sha.ShaOfFile(fname)
		fmt.Println("Sha1 of file", fname, "is")
		fmt.Println("\t", sha.AsString())
		fmt.Println("Expected", expected, "\n")
	}
	{ 
		expected := "201a6b3053cc1422d2c3670b62616221d2290929"
		hasher := sha1.New()
		io.WriteString(hasher, "F")
		io.WriteString(hasher, "o")
		io.WriteString(hasher, "o")
		fmt.Println("Sha1 of chars <F,o,o>", "", "is")
		fmt.Println("\t", hex.EncodeToString(hasher.Sum(nil)))
		fmt.Println("Expected:", expected, "\n")
	}
}
