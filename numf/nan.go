package numf

import (
	"math"
	"strings"

	"github.com/Akkurate/utils/numi"
)

// Checks whether the given float is a valid number instead of NaN or Inf.
func IsValid(value float64) bool {
	if math.IsNaN(value) {
		return false
	}
	if math.IsInf(value, 0) {
		return false
	}
	return true
}

// Replaces given value with a new one, if given value is NaN or Inf.
func ReplaceNan(value float64, replacewith float64) float64 {
	if IsValid(value) {
		return value
	}
	return replacewith
}

// Replaces all NaNs in a slice with given value.
func ReplaceNans(values []float64, replacewith float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		if !IsValid(v) {
			res[i] = replacewith
		} else {
			res[i] = v
		}
	}
	return res
}

// Creates a slice of NaNs of given size.
func NanSlice(size int) (nanslice []float64) {
	nanslice = make([]float64, size)
	for i := range nanslice {
		nanslice[i] = math.NaN()
	}
	return nanslice
}

// Drops NaNs from a slice.
func DropNan(slice []float64) []float64 {
	res := make([]float64, len(slice))
	i := 0
	for _, v := range slice {
		if IsValid(v) {
			res[i] = v
			i++
		}
	}
	return res[:i]
}

// Drops NaNs from a slice in place. This is more memory efficient way than using DropNan.
func DropNanInplace(slice *[]float64) {
	i := 0
	for _, v := range *slice {
		if IsValid(v) {
			(*slice)[i] = v
			i++
		}
	}
	*slice = (*slice)[:i]
}

// Fills NaN values with a value based on given method:
//  "previous" // fills the NaNs with previous value
//  "linear"   // fills the NaNs with linear interpolation
// Filling starts from first valid value, thus leaving any preceding NaNs untouched. By setting prefill = true, first valid value is used to replace also the preceding NaNs.
//
// One sample's lifetime can be set by setting validTime value to > 0. Filling is then performed for validTime samples. Using validTime overrides prefill = true -setting
func FillNan(slice []float64, method string, prefill bool, validTime int) []float64 {
	if len(slice) == 0 {
		return slice
	}
	method = strings.ToLower(method)
	res := make([]float64, len(slice))
	firstvalue := slice[0]
	dropped := DropNan(slice)
	validTime = numi.Select(validTime == 0, len(slice), validTime)
	vt := -1

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
				vt++
				if i == 0 {
					res[i] = firstvalue
				} else if vt >= validTime {
					res[i] = v
				} else {
					res[i] = res[i-1]
				}
			} else {
				res[i] = v
				vt = 0
			}
		}

	case "linear":

		var idx []int
		var val []float64

		res = slice
		res[0] = firstvalue

		for i, v := range res {
			if !math.IsNaN(v) {
				idx = append(idx, i)
				val = append(val, v)
			}
		}
		idx = append(idx, len(res))
		val = append(val, val[len(val)-1])
		res = append(res, math.NaN())

		for i := 0; i < len(val)-1; i++ {
			vt = 0
			gapsize := float64(idx[i+1] - idx[i])
			gapdelta := val[i+1] - val[i]

			if gapsize > 1 {
				linstep := gapdelta / gapsize
				for j := idx[i]; j < idx[i+1]; j++ {
					vt++
					if vt >= validTime {
						res[j+1] = math.NaN()
					} else {
						res[j+1] = res[j] + linstep
					}
				}
			}
		}
		res = res[:len(res)-1]
	default:
		return res

	}

	return res
}
