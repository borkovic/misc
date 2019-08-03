//***************************************************************************
package qsort

import (
	"fmt"
)

//***************************************************************************
// Recursive quicksort.
// 1. Partition into two parts. Note that pivot_high==pivot_low.
// 2. Recurse only on the shorter part to limit stack depth to log(N)
// 3. Continue looping with the longer part
//***************************************************************************
func qsort2r(v []ValType) {
	const debug bool = false

	if debug {
		fmt.Println("Q:", v)
	}
	Begin := IndexType(0)
	End := IndexType(len(v))
	if debug {
		fmt.Println(len(v))
	}

	for Begin+1 < End {
		if debug {
			fmt.Println("F be:", Begin, End)
		}
		if debug {
			fmt.Println("F v[be]:", v[Begin:End])
		}
		//medVal := medianSwap(v[Begin:End])
		medVal := median(v[Begin:End])
		if debug {
			fmt.Println("F medVal:", medVal)
		}
		p, q := dnf2(v[Begin:End], medVal, medVal) // Must call with equal pivots
		if debug {
			fmt.Println("F pq:", p, q)
		}
		if debug {
			fmt.Println("F v[be]:", v[Begin:End])
		}
		p += Begin
		q += Begin
		if debug && p < Begin {
			panic("bad p")
		}
		if debug && q > End {
			panic("bad p")
		}
		//  x in [Begin, p) =>  v[x] < pivot1
		//  x in [p, q)     =>  pivot1 <= v[x] <= pivot2
		//  x in [p, End)   =>  pivot2 < v[x]

		leftSize := p - Begin
		rightSize := End - q
		if leftSize <= rightSize {
			if leftSize > 1 {
				qsort2r(v[Begin:p])
			}
			Begin = q
		} else {
			if rightSize > 1 {
				qsort2r(v[q:End])
			}
			End = p
		}
	} // while (Begin + 1 < End)
}

//***************************************************************************
// Top sorter: call recursive sorter for arrays of len >= 2.
//***************************************************************************
func QSort2(v []ValType) {
	if len(v) < 2 {
		return
	}
	qsort2r(v)
}
