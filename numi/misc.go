
package numi

import (
	"math/rand"
	"time"
)

// Returns an evenly spaced slice of integers, between <start> and <end> with spacing <step>.
func NumRange(start int, end int, step int) (numberrange []int) {
	for x := start; x <= end; x = x + step {
		numberrange = append(numberrange, x)
	}
	return numberrange
}

// Creates a random integer between min and max value.
func RandRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// Replaces default value for a parameter, if the configuration value is zero (ie. initialized empty variable).
// Note that the function does not make any difference, if the configuration is set to zero on purpose.
func SetDefault(defaultval int, confval int) int {
	if confval == 0 {
		return defaultval
	}
	return confval
}
