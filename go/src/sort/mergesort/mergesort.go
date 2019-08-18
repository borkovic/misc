//***************************************************************************
package mergesort
import (
	"fmt"
)
import (
	"../utils"
)

//***************************************************************************
type ValType = utils.ValType
type IndexType = utils.IndexType


//***************************************************************************
// Recursive quicksort.
// 1. Partition into two parts. Note that pivot_high==pivot_low.
// 2. Recurse only on the shorter part to limit stack depth to log(N)
// 3. Continue looping with the longer part
// 4. Sorted array will and in v2
//***************************************************************************
func mergesort(v, v2 []ValType) {
	const debug bool = true

	Begin := IndexType(0)
	End := IndexType(len(v))
	if Begin + 1 >= End {
		return
	}
	if Begin + 2 == End {
		if v[0] > v[1] {
			v[0], v[1] = v[1], v[0]
		}
		return
	}

	mid := (Begin + End) / 2
	mergesort(v[Begin:mid], v2[Begin:mid])
	if debug {
		fmt.Println("Bot:", Begin, ",", mid)
		fmt.Println(v[Begin:mid])
	}
	mergesort(v[mid:End], v2[mid:End])
	if debug {
		fmt.Println("Top:", mid, ",", End)
		fmt.Println("> ",v[mid:End])
	}

	i := IndexType(Begin)
	j := IndexType(mid)
	x := IndexType(Begin)

	for i < mid && j < End {
		var s IndexType
		if v2[i] <= v2[j] {
			s = i
			i++
		} else {
			s = j
			j++
		}
		v[x] = v2[s]
		x++
	}
	for i < mid {
		v[x] = v2[i]
		x++
		i++
	}
	for j < End {
		v[x] = v2[j]
		x++
		j++
	}
	//copy(v2, v)
}

//***************************************************************************
// Top sorter: call recursive sorter for arrays of len >= 2.
//***************************************************************************
func Mergesort(v []ValType) {
	if len(v) < 2 {
		return
	}
	v2 := make([]ValType, len(v))
	mergesort(v, v2)
}
