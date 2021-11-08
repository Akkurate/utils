/* Helper functions for handling conversions between types.
 */
package conv

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// Converts time.Time object to string. Time format:
//  "2006-01-02T15:04:05.000000Z07"
func TimeToString(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000000Z07")
}

// Converts string to time.Time object. Time format:
//  "2006-01-02T15:04:05.000000Z07"
func StringToTime(t string) time.Time {
	tm, _ := time.Parse("2006-01-02T15:04:05.000000Z07", t)
	return tm
}

// Converts string to int. Returns 0 if error detected.
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// Converts float64 to int. Returns NaN if error detected.
func StringToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return math.NaN()
	}
	return f
}

// Converts any input type to string represanttion.
func Use(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
