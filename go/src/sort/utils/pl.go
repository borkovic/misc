package utils

/***********************************************************/
func PrintUlong(m uint64) string {
    if m == 0 {
        return "0"
    }
    const base uint64 = 10
    const digits = "0123456789"
	const squote byte = '\''

	n := m
    var buf [256]byte
    b := 0

    for n > 0 {
        d := n % base
        buf[b] = digits[d]
        b++
        n /= base
    }
    buf[b] = 0
    numDigits := b
    b--

    // digits are reversed in buf = [7654321\0]
    // want buf2 = [1'234'567\0]
    var buf2 [256]byte
    b2 := 0

    for numDigits > 0 {
        buf2[b2] = buf[b]
        b2++
        b--
        numDigits--
        if numDigits > 0 && numDigits%3 == 0 {
            buf2[b2] = squote
            b2++
        }
    }
    buf2[b2] = 0
    s := string(buf2[:])
    return s
}

/***********************************************************/
func PrintLong(m int64) string {
	if m >= 0 {
		return PrintUlong(uint64(m))
	}
    m += 1 // to avoid abs(int64_min) == 1 + abs(int64_max)
    m = -m
    n := uint64(m)
    n += 1
	return "-" + PrintUlong(uint64(m))
}

