extern "C" {
    extern int strlen(const char* s);
    extern int sprintf(char* buf, const char* fmt, ...);
}
#include "printwithcommas.h"

namespace Util {

int
PrintWithCommas(char* buf, const unsigned long ul)
{
    if (0UL == ul) {
        buf[0] = '0';
        buf[1] = '\0';
        return 1;
    }
    return PrintWithCommas2(buf, ul);
    return PrintWithCommasRec(buf, ul);
}

int
PrintWithCommasRec(char* buf, const unsigned long ul)
{
    if (0UL == ul) {
        return 0;
    }

    if (ul < 1000UL) {
        sprintf(buf, "%lu", ul);
        return strlen(buf);
    } else {
        const unsigned long th = ul / 1000UL;
        int nchars = 0;
        if (th > 0UL) {
            nchars = PrintWithCommasRec(buf, th);
        }
        int nc;
        if (nchars > 0) {
            sprintf(buf + nchars, ",%03lu", ul % 1000UL);
            nc = 4;
        } else {
            sprintf(buf + nchars, "%03lu", ul % 1000UL);
            nc = 3;
        }
        return nchars + nc;
    }
}

int
PrintWithCommas2(char* buf, const unsigned long unsLong)
{
    static const char digits[] = "0123456789";
    char* p = buf;
    unsigned long ul = unsLong;

    while (ul > 0UL) {
        const unsigned long q = ul / 1000UL;
        const unsigned long r = ul % 1000UL;

        const unsigned long r0 = r % 10UL;
        const unsigned long r1 = (ul % 100UL - r0) / 10UL;
        const unsigned long r2 = r / 100UL;

        if (q > 0) {
            *p++ = digits[r0];
            *p++ = digits[r1];
            *p++ = digits[r2];
            *p++ = ',';
        } else {
            *p++ = digits[r0];
            if (r1 > 0UL || r2 > 0UL) {
                *p++ = digits[r1];
            }
            if (r2 > 0UL) {
                *p++ = digits[r2];
            }
        }

        ul = q;
    }

    *p = '\0';
    const int n = p - buf;
    char* q = p - 1;
    p = buf;
    // revert buf
    while (p < q) {
        const char c = *p;
        *p = *q;
        *q = c;
        ++p; --q;
    }
    return n;
}

}


