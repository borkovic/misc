package vc

//**********************************************
import (
	//"fmt"
	"strconv"
)

//**********************************************
// compile const for size of VC, could be runtime, but less efficient
const (
	VcSize EngineIdx = 12
)

//**********************************************
// range of VC
type VcVal int32

//**********************************************
// VC, hides individual values
type VC struct {
	timestamps [VcSize]VcVal
}

//**********************************************
// increment one dimension of VC
func (vc *VC) Incr(idx EngineIdx, v VcVal) {
	vc.timestamps[idx] += v
}

//**********************************************
// comparison for equality, all dimensions equal
func (vc1 *VC) Equal(vc2 *VC) bool {
	for i, _ := range vc1.timestamps {
		if vc1.timestamps[i] != vc2.timestamps[i] {
			return false
		}
	}
	return true
}

//**********************************************
// comparison for <  -- all dims <=, at least one dim <
func (vc1 *VC) Less(vc2 *VC) bool {
	hasLess := false
	for i, _ := range vc1.timestamps {
		if vc1.timestamps[i] > vc2.timestamps[i] {
			return false
		} else if vc1.timestamps[i] < vc2.timestamps[i] {
			hasLess = true
		}
	}
	return hasLess
}

//**********************************************
// comparison for <= -- all dims <=
func (vc1 *VC) LessEq(vc2 *VC) bool {
	for i, _ := range vc1.timestamps {
		if vc1.timestamps[i] > vc2.timestamps[i] {
			return false
		}
	}
	return true
}

//**********************************************
func (vc1 *VC) NotEq(vc2 *VC) bool {
	return !vc2.Equal(vc1)
}

//**********************************************
func (vc1 *VC) Greater(vc2 *VC) bool {
	return vc2.Less(vc1)
}

//**********************************************
func (vc1 *VC) GreaterEq(vc2 *VC) bool {
	return vc2.LessEq(vc1)
}

//**********************************************
func (vc1 *VC) Concurrent(vc2 *VC) bool {
	less := false
	greater := false
	for i, _ := range vc1.timestamps {
		if vc1.timestamps[i] < vc2.timestamps[i] {
			less = true
		} else if vc1.timestamps[i] > vc2.timestamps[i] {
			greater = true
		}
	}
	return less && greater
}

//**********************************************
func (vc1 *VC) Maximize(vc2 *VC) {
	if len(vc1.timestamps) != len(vc2.timestamps) {
		panic("VCsw with different lenghts")
	}
	numEngs := EngineIdx(len(vc1.timestamps))
	for engIdx := numEngs * 0; engIdx < numEngs; engIdx++ {
		if vc1.val(engIdx) < vc2.val(engIdx) {
			vc1.timestamps[engIdx] = vc2.val(engIdx)
		}
	}
}

//**********************************************
// string representation of VC
func (vc *VC) String() string {
	s := "["
	first := true
	for _, ts := range vc.timestamps {
		ns := strconv.FormatInt(int64(ts), 10)
		if first {
			first = false
			s += ns
		} else {
			s += "," + ns
		}
	}
	return s + "]"
}

//**********************************************
func (vc1 *VC) val(idx EngineIdx) VcVal {
	return vc1.timestamps[idx]
}
