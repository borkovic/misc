
namespace hs {

typedef int Index; //using Index = int;
using Value = int;
using Vec = std::vector<Value>;


Index Len(const Vec& v) {
    return Index(v.size());
}


/***********************************************************/
/*
    1,2->0
    3,4->1
    5,6->2
* Parent, Left and right child in array based heap
*/
inline Index parent(Index k) {
    return (k-1) / 2;
}

inline Index leftCld(Index k) {
    return 2*k+1;
}

inline Index rightCld(Index k) {
    return 2*k+2;
}

}

