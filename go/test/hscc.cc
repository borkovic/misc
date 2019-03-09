#include <vector>
#include <iostream>
#include <chrono> 

#include "pl.h"

using namespace std::chrono; 

typedef int Index; //using Index = int;
using Value = int;
using Vec = std::vector<Value>;

Index Len(const Vec& v) {
    return Index(v.size());
}


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


using CmpFunc = int (*)(Value l, Value r);


/* Move element k towards root if it small
 */
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
        //fmt.Println(v, k, v[k], smlCld, v[smlCld])
        if (cmp(v[smlCld], val) >= 0) {
            break;
        }
        v[k] = v[smlCld];
        k = smlCld;
    }
    v[k] = val;
}

/* Make heap with elem[0] being root, smallest in heap
 */
void heapify(Vec& v, CmpFunc cmp) {
    const auto last = Len(v)-1;
    for (auto k = parent(last); k >= 0; k--) {
        toLeaves(v, k, last, cmp);
    }
}


/* Heapsort in descending order
 */
void heapsort(Vec& v, CmpFunc cmp) {
    // make heap in linear time
    //fmt.Println("A"); prHeap(v[:], 0, "")
    //fmt.Println(v)
    heapify(v, cmp);
    //fmt.Println("B"); prHeap(v[:], 0, "")
    const auto last = Len(v)-1;
    for (auto k = last; k >= 1; k--) {
        std::swap(v[0], v[k]);
        toLeaves(v, 0, k-1, cmp);
    }
}

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

/*
 * Compare Greater Than
 */
int CmpGT(Value l, Value r) {
    return CmpLT(r, l);
}



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

int main(int argc, char* argv[]) {
    //constexpr const long N = 10*1000*1000;
    const long N = atol(argv[1]);
    Vec v;
    v.resize(N);
    for (long i = 0; i < Len(v); i++) {
        auto r = std::rand();
        r = r < 0 ? -r : r;
        v[i] = Value(r % N);
    }
    const auto cmp = CmpGT;

    const auto start = high_resolution_clock::now(); 
    heapsort(v, cmp);
    const auto stop = high_resolution_clock::now(); 
    const auto duration = duration_cast<seconds>(stop - start); 

    char buf[256];
    printLong(N, &buf);

    std::cout << "CC: Sorting int[" << buf << "]: " << duration.count() << " seconds\n";

    checkSorted(v, cmp);
    return 0;
}

  
  
// To get the value of duration use the count() 
// member function on the duration object 
