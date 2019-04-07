package main

import (
    "fmt"
    "time"
    "math/rand"
    "os"
    "strconv"
    "sort"
)

import (
    hs "../heapsort"
)
type (
    Vec = sort.IntSlice
    Index = hs.Index
)

/***********************************************************/
func checkSortedS(v []int) {
    last := len(v) - 1
    ok := true
    for k := 0; k < last-1; k++ {
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

/***********************************************************/
func main() {
    /*
    const N = 1 * 1000 * 1000
    const N = 10
    const N = 100*1000*1000
    const N = 10 * 1000 * 1000
    */
    ts := os.Args[1]
    N, err := strconv.ParseInt(ts, 10, 64)
    if err != nil {
        panic(err)
    }

    v := make(Vec, N)
    for i := 0; i < len(v); i++ {
        v[i] = int(rand.Int31n(int32(N)))
    }

    start := time.Now()
    v.Sort() //sort.Ints(v)
    elapsed := time.Since(start)
    //elapsed *= 1000.0
    fmt.Printf("GOS: Sorting [%s]%T: %T  %v seconds\n",
        printLong(N), v[0], elapsed, elapsed.Seconds())

    checkSortedS(v[:])
}
