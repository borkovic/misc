extern "C" {
extern int printf(const char* fmt, ...);
}

#include "printwithcommas.h"

using namespace Util;

void
f()
{
    char buf[200];

    {
        signed char cs[] = {
            1 << (8*sizeof(cs[0])-1),
            -2, -1, 0, 1, 2,
            0x7f,
        };

        for (unsigned k = 0; k < sizeof(cs)/sizeof(cs[0]); ++k) {
            PrintWithCommas(buf, cs[k]);
            printf("%hhd   %s\n", cs[k], buf);
        }
    }

    {
        short ss[] = {
            1 << (8*sizeof(ss[0])-1),
            -2, -1, 0, 1, 2,
            -20003, 20003,
            0x7fff,
        };
        for (unsigned k = 0; k < sizeof(ss)/sizeof(ss[0]); ++k) {
            PrintWithCommas(buf, ss[k]);
            printf("%hd   %s\n", ss[k], buf);
        }
    }

    {
        int is[] = {
            1 << (8*sizeof(is[0])-1),
            -2, -1, 0, 1, 2,
            -20003, 20003,
            ~( 1 << (8*sizeof(is[0])-1) ),
        };
        for (unsigned k = 0; k < sizeof(is)/sizeof(is[0]); ++k) {
            PrintWithCommas(buf, is[k]);
            printf("%d   %s\n", is[k], buf);
        }
    }

    {
        long ls[] = {
            1L << (8*sizeof(ls[0])-1),
            -2, -1, 0, 1, 2,
            -20003, 20003,
            ~( 1L << (8*sizeof(ls[0])-1) ),
        };
        for (unsigned k = 0; k < sizeof(ls)/sizeof(ls[0]); ++k) {
            PrintWithCommas(buf, ls[k]);
            printf("%ld   %s\n", ls[k], buf);
        }
    }



    {
        unsigned char ucs[] = {
            0,
            1, 2, 3,
            23,
            0xff,
        };
        for (unsigned k = 0; k < sizeof(ucs)/sizeof(ucs[0]); ++k) {
            PrintWithCommas(buf, ucs[k]);
            printf("%hhu   %s\n", ucs[k], buf);
        }
    }

    {
        unsigned short uss[] = {
            0,
            1, 2, 3,
            20003,
            ~0,
        };
        for (unsigned k = 0; k < sizeof(uss)/sizeof(uss[0]); ++k) {
            PrintWithCommas(buf, uss[k]);
            printf("%hu   %s\n", uss[k], buf);
        }
    }

    {
        unsigned int uis[]= {
            0,
            1, 2, 3,
            20003,
            ~0,
        };
        for (unsigned k = 0; k < sizeof(uis)/sizeof(uis[0]); ++k) {
            PrintWithCommas(buf, uis[k]);
            printf("%u   %s\n", uis[k], buf);
        }
    }

    {
        unsigned long ul = 0;
        ul = ~ul;
        PrintWithCommas(buf, ul);
        printf("%lu   %s\n", ul, buf);
    }
}

int main()
{
    f();
    return 0;
}

