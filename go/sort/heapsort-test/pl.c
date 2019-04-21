#include "pl.h"

/***********************************************************/
void printLong(const long m, char (*buf2)[256]) {
    const long base = 10;
    static const char digits[] = "0123456789";

    char buf[sizeof(*buf2)/sizeof((*buf2)[0])];
    char* p = buf;
    if (0 == m) {
        buf[0] = digits[m];
        buf[1] = '\0';
        return;
    }

    long n = (m >= 0 ? m : -m);

    while (n > 0) {
        long d = n % base;
        n = n / base;
        *p = digits[d];
        p++;
    }
    *p = '\0';

    // suppose m == -1024
    // digits are reverse buf = "4201\0"
    // want buf2 = "-1'024\0"
    int numDigits = p - buf;

    char* p2 = &(*buf2)[0];
    if (m < 0) {
        *p2 = '-';
        p2++;
    }
    --p;

    while (numDigits > 0) {
        *p2 = *p;

        p2++;
        p--;
        numDigits--;
        if (numDigits > 0 && (numDigits%3 == 0)) {
            *p2 = '\'';
            p2++;
        }
    }
    *p2 = '\0';
}

