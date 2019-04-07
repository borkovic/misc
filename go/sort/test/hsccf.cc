#include <vector>
#include <iostream>
#include <chrono> 

#include "heap.h"
#include "pl.h"

using namespace std::chrono; 


/***********************************************************/


/***********************************************************/
/* Move element k towards root if it small
 */
#if 0
static void
toRoot(Vec& v, Index k) {
    const auto val = v[k];
    while (k > 0) {
        const auto p = parent(k);
        //fmt.Println("TR: ", v)
        //fmt.Println("TR: k=",k, "v[k]=",v[k], "p=",p, "v[p]=",v[p])
        if (v[p] >= val) { // cmp
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
static void
toLeaves(Vec& v, Index k, Index last) {
    const auto val = v[k];
    for (auto lCld = leftCld(k); lCld <= last;  lCld = leftCld(k)) { // k has at least one child
        auto smlCld = lCld;
        const auto rCld = lCld + 1;
        if (rCld <= last && v[rCld] >  v[smlCld]) { // cmp
            smlCld = rCld;
        }
        if (v[smlCld] <= val) { // cmp
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
static void
heapify(Vec& v) {
    const auto last = Len(v)-1;
    for (auto k = parent(last); k >= 0; k--) {
        toLeaves(v, k, last);
    }
}


/***********************************************************/
/* Heapsort in descending order
 */
void heapsort(Vec& v) {
    // make heap in linear time
    heapify(v);
    const auto last = Len(v)-1;
    for (auto k = last; k >= 1; k--) {
        std::swap(v[0], v[k]);
        toLeaves(v, 0, k-1);
    }
}

/***********************************************************/
void checkSorted(const Vec& v) {
    const Index last = v.size() - 1;
    bool ok = true;
    for (auto k = Index(0); k < last-1; k++) {
        if (v[k] > v[k+1]) {
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
static void
prHeap(const Vec& v, Index k, const std::string& ident) {
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

/***********************************************************/
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

    const auto start = high_resolution_clock::now(); 
    heapsort(v);
    const auto stop = high_resolution_clock::now(); 
    const auto duration = duration_cast<seconds>(stop - start); 

    char buf[256];
    printLong(N, &buf);

    std::cout << "CC hsortF: Sorting int[" << buf << "]: " << duration.count() << " seconds\n";

    checkSorted(v);
    return 0;
}

  
  
// To get the value of duration use the count() 
// member function on the duration object 
