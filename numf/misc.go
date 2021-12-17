package numf

import (
	"fmt"
)

// Returns an evenly spaced slice of floats, between <start> and <end> with spacing <step>.
func NumRange(start int, end int, step int) (numberrange []float64) {
	if step <= 0 || end < start {
		return numberrange
	}
	for x := start; x <= end; x = x + step {
		numberrange = append(numberrange, float64(x))
	}
	return numberrange
}

// Replaces default value for a parameter, if the configuration value is zero (ie. initialized empty variable).
// Note that the function does not make any difference, if the configuration is set to zero on purpose.
func SetDefault(defaultval float64, confval float64) float64 {
	if confval == 0 {
		return defaultval
	}
	return confval
}

// Returns the given values without changes, if condition == FALSE.
// Returns the values in switched order, if condition == TRUE.
func SwitchIf(condition bool, i1, i2 float64) (float64, float64) {
	if condition {
		return i2, i1
	}
	return i1, i2
}
// Returns i1 if condition == TRUE, else returns i2
func Select(condition bool, i1, i2 float64) float64 {
	if condition {
		return i1
	}
	return i2
}

// Checks if two slices are equal, element-wise. NaN and Inf are also handled.
func IsEqualSlice(s1, s2 []float64) bool {

	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if IsValid(v) {
			if v != s2[i] {
				return false
			}
		} else {
			if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", s2[i]) {
				return false
			}

		}
	}
	return true
}
// Checks if two values have same sign. Zero sign is ambiguous and is considered to be both positive and negative.
func IsSameSign(x float64, y float64) bool {
	return (x >= 0 && y >= 0) || (x <= 0 && y <= 0)
}
