package numf

// NumRange creates a float slice with from start to end, with given step
func NumRange(start int, end int, step int) (numberrange []float64) {
	for x := start; x <= end; x = x + step {
		numberrange = append(numberrange, float64(x))
	}
	return numberrange
}

// SetDefault set default value for a parameter, if the configuration value is zero (=initialized empty variable)
func SetDefault(defaultval float64, confval float64) float64 {
	if confval == 0 {
		return defaultval
	}
	return confval
}
