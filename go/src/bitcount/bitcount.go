package main

import (
	"fmt"
)

//*********************************************
type BitArray uint32

//*********************************************
func countBits(v BitArray) int32 {
	var x BitArray = v
	var cnt int32 = 0

	for x > 0 {
		cnt++
		var y BitArray = x - 1
		x &= y
	}

	return cnt
}

//*********************************************
func printBitCount(v BitArray) {
	cnt := countBits(v)
	fmt.Println("Value: ", v, ", cnt: ", cnt)

}

//*********************************************
func main() {
	var v BitArray
	for v = 0; v < 32; v++ {
		printBitCount(v)
	}
}
