package heapsort

/***********************************************************/
type (
    Index int
    Value int32
    Vec []Value
    CmpFunc func(l Value, r Value) int
)


/***********************************************************/
func Len(v Vec) Index {
    return Index(len(v))
}




/***********************************************************/
/*
Parent, Left and right child in array based heap (first index = 0)
Heap of 6 elements: 
    0
 1     2
3 4   5

left,right -> parent
     1,2   ->   0
     3,4   ->   1
       5   ->   2

left(n) = 2*n+1
right(n) = left(n)+1
parent(n) = floor((n-1)/2)
*/

/***********************************************************/
func parent(k Index) Index {
    if k <= 0 {
        panic("Negative index in parent")
    }
    return (k - 1) / 2
}

/***********************************************************/
func LeftCld(k Index) Index {
    return 2*k + 1
}

/***********************************************************/
func RightCld(k Index) Index {
    return 2*k + 2
}

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
