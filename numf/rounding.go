package numf

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Akkurate/utils/str"
)

type Rounded struct {
	Rawvalue       float64 // raw value with given digits but no prefix
	Prefix         string  // prefix string
	Value          float64 // rounded value with given digits and prefix
	Response       string  // response string as <Value> <Prefix><unit>
	Prettyvalue    float64 // rounded value with one decimal
	Prettyresponse string  // response string as <Value> <Prefix><unit> ; if value is <0.1 , shows "less than 0.1"
}

// Returns the rounded value to given digits and correct prefix (M for Megas, k for Kilos etc.)
// Special case is abs(value) between 1000....10000, which is not converted to kilos (because it looks nicer)
// set prefix to force certain prefix, otherwise the function figures it out on its' own.
// These units are excluded from having a prefix
//  noprefixUnits := []string{"%", "cycles", "years", "°c", "°lon", "°lat", "events", "", " "}
func RoundWithPrefix(v float64, digits int, unit string, prefix string) Rounded {
	// initialize response
	resp := Rounded{
		Rawvalue: RoundTo(v, digits),
	}
	doResponses := func(value, prettyvalue float64, prf string) {
		resp.Response = fmt.Sprintf("%v %v%v", value, prf, unit)
		var sign string
		if value < 0 {
			sign = "-"
		}
		if prettyvalue == 0 {
			resp.Prettyresponse = fmt.Sprintf("is under %v0.1 %v%v", sign, prf, unit)
		} else {
			resp.Prettyresponse = fmt.Sprintf("%v %v%v", prettyvalue, prf, unit)
		}
	}
	lowercaseunit := strings.ToLower(unit) // force lowercase

	prefixes := []string{"G", "M", "k", "m", "u", "n"}
	powers := []float64{1e9, 1e6, 1e3, 1e-3, 1e-6, 1e-9}
	noprefixUnits := []string{"%", "cycles", "years", "°c", "°lon", "°lat", "events", "", " "}

	setprefix := true
	// no prefix if noprefixUnit given
	if str.Contains(noprefixUnits, lowercaseunit) {
		setprefix = false
		resp.Prefix = ""
	}

	if setprefix {
		prefixpos, _ := str.FindIndex(prefixes, prefix)
		if prefixpos >= 0 {
			resp.Value = RoundTo(v/powers[prefixpos], digits)
			resp.Prettyvalue = math.Round(10*resp.Value) / 10
			doResponses(resp.Value, resp.Prettyvalue, prefix)
			return resp
		}

		resp.Prefix = ""
		if math.Abs(v) >= 0.01 && math.Abs(v) <= 9999 {
			resp.Value = RoundTo(v, digits)
			resp.Prettyvalue = math.Round(10*resp.Value) / 10
			doResponses(resp.Value, resp.Prettyvalue, resp.Prefix)
			return resp
		}

		for i, p := range prefixes {
			if math.Abs(v) >= powers[i] {
				resp.Prefix = p
				resp.Value = RoundTo(v/powers[i], digits)
				resp.Prettyvalue = math.Round(10*resp.Value) / 10
				doResponses(resp.Value, resp.Prettyvalue, resp.Prefix)
				return resp
			}
		}
	}

	resp.Value = RoundTo(v, digits)
	resp.Prettyvalue = math.Round(10*resp.Value) / 10
	doResponses(resp.Value, resp.Prettyvalue, prefix)

	return resp

}

// Rounds the number to given significant digits.
func RoundTo(val float64, digits int) float64 {

	s := fmt.Sprintf("%."+strconv.Itoa(digits)+"g", val)
	r, _ := strconv.ParseFloat(s, 64)
	return r
}

// Rounds <value> to <nearest> value
func RoundToNearest(value float64, nearest float64) float64 {

	modulo := math.Mod(value, nearest)
	if modulo >= nearest/2 {
		return value - modulo + nearest
	}
	return value - modulo
}
