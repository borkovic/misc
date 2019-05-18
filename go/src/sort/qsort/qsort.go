//***************************************************************************
package qsort

import "../utils"

//***************************************************************************
type ValType = utils.ValType
type IndexType = utils.IndexType

//***************************************************************************
func median3(v []ValType) ValType {
	f := v[0]
	m := v[len(v)/2]
	l := v[len(v)-1]
	if f < m { // fm
		if m < l { // fm,ml => fml
			return m
		} else if f < l { // fm,lm,fl => flm
			return l
		} else { // fm,lm,lf => lfm
			return f
		}
	} else { // mf
		if m > l { // mf,lm => lmf
			return m
		} else if f > l { // mf,ml,lf => mlf
			return l
		} else { // mf,ml,fl => mfl
			return f
		}
	}
}

//***************************************************************************
// Invariant:
// B     L-1  L    M-1  M   R-1  R     E-1  E
// [ x < p ]  [ p==x ]  [ ??? ]  [ p < x ]
//***************************************************************************
func dnf1(v []ValType, pivot ValType) (IndexType, IndexType) {
	L := IndexType(0)
	M := L
	R := IndexType(len(v))

	for M < R { // ??? section non-empty
		if v[M] < pivot {
			v[L], v[M] = v[M], v[L]
			L++
			M++
		} else if pivot < v[M] {
			v[M], v[R-1] = v[R-1], v[M]
			R--
		} else { // pivot == v[M]
			M++
		}
	}
	if L >= M {
		panic("Middle section must be non-empty since pivot is from the array")
	}
	return L, M
} // dnf1

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

	for M < R { // ??? section non-empty
		if v[M] < pivotLow {
			v[L], v[M] = v[M], v[L]
			L++
			M++
		} else if pivotHigh < v[M] {
			v[M], v[R-1] = v[R-1], v[M]
			R--
		} else { // pivotLow <= v[M] <= pivotHigh
			M++
		}
	}
	if L >= M {
		panic("Middle section must be non-empty since pivot is from the array")
	}
	return L, M
} // dnf2

//***************************************************************************
func qsort1r(v []ValType) {
	Begin := IndexType(0)
	End := IndexType(len(v))

	for Begin+1 < End {
		medVal := median3(v[Begin:End])
		var p, q IndexType
		if false {
			p, q = dnf2(v[Begin:End], medVal, medVal) // Must call with equal pivots
		} else {
			p, q = dnf1(v[Begin:End], medVal) // Must call with equal pivots
		}
		p += Begin // p,q are relative to Begin
		q += Begin // so need to adjust
		if p < Begin {
			panic("bad p")
		}
		if q > End {
			panic("bad q")
		}
		//  j in [Begin, p) =>  v[j] < pivot1
		//  j in [p, q)     =>  pivot1 <= v[j] <= pivot2
		//  j in [q, End)   =>  pivot2 < v[j]

		leftSize := p - Begin
		rightSize := End - q
		if leftSize <= rightSize {
			if leftSize > 1 {
				qsort1r(v[Begin:p])
			}
			Begin = q
		} else {
			if rightSize > 1 {
				qsort1r(v[q:End])
			}
			End = p
		}
	} // while (Begin + 1 < End)
}

//***************************************************************************
func QSort(v []ValType) {
	if len(v) < 2 {
		return
	}
	qsort1r(v)
}
