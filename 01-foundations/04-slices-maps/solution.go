package slicesmaps

import "slices"

// AppendUnique appends v to slice only if v is not already present (linear search).
// Returns the new slice (may share backing array with s per append rules).
func AppendUnique(s []int, v int) []int {
	if slices.Contains(s,v){
		return s
	}
	return append(s, v)
}

// SafeSlice returns a sub-slice of s from low to high (exclusive).
// If high > len(s), high is clamped to len(s). If low > high after clamping, return empty slice (not nil).
// If low < 0, low is treated as 0.
func SafeSlice(s []int, low, high int) []int {
	if high>len(s){
		high=len(s)
	}
	if low<0{
		low=0
	}
	if low>high{
		return []int{}
	}
	return s[low:high]
}

// MapInvert swaps keys and values. If duplicate values exist, the last key wins.
func MapInvert(m map[string]int) map[int]string {
	inverted:=make(map[int]string,len(m))
	for k,v:=range m{
		inverted[v]=k
	}
	return inverted
}

// CapAfterAppend demonstrates understanding of len/cap: given s with len/cap,
// append one element and return the new cap. Used to verify you understand append growth.
func CapAfterAppend(s []int) int {
	s = append(s, 7)
	return cap(s)
}
