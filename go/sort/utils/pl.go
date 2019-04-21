package utils

/***********************************************************/
func PrintLong(m int64) string {
    if m == 0 {
        return "0"
    }
    var positive bool = true
    var n uint64
    if (m < 0) {
        positive = false
        m += 1 // to avoid abs(int64_min) == 1 + abs(int64_max)
        m = -m
        n = uint64(m)
        n -= 1
    } else {
        n = uint64(m)
    }
    const base uint64 = 10
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
    if !positive {
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

