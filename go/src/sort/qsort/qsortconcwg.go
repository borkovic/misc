//***************************************************************************
package qsort

import (
	"fmt"
	"sync"
)

//***************************************************************************
// Concurrent single recursive quicksort.
// 1. Partition into two parts. Note that pivot_high==pivot_low.
// 2. Concurrent Recurse only on the shorter part to limit stack depth to log(N)
// 3. Continue looping with the longer part
//***************************************************************************
func qsort2rcwg(v []ValType, wg *sync.WaitGroup) {
	const debug bool = false
	const (
		goLim int = 1000
	)

	if debug {
		fmt.Println("Q:", v)
	}
	Begin := IndexType(0)
	End := IndexType(len(v))
	if debug {
		fmt.Println(len(v))
	}
	var wg2 sync.WaitGroup

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
				if wg != nil && goLim < leftSize {
					wg2.Add(1)
					go qsort2rcwg(v[Begin:p], &wg2)
				} else {
					qsort2rcwg(v[Begin:p], nil)
				}
			}
			Begin = q
		} else {
			if rightSize > 1 {
				if wg != nil && goLim < rightSize {
					wg2.Add(1)
					go qsort2rcwg(v[q:End], &wg2)
				} else {
					qsort2rcwg(v[q:End], nil)
				}
			}
			End = p
		}
	} // while (Begin + 1 < End)

	wg2.Wait()

	if wg != nil {
		wg.Done()
	}
}

//***************************************************************************
// Top sorter: call recursive sorter for arrays of len >= 2.
//***************************************************************************
func QSort2cwg(v []ValType) {
	if len(v) < 2 {
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	qsort2rcwg(v, &wg)
	wg.Wait()
}
