package main

import (
    "fmt"
    "time"
    "math/rand"
    "os"
    "strconv"
)
import (
    hs "../heapsort"
)

type (
    Vec = hs.Vec
    CmpFunc = hs.CmpFunc
    Index = hs.Index
    Value = hs.Value
)

/***********************************************************/
func printLong(m int64) string {
    if m == 0 {
        return "0"
    }
    n := m
    if (m < 0) {
        n = -m
    }
    const base int64 = 10
    const digits = "0123456789"

    var buf [256]byte
    b := 0

    for n > 0 {
        d := n % base
        n = n / base
        buf[b] = digits[d]
        b++
    }
    buf[b] = 0
    numDigits := b
    b--

    // digits are reverse buf = [4201\0]
    // want buf2 = [1'024\0]
    var buf2 [256]byte
    b2 := 0
    if m < 0 {
        buf2[b2] = '-'
        b2++
    }

    for numDigits > 0 {
        buf2[b2] = buf[b]
        b2++
        b--
        numDigits--
        if numDigits > 0 && numDigits%3 == 0 {
            buf2[b2] = '\''
            b2++
        }
    }
    buf2[b2] = 0
    s := string(buf2[:])
    return s
}

/***********************************************************/
func checkSortedF(v Vec) {
    last := hs.Len(v) - 1
    ok := true
    for k := Index(0); k < last-1; k++ {
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

/***********************************************************/

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
        printLong(N), v[0], elapsed, elapsed.Seconds())

    checkSortedF(v[:])
}
