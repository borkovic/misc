#include <vector>
#include <iostream>

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

Index lChild(Index k) {
    return 2*k+1;
}

Index rChild(Index k) {
    return 2*k+2;
}


using CmpFunc = int (*)(Value l, Value r);


/* Move element k towards root if it small
 */
void toRoot(Vec& v, Index k, CmpFunc cmp) {
    while (k > 0) {
        auto p = parent(k);
        //fmt.Println("TR: ", v)
        //fmt.Println("TR: k=",k, "v[k]=",v[k], "p=",p, "v[p]=",v[p])
        if (cmp(v[p], v[k]) <= 0) {
            break;
        }
        auto t = v[p];
        v[p] = v[k];
        v[k] = t;
        k = p;
    }
}

/* Move element k toward leaves if it is large
 */
void toLeaves(Vec& v, Index k, Index last, CmpFunc cmp) {
    for (auto leftChild = lChild(k); leftChild <= last;  leftChild = lChild(k)) { // k has at least one child
        auto smallChild = leftChild;
        auto rightChild = leftChild + 1;
        if (rightChild <= last && cmp(v[rightChild], v[smallChild]) < 0) {
            smallChild = rightChild;
        }
        //fmt.Println(v, k, v[k], smallChild, v[smallChild])
        if (cmp(v[smallChild], v[k]) >= 0) {
            break;
        }
        auto t = v[smallChild];
        v[smallChild] = v[k];
        v[k] = t;

        k = smallChild;
    }
}

/* Make heap with elem[0] being root, smallest in heap
 */
void heapify(Vec& v, CmpFunc cmp) {
    auto last = Len(v)-1;
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
    auto last = Len(v)-1;
    for (auto k = last; k >= 1; k--) {
        auto t = v[0];
        v[0] = v[k];
        v[k] = t;
        toLeaves(v, 0, k-1, cmp);
    }
}

void checkSorted(const Vec& v, CmpFunc cmp) {
    bool ok = true;
    for (auto k = Index(0); k < last-1; k++) {
        if (cmp(v[k], v[k+1]) < 0) {
            std::cout <<
            "Error: v[" << k << "]=" << v[k] << "v[" << k+1 << "]=" << v[k+1];
            ok = false;
        }
    }
    if (ok) {
        std::cout << "OK\n";
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
    auto last = Len(v)-1;
    auto leftChild = lChild(k);
    auto rightChild = rChild(k);
    if  (leftChild <= last) {
        prHeap(v, leftChild, ident+"  ");
    }
    if  (rightChild <= last) {
        prHeap(v, rightChild, ident+"  ");
    }
}

int main() {
    constexpr long N = 100*1000*1000;
    Vec v;
    v.resize(N);
    for (long i = 0; i < Len(v); i++) {
        v[i] = Value(rand.Int31n(N))
    }
    heapsort(v, CmpGT);
    return 0;
}

