package misc

import (
	"fmt"
	"math"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// ReplaceNan Replaces NaN and Inf with given value
func ReplaceNan(value float64, replacewith float64) float64 {
	if math.IsNaN(value) {
		return replacewith
	}
	if math.IsInf(value, 0) {
		return replacewith
	}
	return value
}

// IsValidFloat checks whether the given float is a valid number instead of NaN or Inf
func IsValidFloat(value float64) bool {
	if math.IsNaN(value) {
		return false
	}
	if math.IsInf(value, 0) {
		return false
	}
	return true
}

// ReplaceNanSlice Replace NaNs in a slice with given value
func ReplaceNanSlice(values []float64, replacewith float64) []float64 {
	res := make([]float64, len(values))
	for i, v := range values {
		if math.IsNaN(v) {
			res[i] = replacewith
		} else {
			res[i] = v
		}
	}
	return res
}

// NaNSlice Create a slice of NaNs of given size
func NanSlice(size int) (nanslice []float64) {
	nanslice = make([]float64, size)
	for i := range nanslice {
		nanslice[i] = math.NaN()
	}
	return nanslice
}

// Numberrange creates a slice with numbers from start to end, with given step (integer)
func Numberrange(start int, end int, step int) (numberrange []float64) {
	for x := start; x <= end; x = x + step {
		numberrange = append(numberrange, float64(x))
	}
	return numberrange
}

//UniqueStrings returns unique strings from given slice of strings
func UniqueStrings(t []string) []string {

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

//UniqueInts returns unique ints from given slice of ints
func UniqueInts(t []int) []int {

	var res []int
	unique := make(map[int]bool)

	for _, tt := range t {
		if !unique[tt] {
			unique[tt] = true
			res = append(res, tt)
		}
	}

	return res
}

// Dropna drop NaNs from a slice
func Dropna(slice []float64) []float64 {

	res := make([]float64, len(slice))
	i := 0
	for _, v := range slice {
		if !math.IsNaN(v) {
			res[i] = v
			i++
		}
	}
	return res[:i]
}

// Fillna
// Fills NaN values with a value base on method.
// method "previous" fills with previous value,method "linear" fills the NaNs with linear interpolation
// If no previous value found for filling, first valid value can be used by setting prefill=true
// validTime (higher number than 1) sets the sample's lifetime. If the sample "dies", NaN is written.  This overrides prefill = true -setting
func Fillna(slice []float64, method string, prefill bool, validTime int) []float64 {

	method = strings.ToLower(method)
	res := make([]float64, len(slice))
	firstvalue := math.NaN()
	dropped := Dropna(slice)

	if len(dropped) == 0 {
		return slice
	}

	if prefill {
		firstvalue = dropped[0]
	}

	switch method {

	case "previous":

		for i, v := range slice {
			if math.IsNaN(v) {
				if i == 0 {
					res[i] = firstvalue
				} else {
					res[i] = res[i-1]
				}
			} else {
				res[i] = v
			}
		}

	case "linear":

		var idx []int
		var val []float64

		linearresult := slice

		for i, v := range slice {
			if !math.IsNaN(v) {
				idx = append(idx, i)
				val = append(val, v)
			}
		}

		for i := 0; i < len(val)-1; i++ {

			gap := idx[i+1] - idx[i]
			gapval := val[i+1] - val[i]

			if gap > 1 {
				step := gapval / float64(gap)
				for j := idx[i]; j < idx[i+1]; j++ {
					linearresult[j+1] = linearresult[j] + step
				}
			}
		}
		res = Fillna(linearresult, "previous", prefill, validTime)

	}
	if validTime > 1 {
		filterRes := NanSlice(len(slice))
		for i, v := range slice {
			if !math.IsNaN(v) {
				for j := i; j < i+validTime; j++ {
					if j < len(slice) {
						filterRes[j] = v
					}
				}
			}
		}
		for i, v := range filterRes {
			if math.IsNaN(v) {
				res[i] = v
			}
		}
	}
	return res
}

// DeltaI returns delta between consecutive ints in a slice
func DeltaI(ints []int) []int {

	res := make([]int, len(ints))

	for i := 1; i < len(ints); i++ {
		res[i-1] = ints[i] - ints[i-1]
	}
	return res
}

// DeltaF returns delta between consecutive floats in a slice
func DeltaF(floats []float64) []float64 {

	res := make([]float64, len(floats)-1)

	for i := 1; i < len(floats); i++ {
		res[i-1] = floats[i] - floats[i-1]
	}
	return res
}

// CompareFloats returns maximum and minimum of two floats taking NaNs into account
func CompareFloats(x float64, y float64) (max float64, min float64) {
	max = math.NaN()
	min = math.NaN()
	if !math.IsNaN(x) {
		max = x
		min = x
		if !math.IsNaN(y) {
			max = math.Max(max, y)
			min = math.Min(min, y)
		}
	} else {
		if !math.IsNaN(y) {
			max = y
			min = y
		}
	}
	return max, min
}

type Rounded struct {
	Rawvalue float64 // raw value with given digits
	Prefix   string  // prefix string
	Value    float64 // rounded value with given digits and prefix
	Response string  // response as <Value> <Prefix><unit>
}

// RoundPrefix returns the rounded value to given digits and correct prefix (Megas, Kilos etc.)
// Special case is abs value between 1000....10000 which is not converted to kilos (because I like it like that)
// set prefix to force certain prefix, otherwise the function figures it out on its' own.
func RoundPrefix(v float64, digits int, unit string, prefix string) Rounded {

	lowercaseunit := strings.ToLower(unit) // force lowercase

	prefixes := []string{"G", "M", "k", "m", "u", "n"}
	powers := []float64{1e9, 1e6, 1e3, 1e-3, 1e-6, 1e-9}
	noprefixUnits := []string{"%", "cycles", "years", "°c", "°lon", "°lat", "events", "", " "}

	// initialize response
	resp := Rounded{
		Prefix:   prefix,
		Rawvalue: RoundTo(v, digits),
	}

	// no prefix if noprefixUnit given
	if SliceContains(noprefixUnits, lowercaseunit) {
		prefix = "-"
	}

	if prefix != "-" {
		prefixpos := FindStringposition(prefixes, prefix)
		if prefixpos >= 0 {
			resp.Value = RoundTo(v/powers[prefixpos], digits)
			resp.Response = fmt.Sprintf("%v %v%v", resp.Prefix, resp.Value, unit)
			return resp
		}

		resp.Prefix = ""
		if math.Abs(v) >= 0.01 && math.Abs(v) <= 9999 {
			resp.Value = RoundTo(v, digits)
			resp.Response = fmt.Sprintf("%v %v%v", resp.Value, resp.Prefix, unit)
			return resp
		}

		for i, p := range prefixes {
			if math.Abs(v) >= powers[i] {
				resp.Prefix = p
				resp.Value = RoundTo(v/powers[i], digits)
				resp.Response = fmt.Sprintf("%v %v%v", resp.Value, resp.Prefix, unit)
				return resp
			}
		}
	}

	resp.Value = RoundTo(v, digits)
	resp.Response = fmt.Sprintf("%v %v%v", resp.Value, resp.Prefix, unit)
	return resp

}

//RoundTo rounds the number to given digits
func RoundTo(val float64, digits int) float64 {

	s := fmt.Sprintf("%."+strconv.Itoa(digits)+"g", val)
	r, _ := strconv.ParseFloat(s, 64)
	return r
}

// FindStringposition returns the first occurrence index of given string. Returns -1 if string was not found.
func FindStringposition(slice []string, s string) int {
	for i, a := range slice {
		if a == s {
			return i
		}
	}
	return -1
}

// FindFloatIndex finds the index of first occurrence of the given value
func FindFloatIndex(slice []float64, val float64) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}

func Xnor(a, b bool) bool {
	return !((a || b) && (!a || !b))
}

func GetSysmemory() float64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return RoundTo(float64(m.Sys)/1024/1024, 5)
}

