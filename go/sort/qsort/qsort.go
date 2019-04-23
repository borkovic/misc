//***************************************************************************
package qsort

import "fmt"

import "../utils"

//***************************************************************************
type ValType = utils.ValType
type IndexType = utils.IndexType



//***************************************************************************
func gMedian(v []ValType) ValType {
    f := v[0]
    m := v[len(v)/2]
    l := v[len(v) - 1]
    if f < m {
        if m < l {
            return m
        } else if f < l {
            return l
        } else {
            return f
        }
    } else {
        if m > l {
            return m
        } else if f > l {
            return l
        } else {
            return f
        }
    }
}

//***************************************************************************
// Invariant:
// B      L-1  L         M-1  M   R-1  R      E-1  E
// [ x < p1 ]  [ p1<=x<=p2 ]  [ ??? ]  [ p2 < x ]
//***************************************************************************
func dnf2(v []ValType, pivotLow, pivotHigh ValType) (IndexType, IndexType) {
	if pivotLow > pivotHigh {
		panic("Lower pivot actually greater")
	}
	L := IndexType(0)
	M := L
	R := IndexType(len(v))

  	for M < R {
        if v[M] < pivotLow {
			tmp := v[L]
			v[L] = v[M]
			v[M] = tmp
   			L++
			M++
    	} else if pivotHigh < v[M] {
			tmp := v[M]
			v[M] = v[R-1]
			v[R-1] = tmp
  			R--
  		} else { // pivotLow <= v[M] <= pivotHigh
  			M++
  		}
  	}
	if L >= M {
		panic("Middle section must be non-empty")
	}
	return L, M
} // dnf2


//***************************************************************************
func qSort2r(v []ValType) {
    const debug bool = false

    if debug { fmt.Println("Q:", v) }
    Begin := IndexType(0)
    End := IndexType(len(v))
    if debug { fmt.Println(len(v)) }

    for Begin + 1 < End {
        if debug { fmt.Println("F be:", Begin, End) }
        if debug { fmt.Println("F v[be]:", v[Begin:End]) }
        medVal := gMedian(v[Begin:End])
        if debug { fmt.Println("F medVal:", medVal) }
        p, q := dnf2(v[Begin:End], medVal, medVal) // Must call with equal pivots
        if debug { fmt.Println("F pq:", p, q) }
        if debug {fmt.Println("F v[be]:", v[Begin:End]) }
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
    
        leftSize  := p - Begin
        rightSize := End - q
        if leftSize <= rightSize {
            if leftSize > 1 {
                qSort2r(v[Begin:p])
            }
            Begin = q
        } else {
            if rightSize > 1 {
                qSort2r(v[q:End])
            }
            End = p
        }
    } // while (Begin + 1 < End)
}

func QSort2(v []ValType) {
    if len(v) < 2 {
        return
    }
    qSort2r(v)
}

