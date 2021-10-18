package str

import "strings"

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

// Contains sliceContains
func Contains(slice []string, s string) bool {
	for _, a := range slice {
		if a == s {
			return true
		}
	}
	return false
}

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


