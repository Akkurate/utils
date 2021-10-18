package numi

//Unique returns unique ints from given slice of ints
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

// Delta returns the delta between all consecutive ints
func Delta(ints []int) []int {

	res := make([]int, len(ints))

	for i := 1; i < len(ints); i++ {
		res[i-1] = ints[i] - ints[i-1]
	}
	return res
}

// Max returns bigger int from two ints
func Max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

// Min returns smaller int from two ints
func Min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

// FindIndex finds the index of first occurrence of the given value
func FindIndex(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// Contains Check if the input slice contains given integer
func Contains(slice []int, s int) bool {
	for _, a := range slice {
		if a == s {
			return true
		}
	}
	return false
}

// FindMissingIndexes compares <newslice> against <master>. Those numbers' indexes from <master> are collected whic are not found from <newslice>
func FindMissingIndexes(newslice []int, master []int) (idx []int) {

	for i, v := range master {
		_, b := FindIndex(newslice, v)
		if !b {
			idx = append(idx, i)
		}
	}
	return idx
}
