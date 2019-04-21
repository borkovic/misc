package main

import (
    "fmt"
    "time"
    "math/rand"
    "os"
    "strconv"
)
import (
    "../utils"
    hs "../heapsort"
)

type (
    Vec = hs.Vec
    CmpFunc = hs.CmpFunc
    Index = hs.Index
    Value = hs.Value
)

/***********************************************************/
func prHeap(v Vec, k Index, ident string) {
    fmt.Println(ident, v[k])
    last := hs.Len(v) - 1
    lCld := hs.LeftCld(k)
    rCld := hs.RightCld(k)
    if lCld <= last {
        prHeap(v, lCld, ident+"  ")
    }
    if rCld <= last {
        prHeap(v, rCld, ident+"  ")
    }
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

    v := make([]Value, N) // [N]Value
    for i := 0; i < len(v); i++ {
        v[i] = Value(rand.Int31n(int32(N)))
    }

    start := time.Now()
    hs.HeapsortF(v[:])
    elapsed := time.Since(start)
    //elapsed *= 1000.0
    fmt.Printf("GOF: Sorting [%s]%T: %T  %v seconds\n",
        utils.PrintLong(N), v[0], elapsed, elapsed.Seconds())

    utils.CheckSorted(v[:])
}
