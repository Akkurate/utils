package misc

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/Akkurate/utils/logging"
)

// SliceContains sliceContains
func SliceContains(slice []string, s string) bool {
	for _, a := range slice {
		if a == s {
			return true
		}
	}
	return false
}

// SliceNumberContains SliceNumberContains
func SliceNumberContains(slice []int, s int) bool {
	for _, a := range slice {
		if a == s {
			return true
		}
	}
	return false
}

func RandRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// GetEnvOrDefault
func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// TimeToString TimeToString
func TimeToString(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000000Z07")
}

// StringToTime StringToTime
func StringToTime(t string) time.Time {
	tm, _ := time.Parse("2006-01-02T15:04:05.000000Z07", t)
	return tm
}

// StringToInt StringToInt
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
		return 0
	}
	return f
}

// TimeToMicroseconds TimeToMicroseconds
func TimeToMicroseconds(tm time.Time) int64 {
	return tm.UnixNano() / int64(time.Microsecond)
}

// Use Use
func Use(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

// TimeMeasure GoPython
type TimeMeasure struct {
	start time.Time
}

// Print print
func (t *TimeMeasure) Print(msg string) {
	duration := time.Since(t.start)

	logging.Info(msg+" <yellow>%v ms</>", duration.Milliseconds())
}

// GetMilliseconds GetMilliseconds
func (t *TimeMeasure) GetMilliseconds() int64 {
	duration := time.Since(t.start)
	return int64(duration.Milliseconds())
}

// NewTimeMeasure NewTimeMeasure
func NewTimeMeasure() *TimeMeasure {
	x := &TimeMeasure{
		start: time.Now(),
	}
	return x
}

// PrettyByte PrettyByte
func PrettyByte(bytes []byte) string {
	b := len(bytes)
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

// Find Find
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// FindInt FindInt
func FindInt(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// CompareSliceInts CompareSliceInts
func CompareSliceInts(newslice []int, master []int) (idx []int) {

	for i, v := range master {
		_, b := FindInt(newslice, v)
		if !b {
			idx = append(idx, i)
		}
	}
	return idx
}

// InsertIntoSlice InsertIntoSlice
func InsertIntoSlice(slice []float64, idx int, val float64) []float64 {

	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = val
	return slice

}

// ParseTime ParseTime
func ParseTime(t string) time.Time {
	tm, _ := time.Parse("2006-01-02T15:04:05Z", t)
	return tm
}

// FileExists FileExists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// EnsureFolder EnsureFolder
func EnsureFolder(root string) string {
	newpath := filepath.Join(root)
	os.MkdirAll(newpath, os.ModePerm)
	return newpath
}

// MatchAll MatchAll
// example: Avant(?P<DeviceID>\d{1,})(_S(?P<SlaveID>\d{1,})_C(?P<CellID>\d{1,}))?
func MatchAll(a string, url string) (paramsMap map[string]string) {

	compRegEx := regexp.MustCompile(a)

	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}

// NanosecondsToTime NanosecondsToTime
func NanosecondsToTime(tm int64) time.Time {
	return time.Unix(0, tm*int64(time.Nanosecond))
}
