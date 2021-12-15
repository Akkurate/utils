/* Helper functions for handling strings.
 */
package str

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// Converts string to SnakeCase.
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.Title(strings.ToLower(snake))
}

// Returns unique strings from given slice of strings.
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

// Finds the index of first occurrence of the given value and returns the index and a boolean. Boolean is true when index is found, otherwise index = -1.
func FindIndex(slice []string, s string) (int, bool) {
	for i, a := range slice {
		if a == s {
			return i, true
		}
	}
	return -1, false
}

// Checks if given string exists in the slice.
func Contains(slice []string, s string) bool {
	for _, a := range slice {
		if a == s {
			return true
		}
	}
	return false
}

// Checks if given string exists in the slice. Case insensitive.
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

// Removes the first occurence of given string from a slice and returns the result as new slice
func Remove(slice []string, s string) []string {
	c := make([]string, len(slice))
	copy(c, slice)
	for i, v := range c {
		if v == s {
			c = append(c[:i], c[i+1:]...)
			break
		}
	}
	return slice
}

// Inserts given string to given index into a slice and returns the result as new slice
func Insert(slice []string, idx int, val string) []string {
	c := make([]string, len(slice))
	copy(c, slice)
	c = append(c, "")
	copy(c[idx+1:], c[idx:])
	c[idx] = val
	return c

}

// Removes a string from given index and returns the result as new slice
func RemoveFrom(slice []string, s int) []string {
	c := make([]string, len(slice))
	copy(c, slice)
	return append(c[:s], c[s+1:]...)
}

func CleanUp(str string) string {
	// trim str
	str = strings.TrimSpace(str)
	// remove extra spaces using regex
	str = strings.Trim(str, "")
	//str = regexp.MustCompile(`(\n|\t|)`).ReplaceAllString(str, " ")
	str = regexp.MustCompile(`(\s+)`).ReplaceAllString(str, " ")
	return str

}
