package numf

// Returns an evenly spaced slice of floats, between <start> and <end> with spacing <step>.
func NumRange(start int, end int, step int) (numberrange []float64) {
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
