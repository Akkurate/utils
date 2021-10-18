package conv

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// TimeToString TimeToString
func TimeToString(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000000Z07")
}

// StringToTime StringToTime
func StringToTime(t string) time.Time {
	tm, _ := time.Parse("2006-01-02T15:04:05.000000Z07", t)
	return tm
}

// StringToInt Convert string to int, returns 0 if error detected
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
func StringToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return math.NaN()
	}
	return f
}

// Use Use anything as string
func Use(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
