package numf

import (
	"math"
	"strings"
)

// IsValid checks whether the given float is a valid number instead of NaN or Inf
func IsValid(value float64) bool {
	if math.IsNaN(value) {
		return false
	}
	if math.IsInf(value, 0) {
		return false
	}
	return true
}

// ReplaceNan
// If <value> is NaN or Inf, replaces <value> with <replacewith>
func ReplaceNan(value float64, replacewith float64) float64 {
	if math.IsNaN(value) {
		return replacewith
	}
	if math.IsInf(value, 0) {
		return replacewith
	}
	return value
}

// ReplaceNanSlice Replace NaNs in a slice with given value
func ReplaceNans(values []float64, replacewith float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		if math.IsNaN(v) {
			res[i] = replacewith
		} else {
			res[i] = v
		}
	}
	return res
}

// NaNSlice Create a slice of NaNs of given size
func NanSlice(size int) (nanslice []float64) {
	nanslice = make([]float64, size)
	for i := range nanslice {
		nanslice[i] = math.NaN()
	}
	return nanslice
}

// Dropna drop NaNs from a slice
func DropNan(slice []float64) []float64 {

	res := make([]float64, len(slice))
	i := 0
	for _, v := range slice {
		if !math.IsNaN(v) {
			res[i] = v
			i++
		}
	}
	return res[:i]
}

// Fillna
// Fills NaN values with a value base on method.
// method "previous" fills with previous value,method "linear" fills the NaNs with linear interpolation
// If no previous value found for filling, first valid value can be used by setting prefill=true
// validTime (higher number than 1) sets the sample's lifetime. If the sample "dies", NaN is written.  This overrides prefill = true -setting
func FillNan(slice []float64, method string, prefill bool, validTime int) []float64 {

	method = strings.ToLower(method)
	res := make([]float64, len(slice))
	firstvalue := math.NaN()
	dropped := DropNan(slice)

	if len(dropped) == 0 {
		return slice
	}

	if prefill {
		firstvalue = dropped[0]
	}

	switch method {

	case "previous":

		for i, v := range slice {
			if math.IsNaN(v) {
				if i == 0 {
					res[i] = firstvalue
				} else {
					res[i] = res[i-1]
				}
			} else {
				res[i] = v
			}
		}

	case "linear":

		var idx []int
		var val []float64

		linearresult := slice

		for i, v := range slice {
			if !math.IsNaN(v) {
				idx = append(idx, i)
				val = append(val, v)
			}
		}

		for i := 0; i < len(val)-1; i++ {

			gap := idx[i+1] - idx[i]
			gapval := val[i+1] - val[i]

			if gap > 1 {
				step := gapval / float64(gap)
				for j := idx[i]; j < idx[i+1]; j++ {
					linearresult[j+1] = linearresult[j] + step
				}
			}
		}
		res = FillNan(linearresult, "previous", prefill, validTime)

	}
	if validTime > 1 {
		filterRes := NanSlice(len(slice))
		for i, v := range slice {
			if !math.IsNaN(v) {
				for j := i; j < i+validTime; j++ {
					if j < len(slice) {
						filterRes[j] = v
					}
				}
			}
		}
		for i, v := range filterRes {
			if math.IsNaN(v) {
				res[i] = v
			}
		}
	}
	return res
}