// ExtractSources returns the levels a,b,c as integer list
func ExtractSources(input string) [3]int {
	sources := [3]int{-1, -1, -1}

	r := regexp.MustCompile(`[^\d]+(\d{1,})`)
	matches := r.FindAllStringSubmatch(input, -1)

	for index, v := range matches {
		sources[index] = StringToInt(v[1])
	}
	return sources
}

// MaxI returns bigger int from two ints
func MaxI(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

// MinI returns smaller int from two ints
func MinI(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

// SetDefaultI set default value for a parameter, if the configuration value is zero (=initialized empty variable)
func SetDefaultI(defaultval int, confval int) int {
	if confval == 0 {
		return defaultval
	}
	return confval
}

// SetDefaultF set default value for a parameter, if the configuration value is zero (=initialized empty variable)
func SetDefaultF(defaultval float64, confval float64) float64 {
	if confval == 0 {
		return defaultval
	}
	return confval
}
func StringInSlice(slice []string, s string, ignorecase bool) bool {
	if ignorecase {
		s = strings.ToLower(s)
	}
	for _, a := range slice {
		if ignorecase {
			a = strings.ToLower(a)
		}
		if a == s {
			return true
		}
	}
	return false
}

// FirstOfMonth returns a TZ timestamp of first day of given month
func FirstOfMonth(t time.Time) time.Time {
	return ParseTime(t.Format("2006-01") + "-01T00:00:00Z")
}
