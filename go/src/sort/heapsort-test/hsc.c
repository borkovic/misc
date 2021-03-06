#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <time.h>

#include "pl.h"

typedef int Index; //using Index = int;
typedef int Value;
typedef enum Bool { False = 0, True = 1 } Bool;







/***********************************************************/
/*
    1,2->0
    3,4->1
    5,6->2
* Parent, Left and right child in array based heap
*/
Index parent(Index k) {
    return (k-1) / 2;
}

Index leftCld(Index k) {
    return 2*k+1;
}

Index rightCld(Index k) {
    return 2*k+2;
}


/***********************************************************/
typedef int (*CmpFunc)(Value l, Value r);


/***********************************************************/
/* Move element k towards root if it small
 */
void toRoot(Value* v, Index k, CmpFunc cmp) {
    const Value val = v[k];
    while (k > 0) {
        const Index p = parent(k);
        //fmt.Println("TR: ", v)
        //fmt.Println("TR: k=",k, "v[k]=",v[k], "p=",p, "v[p]=",v[p])
        if (cmp(v[p], val) <= 0) {
            break;
        }
        v[k] = v[p];
        k = p;
    }
    v[k] = val;
}

/***********************************************************/
/* Move element k toward leaves if it is large
 */
void toLeaves(Value* v, Index k, Index last, CmpFunc cmp) {
    const Value val = v[k];
    for (Index lCld = leftCld(k); lCld <= last;  lCld = leftCld(k)) { // k has at least one child
        Index smlCld = lCld;
        const Index rCld = lCld + 1;
        if (rCld <= last && cmp(v[rCld], v[smlCld]) < 0) {
            smlCld = rCld;
        }
        if (cmp(v[smlCld], val) >= 0) {
            break;
        }
        v[k] = v[smlCld];
        k = smlCld;
    }
    v[k] = val;
}

/***********************************************************/
/* Make heap with elem[0] being root, smallest in heap
 */
void heapify(Value* v, Index sz, CmpFunc cmp) {
    const Index last = sz-1;
    for (Index k = parent(last); k >= 0; k--) {
        toLeaves(v, k, last, cmp);
    }
}


/***********************************************************/
/* Heapsort in descending order
 */
void Heapsort(Value* v, Index sz, CmpFunc cmp) {
    // make heap in linear time
    heapify(v, sz, cmp);
    const Index last = sz-1;
    for (Index k = last; k >= 1; k--) {
        Value t = v[0]; v[0] = v[k]; v[k] = t;
        toLeaves(v, 0, k-1, cmp);
    }
}

/***********************************************************/
void checkSorted(const Value* v, Index sz, CmpFunc cmp) {
    const Index last = sz - 1;
    Bool ok = True;
    for (Index k = (Index)(0); k < last-1; k++) {
        if (cmp(v[k], v[k+1]) < 0) {
            printf("Error: v[%d]=%d, v[%d]=%d\n", k, v[k], k+1, v[k+1]);
            ok = False;
        }
    }
    if (ok) {
        printf("hs3: OK\n");
    } else {
        printf("hs3: BAD\n");
    }
    //fmt.Println("C"); prHeap(v[:], 0, 0)
    //fmt.Println(v)
}


/***********************************************************/
/*
 * Compare Less Than
 */
int CmpLT(Value l, Value r) {
    if (l < r) {
        return -1;
    } else if (l > r) {
        return 1;
    } else {
        return 0;
    }
}

/***********************************************************/
/*
 * Compare Greater Than
 */
int CmpGT(Value l, Value r) {
    return CmpLT(r, l);
}

/***********************************************************/
void prHeap(const Value* v, Index sz, Index k, int ident) {
    char buf[512];

    memset(buf, ' ', 2*ident);
    const int nc = sprintf(buf+2*ident, "%d", v[k]);
    buf[2*ident+nc] = '\0';
    printf("%s\n", buf);

    const Index last = sz-1;
    const Index lCld = leftCld(k);
    const Index rCld = rightCld(k);
    if  (lCld <= last) {
        prHeap(v, sz, lCld, ident+1);
    }
    if  (rCld <= last) {
        prHeap(v, sz, rCld, ident+1);
    }
}

/***********************************************************/
int main(int argc, char* argv[]) {
    //const long N = 10*1000*1000;
    const long N = atol(argv[1]);

    Value* const v = (Value*)(malloc(N*sizeof(v[0])));
    for (long i = 0; i < N; i++) {
        Value r = rand();
        r = r < 0 ? -r : r;
        v[i] = (Value)(r % N);
    }
    const CmpFunc cmp = CmpGT;
    clock_t start, end;
     
    start = clock();
    Heapsort(v, N, cmp);
    end = clock();

    double cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
    char buf[256];
    printLong(N, &buf);
    printf("C: Sorting int[%s]: %.2f seconds\n", buf, cpu_time_used);

    checkSorted(v, N, cmp);
    return 0;
}


