//***************************************************************************
package qsort

import (
	"fmt"
)

const (
	goLim int = 1000
)

//***************************************************************************
// Concurrent single recursive quicksort.
// 1. Partition into two parts. Note that pivot_high==pivot_low.
// 2. Concurrent Recurse only on the shorter part to limit stack depth to log(N)
// 3. Continue looping with the longer part
//***************************************************************************
func qsort2rc(v []ValType, Ret chan<- struct{}) {
	const debug bool = false

	if debug {
		fmt.Println("Q:", v)
	}
	Begin := IndexType(0)
	End := IndexType(len(v))
	if debug {
		fmt.Println(len(v))
	}
	numGos := 0
	var ret chan struct{} = nil
	if Ret != nil {
		ret = make(chan struct{}, 20)
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
				if ret != nil && goLim < leftSize {
					numGos++
					go qsort2rc(v[Begin:p], ret)
				} else {
					qsort2rc(v[Begin:p], nil)
				}
			}
			Begin = q
		} else {
			if rightSize > 1 {
				if ret != nil && goLim < rightSize {
					numGos++
					go qsort2rc(v[q:End], ret)
				} else {
					qsort2rc(v[q:End], nil)
				}
			}
			End = p
		}
	} // while (Begin + 1 < End)

	if ret != nil {
		for i := 0; i < numGos; i++ {
			<-ret
		}
		if ret != nil {
			close(ret)
		}
	}
	if Ret != nil {
		Ret <- struct{}{}
	}
}

//***************************************************************************
// Top sorter: call recursive sorter for arrays of len >= 2.
//***************************************************************************
func QSort2c(v []ValType) {
	if len(v) < 2 {
		return
	}
	ret := make(chan struct{}, 2)
	qsort2rc(v, ret)
	<-ret
	close(ret)
}
