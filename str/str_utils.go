package str

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.Title(strings.ToLower(snake))
}

//Unique returns unique strings from given slice of strings
func Unique(t []string) []string {

	var res []string
	unique := make(map[string]bool)

	for _, tt := range t {
		if !unique[tt] {
			unique[tt] = true
			res = append(res, tt)
		}
	}

	return res
}

// FindIndex returns the first occurrence index of given string. Returns -1 if string was not found.
func FindIndex(slice []string, s string) (int, bool) {
	for i, a := range slice {
		if a == s {
			return i, true
		}
	}
	return -1, false
}

// Contains checks if given string exists in the slice
func Contains(slice []string, s string) bool {
	for _, a := range slice {
		if a == s {
			return true
		}
	}
	return false
}

// ContainsIgnorecase checks if given string exists in the slice. Case insensitive.
func ContainsIgnorecase(slice []string, s string) bool {
	s = strings.ToLower(s)

	for _, a := range slice {
		a = strings.ToLower(a)
		if a == s {
			return true
		}
	}
	return false
}

// Remove removes given string from a slice
func Remove(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return slice
}

// Insert inserts given value to given index into a slice
func Insert(slice []string, idx int, val string) []string {

	slice = append(slice, "")
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = val
	return slice

}

// RemoveFrom removes a string from given index
func RemoveFrom(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
