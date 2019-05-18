package utils

import "fmt"

/***********************************************************/
func CheckSorted(v []ValType) {
    ok := true
    ultimate := IndexType(len(v) - 2)
    for k := IndexType(0); k < ultimate; k++ {
        if v[k] > v[k+1] {
            fmt.Println("Error: v[", k, "]=", v[k], "v[", k+1, "]=", v[k+1])
            ok = false
        }
    }
    if ok {
        fmt.Println("hs1: OK")
    } else {
        fmt.Println("hs1: BAD")
    }
    //fmt.Println("C")
    //prHeap(v[:], 0, "")
    //fmt.Println(v)
}

