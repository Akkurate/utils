package numi

import (
	"math/rand"
	"time"
)

func RandRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// SetDefault set default value for a parameter, if the configuration value is zero (=initialized empty variable)
func SetDefault(defaultval int, confval int) int {
	if confval == 0 {
		return defaultval
	}
	return confval
}
