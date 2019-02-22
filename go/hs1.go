package main

import "fmt"
import "math/rand"

type Index int
type Value int32
type Vec []Value

func Len(v Vec) Index {
    return Index(len(v))
}

/*
            0
         1     2
        3 4   5 6
    children->parent
    1,2     ->0
    3,4     ->1
    5,6     ->2
* Parent, Left and right child in array based heap
*/
func parent(k Index) Index {
    if k <= 0 {
        panic("Negative index in parent")
    }
    return (k - 1) / 2
}

func lChild(k Index) Index {
    return 2*k + 1
}

func rChild(k Index) Index {
    return 2*k + 2
}

type CmpFunc func(l Value, r Value) int

/* Move element k towards root if it small
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

/* Move element k toward leaves if it is large
 */
func toLeaves(v Vec, k Index, last Index, cmp CmpFunc) {
    val := v[k]
    for leftChild := lChild(k); leftChild <= last; leftChild = lChild(k) { // k has at least one child
        smallChild := leftChild
        rightChild := leftChild + 1
        if rightChild <= last && cmp(v[rightChild], v[smallChild]) < 0 {
            smallChild = rightChild
        }
        //fmt.Println(v, k, v[k], smallChild, v[smallChild])
        if cmp(v[smallChild], val) >= 0 {
            break
        }
        v[k] = v[smallChild]
        k = smallChild
    }
    v[k] = val
}

/* Make heap with elem[0] being root, smallest in heap
 */
func heapify(v Vec, cmp CmpFunc) {
    last := Len(v) - 1
    for k := parent(last); k >= 0; k-- {
        toLeaves(v, k, last, cmp)
    }
}

/* Heapsort in descending order
 */
func heapsort(v Vec, cmp CmpFunc) {
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

func checkSorted(v Vec, cmp CmpFunc) {
    last := Len(v) - 1
    ok := true
    for k := Index(0); k < last-1; k++ {
        if cmp(v[k], v[k+1]) < 0 {
            fmt.Println("Error: v[", k, "]=", v[k], "v[", k+1, "]=", v[k+1])
            ok = false
        }
    }
    if ok {
        fmt.Println("OK")
    }
    //fmt.Println("C"); prHeap(v[:], 0, "")
    //fmt.Println(v)
}

/*
 * Compare Less Than
 */
func CmpLT(l Value, r Value) int {
    if l < r {
        return -1
    } else if l > r {
        return 1
    } else {
        return 0
    }
}

/*
 * Compare Greater Than
 */
func CmpGT(l Value, r Value) int {
    return CmpLT(r, l)
}

func prHeap(v Vec, k Index, ident string) {
    fmt.Println(ident, v[k])
    last := Len(v) - 1
    leftChild := lChild(k)
    rightChild := rChild(k)
    if leftChild <= last {
        prHeap(v, leftChild, ident+"  ")
    }
    if rightChild <= last {
        prHeap(v, rightChild, ident+"  ")
    }
}

func main() {
    const N = 10 * 1000 * 1000
    //const N = 100*1000*1000
    //const N = 10

    var v [N]Value
    for i := 0; i < len(v); i++ {
        v[i] = Value(rand.Int31n(N))
    }
    cmp := CmpGT
    heapsort(v[:], cmp)
    checkSorted(v[:], cmp)
}
