package heapsort
import "../utils"

/***********************************************************/
type (
    Index = utils.IndexType
    Value = utils.ValueType

    Vec []Value
    CmpFunc func(l Value, r Value) int
)


/***********************************************************/
func Len(v Vec) Index {
    return Index(len(v))
}




/***********************************************************/
/*
Parent, Left and right child in array based heap (first index = 0)
Heap of 6 elements: 
    0
 1     2
3 4   5

left,right -> parent
     1,2   ->   0
     3,4   ->   1
       5   ->   2

left(n) = 2*n+1
right(n) = left(n)+1
parent(n) = floor((n-1)/2)
*/

/***********************************************************/
func parent(k Index) Index {
    if k <= 0 {
        panic("Negative index in parent")
    }
    return (k - 1) / 2
}

/***********************************************************/
func LeftCld(k Index) Index {
    return 2*k + 1
}

/***********************************************************/
func RightCld(k Index) Index {
    return 2*k + 2
}

/***********************************************************/
/*
 * Compare Less Than
 */
func CmpLT(l Value, r Value) int {
    if l < r {
        return -1
    } else if l > r {
        return 1
    } else {
        return 0
    }
}

/***********************************************************/
/*
 * Compare Greater Than
 */
func CmpGT(l Value, r Value) int {
    if l > r {
        return -1
    } else if l < r {
        return 1
    } else {
        return 0
    }
}

