#include <vector>
#include <iostream>
#include <chrono> 

#include "heap.h"
#include "pl.h"

using namespace std::chrono; 

namespace hs {

typedef int Index; //using Index = int;
using Value = int;
using Vec = std::vector<Value>;



/***********************************************************/
using CmpFunc = int (*)(Value l, Value r);


/***********************************************************/
/* Move element k towards root if it small
 */
#if 0
void toRoot(Vec& v, Index k, CmpFunc cmp) {
    const auto val = v[k];
    while (k > 0) {
        const auto p = parent(k);
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
#endif

/***********************************************************/
/* Move element k toward leaves if it is large
 */
void toLeaves(Vec& v, Index k, Index last, CmpFunc cmp) {
    const auto val = v[k];
    for (auto lCld = leftCld(k); lCld <= last;  lCld = leftCld(k)) { // k has at least one child
        auto smlCld = lCld;
        const auto rCld = lCld + 1;
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
void heapify(Vec& v, CmpFunc cmp) {
    const auto last = Len(v)-1;
    for (auto k = parent(last); k >= 0; k--) {
        toLeaves(v, k, last, cmp);
    }
}


/***********************************************************/
/* Heapsort in descending order
 */
void heapsort(Vec& v, CmpFunc cmp) {
    // make heap in linear time
    heapify(v, cmp);
    const auto last = Len(v)-1;
    for (auto k = last; k >= 1; k--) {
        std::swap(v[0], v[k]);
        toLeaves(v, 0, k-1, cmp);
    }
}

/***********************************************************/
void checkSorted(const Vec& v, CmpFunc cmp) {
    const Index last = v.size() - 1;
    bool ok = true;
    for (auto k = Index(0); k < last-1; k++) {
        if (cmp(v[k], v[k+1]) < 0) {
            std::cout <<
            "Error: v[" << k << "]=" << v[k] << "v[" << k+1 << "]=" << v[k+1];
            ok = false;
        }
    }
    if (ok) {
        std::cout << "hs2: OK\n";
    } else {
        std::cout << "hs2: BAD\n";
    }
    //fmt.Println("C"); prHeap(v[:], 0, "")
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
void prHeap(const Vec& v, Index k, const std::string& ident) {
    std::cout << ident << " " << v[k] << "\n";
    const auto last = Len(v)-1;
    const auto lCld = leftCld(k);
    const auto rCld = rightCld(k);
    if  (lCld <= last) {
        prHeap(v, lCld, ident+"  ");
    }
    if  (rCld <= last) {
        prHeap(v, rCld, ident+"  ");
    }
}
}




/***********************************************************/
int main(int argc, char* argv[]) {
    //constexpr const long N = 10*1000*1000;
    const long N = atol(argv[1]);
    //Vec v;
    hs::Vec v;

    v.resize(N);
    for (long i = 0; i < hs::Len(v); i++) {
        auto r = std::rand();
        r = r < 0 ? -r : r;
        v[i] = hs::Value(r % N);
    }
    const auto cmp = hs::CmpGT;

    const auto start = high_resolution_clock::now(); 
    hs::heapsort(v, cmp);
    const auto stop = high_resolution_clock::now(); 
    const auto duration = duration_cast<seconds>(stop - start); 

    char buf[256];
    printLong(N, &buf);

    std::cout << "CC hsort: Sorting int[" << buf << "]: " << duration.count() << " seconds\n";

    hs::checkSorted(v, cmp);
    return 0;
}

