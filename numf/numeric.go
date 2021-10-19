package numf

import (
	"math"
)

// Delta returns the delta between all consecutive floats
func Delta(floats []float64) []float64 {

	res := make([]float64, len(floats)-1)

	for i := 1; i < len(floats); i++ {
		res[i-1] = floats[i] - floats[i-1]
	}
	return res
}

// Compare returns maximum and minimum of two floats taking NaNs into account
func Compare(x float64, y float64) (max float64, min float64) {
	max = math.NaN()
	min = math.NaN()
	if !math.IsNaN(x) {
		max = x
		min = x
		if !math.IsNaN(y) {
			max = math.Max(max, y)
			min = math.Min(min, y)
		}
	} else {
		if !math.IsNaN(y) {
			max = y
			min = y
		}
	}
	return max, min
}

// FindIndex finds the index of first occurrence of the given value
func FindIndex(slice []float64, val float64) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// Insert inserts given value to given index into a slice
func Insert(slice []float64, idx int, val float64) []float64 {

	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = val
	return slice

}

// RemoveFrom removes an integer from given index
func RemoveFrom(slice []float64, s int) []float64 {
	return append(slice[:s], slice[s+1:]...)
}
