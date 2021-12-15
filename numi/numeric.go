/* Helper functions for handling integers.
 */
package numi

// Returns unique integers from given slice.
func Unique(t []int) []int {

	var res []int
	unique := make(map[int]bool)

	for _, tt := range t {
		if !unique[tt] {
			unique[tt] = true
			res = append(res, tt)
		}
	}

	return res
}

// Returns the delta between all consecutive ints. Returned slice length is one item shorter.
func Delta(ints []int) []int {

	res := make([]int, len(ints))

	for i := 1; i < len(ints); i++ {
		res[i-1] = ints[i] - ints[i-1]
	}
	return res
}

// Returns the larger of given integers.
func Max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

// Returns the smaller of given integers.
func Min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

// Finds the index of first occurrence of the given value and returns the index and a boolean. Boolean is true when index is found, otherwise index = -1.
func FindIndex(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// Checks if given integer exists in the slice.
func Contains(slice []int, s int) bool {
	for _, a := range slice {
		if a == s {
			return true
		}
	}
	return false
}

// Compares <newslice> against <slice>. Those numbers' indexes from <slice> are collected which are not found from <newslice>.
func FindMissingIndexes(newslice []int, slice []int) (idx []int) {

	for i, v := range slice {
		_, b := FindIndex(newslice, v)
		if !b {
			idx = append(idx, i)
		}
	}
	return idx
}

// RemoveFrom removes an integer from given index and returns the result as a new slice.
func RemoveFrom(slice []int, s int) []int {
	c := make([]int, len(slice))
	copy(c, slice)
	return append(c[:s], c[s+1:]...)
}

// Insert inserts given value to given index into a slice and returns the result as a new slice.
func Insert(slice []int, idx int, val int) []int {
	c := make([]int, len(slice))
	copy(c, slice)
	c = append(c, 0)
	copy(c[idx+1:], c[idx:])
	c[idx] = val
	return c

}

// Removes first occurence of given integer from a slice and returns the result as a new slice.
func Remove(slice []int, s int) []int {
	c := make([]int, len(slice))
	copy(c, slice)
	for i, v := range c {
		if v == s {
			c = append(c[:i], c[i+1:]...)
			break
		}
	}
	return slice
}

// SliceOf creates a slice of given size filled with value.
func SliceOf(value int, size int) []int {
	s := make([]int, size)
	for i := range s {
		s[i] = value
	}
	return s
}

// Calculates a slice of cumulative sum from given slice.
func Cumsum(slice []int) []int {
	s := make([]int, len(slice))
	var previous int
	for i, v := range slice {
		s[i] = previous + v
		previous = s[i]
	}
	return s
}
