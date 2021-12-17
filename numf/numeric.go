/* Helper functions for handling floats.
 */
package numf

import (
	"math"
)

// Returns the delta between all consecutive floats. Returned slice length is one item shorter.
func Delta(slice []float64) []float64 {
	if len(slice) < 2 {
		return nil
	}
	res := make([]float64, len(slice)-1)

	for i := 1; i < len(slice); i++ {
		res[i-1] = slice[i] - slice[i-1]
	}
	return res
}

// Compares and returns maximum and minimum of two floats taking NaNs into account.
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

// Finds the index of first occurrence of the given value.
func FindIndex(slice []float64, val float64) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// Inserts given value to given index into a slice.
func Insert(slice []float64, idx int, val float64) []float64 {

	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = val
	return slice

}

// Removes an integer from given index
func RemoveFrom(slice []float64, s int) []float64 {
	return append(slice[:s], slice[s+1:]...)
}

// Checks if given float exists in the slice.
func Contains(slice []float64, s float64) bool {
	for _, a := range slice {
		if a == s {
			return true
		}
	}
	return false
}

// Creates a slice of given size filled with given value.
func SliceOf(value float64, size int) []float64 {
	if size <= 0 {
		return nil
	}
	s := make([]float64, size)
	for i := range s {
		s[i] = value
	}
	return s
}

// Calculates a slice of cumulative sum from given slice.
func Cumsum(slice []float64) []float64 {
	s := make([]float64, len(slice))
	var previous float64
	for i, v := range slice {
		s[i] = previous + v
		previous = s[i]
	}
	return s
}

// Multiplies two slices of same length element-wise.
func MulSlices(s1, s2 []float64) []float64 {
	if len(s1) != len(s2) {
		return nil
	}
	res := make([]float64, len(s1))
	for i, v1 := range s1 {
		res[i] = v1 * s2[i]
	}
	return res
}
