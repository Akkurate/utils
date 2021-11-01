package system

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/Akkurate/utils/numf"
)

func GetSysmemory() float64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return numf.RoundTo(float64(m.Sys)/1024/1024, 5)
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

// GetEnvOrDefault
func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// isInSlice first argument is the slice, second is the value
// using reflect to check if the value is in the slice
// verify that the length of args should be equal to 2
func isInSlice(args ...interface{}) bool {
	if len(args) != 2 {
		return false
	}
	slice := reflect.ValueOf(args[0])
	value := reflect.ValueOf(args[1])
	for i := 0; i < slice.Len(); i++ {
		if slice.Index(i).Interface() == value.Interface() {
			return true
		}
	}
	return false
}
