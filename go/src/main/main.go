package main

import "io"
import "fmt"
import "strings"
import "bytes"
import "encoding/hex"
import "crypto/sha1"

import "gitlet"

func main() {
	type StringSha struct {
		s            string
		expected_sha string
	}
	ss2 := [2]StringSha{
		{"Foo", "201a6b3053cc1422d2c3670b62616221d2290929"},
		{"His money is twice tainted: 'taint yours and 'taint mine.",
			"597f6a540010f94c15d71806a99a2c8710e747bd"},
	}

	{
		for _, ss := range ss2 {
			var sha gitlet.ShaId
			sha.ShaOfString(ss.s)
			fmt.Println("Sha1 of string", ("<" + ss.s + ">"), "is")
			fmt.Println("\t\t", sha.AsString())
			fmt.Println("Expected:\t", ss.expected_sha, "\n")
		}
	}

	{
		expected := "201a6b3053cc1422d2c3670b62616221d2290929"
		fname := "File"
		var sha gitlet.ShaId
		sha.ShaOfFile(fname)
		fmt.Println("Sha1 of io.copy file", fname, "is")
		fmt.Println("\t\t", sha.AsString())
		fmt.Println("Expected\t", expected, "\n")
	}
	{
		expected := "201a6b3053cc1422d2c3670b62616221d2290929"
		hasher := sha1.New()
		io.WriteString(hasher, "F")
		io.WriteString(hasher, "o")
		io.WriteString(hasher, "o")
		fmt.Println("Sha1 of chars <F,o,o> is")
		fmt.Println("\t\t", hex.EncodeToString(hasher.Sum(nil)))
		fmt.Println("Expected:\t", expected, "\n")
	}
	{
		for _, ss := range ss2 {
			src := strings.NewReader(ss.s)
			hasher := sha1.New()
			io.Copy(hasher, src)
			fmt.Println("Sha1 of io.copy string", ("<" + ss.s + ">"), "is")
			fmt.Println("\t\t", hex.EncodeToString(hasher.Sum(nil)))
			fmt.Println("Expected:\t", ss.expected_sha, "\n")
		}
	}
	{
		for _, ss := range ss2 {
			bs := []byte(ss.s)
			src := bytes.NewReader(bs)
			hasher := sha1.New()
			io.Copy(hasher, src)
			fmt.Println("Sha1 of io.copy bytes", ("<" + ss.s + ">"), "is")
			fmt.Println("\t\t", hex.EncodeToString(hasher.Sum(nil)))
			fmt.Println("Expected:\t", ss.expected_sha, "\n")
		}
	}
}
