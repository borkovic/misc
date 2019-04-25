package heapsort

/***********************************************************/
/* Move element k towards root if smaller than descendants
 */
func toRootF(v Vec, k Index) {
    val := v[k]
    for k > 0 {
        p := parent(k)
        //fmt.Println("TR: ", v)
        //fmt.Println("TR: k=",k, "v[k]=",v[k], "p=",p, "v[p]=",v[p])
        if (v[p] >= val) { // cmp
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
func toLeavesF(v Vec, k Index, last Index) {
    val := v[k]
    for lCld := LeftCld(k); lCld <= last; lCld = LeftCld(k) { // k has at least one child
        smlCld := lCld
        rCld := lCld + 1
        if rCld <= last && v[rCld] > v[smlCld] { // cmp
            smlCld = rCld
        }
        //fmt.Println(v, k, v[k], smlCld, v[smlCld])
        if v[smlCld] <= val { // cmp
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
func heapifyF(v Vec) {
    last := Len(v) - 1
    for k := parent(last); k >= 0; k-- {
        toLeavesF(v, k, last)
    }
}

/***********************************************************/
/* Heapsort in descending order
 */
func HeapsortF(v Vec) {
    // make heap in linear time
    //fmt.Println("A"); prHeap(v[:], 0, "")
    //fmt.Println(v)
    heapifyF(v)
    //fmt.Println("B"); prHeap(v[:], 0, "")
    last := Len(v) - 1
    for k := last; k >= 1; k-- {
        v[0], v[k] = v[k], v[0]
        toLeavesF(v, 0, k-1)
    }
}

