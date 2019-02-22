package main

import "fmt"
import "math/rand"

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
        fmt.Println("hs1: OK")
    } else {
        fmt.Println("hs1: BAD")
    }
    //fmt.Println("C"); prHeap(v[:], 0, "")
    //fmt.Println(v)
}

/*
 * Compare Less Than
 */
func cmpLT(l Value, r Value) int {
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
func cmpGT(l Value, r Value) int {
    return cmpLT(r, l)
}

func prHeap(v Vec, k Index, ident string) {
    fmt.Println(ident, v[k])
    last := Len(v) - 1
    lCld := leftCld(k)
    rCld := rightCld(k)
    if lCld <= last {
        prHeap(v, lCld, ident+"  ")
    }
    if rCld <= last {
        prHeap(v, rCld, ident+"  ")
    }
}

func main() {
    //const N = 1 * 1000 * 1000
    const N = 10 * 1000 * 1000
    //const N = 100*1000*1000
    //const N = 10

    var v [N]Value
    for i := 0; i < len(v); i++ {
        v[i] = Value(rand.Int31n(N))
    }
    cmp := cmpGT
    Heapsort(v[:], cmp)
    checkSorted(v[:], cmp)
}
