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
// Recursive mergesort`.
// v2 is aux array used for merging
//***************************************************************************
func mergesort(v, v2 []ValType) {
	const debug bool = false

	Begin := IndexType(0)
	End := IndexType(len(v))
	if End-Begin <= 1 {
		return
	}
	if End-Begin <= 2 {
		if v[0] > v[1] {
			v[0], v[1] = v[1], v[0]
		}
		return
	}

	mid := (Begin + End) / 2
	mergesort(v[Begin:mid], v2[Begin:mid])
	if debug {
		fmt.Println("Bot:", Begin, ",", mid)
		fmt.Println("> ", v[Begin:mid])
	}
	mergesort(v[mid:End], v2[mid:End])
	if debug {
		fmt.Println("Top:", mid, ",", End)
		fmt.Println("> ", v[mid:End])
	}

	i := Begin
	j := mid
	x := Begin

	for i < mid && j < End {
		if v[i] <= v[j] {
			v2[x] = v[i]
			i++
		} else {
			v2[x] = v[j]
			j++
		}
		x++
	}
	for i < mid {
		v2[x] = v[i]
		x++
		i++
	}
	for j < End {
		v2[x] = v[j]
		x++
		j++
	}
	copy(v, v2)
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
