<<<<<<< HEAD
package heapsort

// if cmp is cmpLT, sort will be descending
// if cmp is cmpGT, sort will be ascending

/***********************************************************/
/* Move element k towards root if smaller than descendants
 */
func toRoot(v Vec, k Index, cmp CmpFunc) {
	val := v[k]
	for k > 0 {
		p := parent(k)
		if cmp(v[p], val) <= 0 {
			break
		}
		v[k] = v[p]
		k = p
	}
	v[k] = val
}

/***********************************************************/
/* Move element k toward leaves if it is large
 */
func toLeaves(v Vec, k Index, last Index, cmp CmpFunc) {
	val := v[k]
	for lCld := LeftCld(k); lCld <= last; lCld = LeftCld(k) { // k has at least one child
		smlCld := lCld
		rCld := lCld + 1
		if rCld <= last && cmp(v[rCld], v[smlCld]) < 0 {
			smlCld = rCld
		}
		if cmp(v[smlCld], val) >= 0 {
			break
		}
		v[k] = v[smlCld]
		k = smlCld
	}
	v[k] = val
}

/***********************************************************/
/* Make heap with elem[0] being root, smallest in heap
 */
func heapify(v Vec, cmp CmpFunc) {
	last := Len(v) - 1
	for k := parent(last); k >= 0; k-- {
		toLeaves(v, k, last, cmp)
	}
}

/***********************************************************/
/* Heapsort in descending order
 */
func Heapsort(v Vec, cmp CmpFunc) {
	// make heap in linear time
	heapify(v, cmp)
	last := Len(v) - 1
	for k := last; k >= 1; k-- {
		v[0], v[k] = v[k], v[0]
		toLeaves(v, 0, k-1, cmp)
	}
}
||||||| merged common ancestors
=======
package heapsort

// if cmp is cmpLT, sort will be descending
// if cmp is cmpGT, sort will be ascending

/***********************************************************/
/* Move element k towards root if smaller than descendants
 */
func toRoot(v Vec, k Index, cmp CmpFunc) {
	val := v[k]
	for k > 0 {
		p := parent(k)
		//fmt.Println("TR: ", v)
		//fmt.Println("TR: k=",k, "v[k]=",v[k], "p=",p, "v[p]=",v[p])
		if cmp(v[p], val) <= 0 {
			break
		}
		v[k] = v[p]
		k = p
	}
	v[k] = val
}

/***********************************************************/
/* Move element k toward leaves if it is large
 */
func toLeaves(v Vec, k Index, last Index, cmp CmpFunc) {
	val := v[k]
	for lCld := LeftCld(k); lCld <= last; lCld = LeftCld(k) { // k has at least one child
		smlCld := lCld
		rCld := lCld + 1
		if rCld <= last && cmp(v[rCld], v[smlCld]) < 0 {
			smlCld = rCld
		}
		//fmt.Println(v, k, v[k], smlCld, v[smlCld])
		if cmp(v[smlCld], val) >= 0 {
			break
		}
		v[k] = v[smlCld]
		k = smlCld
	}
	v[k] = val
}

/***********************************************************/
/* Make heap with elem[0] being root, smallest in heap
 */
func heapify(v Vec, cmp CmpFunc) {
	last := Len(v) - 1
	for k := parent(last); k >= 0; k-- {
		toLeaves(v, k, last, cmp)
	}
}

/***********************************************************/
/* Heapsort in descending order
 */
func Heapsort(v Vec, cmp CmpFunc) {
	// make heap in linear time
	//fmt.Println("A"); prHeap(v[:], 0, "")
	//fmt.Println(v)
	heapify(v, cmp)
	//fmt.Println("B"); prHeap(v[:], 0, "")
	last := Len(v) - 1
	for k := last; k >= 1; k-- {
		v[0], v[k] = v[k], v[0]
		toLeaves(v, 0, k-1, cmp)
	}
}
>>>>>>> origin/master
