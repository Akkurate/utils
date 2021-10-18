package system

import (
	"os"
	"path/filepath"
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
