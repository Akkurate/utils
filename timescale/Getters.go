package timescale

import (
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/Akkurate/utils/logging"
	"github.com/Akkurate/utils/numf"
	"github.com/Akkurate/utils/str"
)

// Row contains Data of one Timescale row as a map where map keys are column names.
type Row struct {
	Data map[string]interface{}
}

// Get data from given column as Value. Returns nil if no data found.
func (r *Row) Get(field string) *Value {
	v, e := r.Data[field]
	if !e {
		return nil
	}
	if v == nil {
		return nil
	}
	return &Value{Raw: v}
}

// Get data from given column as Value. Returns defaultvalue if no data found.
func (r *Row) GetDefault(field string, defaultValue interface{}) *Value {
	v, e := r.Data[field]
	if !e {
		return &Value{Raw: defaultValue}
	}
	if v == nil {
		return &Value{Raw: defaultValue}
	}
	return &Value{Raw: v}
}

// Get data from given column as float64. Returns NaN if no data found.
func (r *Row) GetFloat(field string) float64 {
	v := r.Get(field)
	if v == nil {
		return math.NaN()
	}
	return v.AsFloat()
}

// Get data from given column as int. Returns -1 if no data found.
func (r *Row) GetInt(field string) int {
	v := r.Get(field)
	if v == nil {
		return -1
	}
	return int(v.AsInt64())
}

// Get data from given column as string.
func (r *Row) GetString(field string) string {
	v := r.Get(field)
	if v == nil {
		return ""
	}
	return v.AsString()
}

// Get data from given column as time.Time.
func (r *Row) GetTime(field string) time.Time {
	v := r.Get(field)
	if v == nil {
		return time.Time{}
	}
	return v.AsTime()
}

// Get all convertible data as float64 and return as a map. Blacklist of columns can be given to exclude certain columns.
func (r *Row) GetFloats(blacklist []string) *map[string]float64 {
	m := make(map[string]float64)
	for k, v := range r.Data {
		if !str.Contains(blacklist, k) {
			val := &Value{Raw: v}
			f := val.AsFloat()
			if numf.IsValid(f) {
				m[k] = f
			}
		}
	}
	return &m
}

// Fills the given struct with data from Timescale row. The struct to be filled is given as a pointer.
// Function will fill all the struct's fields it finds and leaves the rest of the fields untouched. Thus default values can be set to struct prior to filling it.
// NOTE: Struct field name must match the timescale column name, except that struct field always needs capital letter
func (r *Row) GetStruct(structPointer interface{}) {

	row := *r

	valueof := reflect.ValueOf(structPointer).Elem() // struct "value"
	typeof := reflect.TypeOf(structPointer).Elem()   // struct type

	// loop through all the struct fields...
	for i := 0; i < valueof.NumField(); i++ {

		dataValue := valueof.Field(i)                  // value of struct field
		field := strings.ToLower(typeof.Field(i).Name) // name of struct field in lowercase to match timescale column name

		fieldValue, found := row.Data[field] // try to get the fieldValue from timescale column ...

		if !found || fieldValue == nil {
			continue // no field or fieldValue was nil -- continue. Struct value will remain untouched.
		}

		fv := reflect.ValueOf(fieldValue)

		// set the struct fields dataValue with fieldValue from map. Some field type kinds may need special attention.
		switch reflect.TypeOf(fieldValue).Kind() {

		default:
			dataValue.Set(fv)

			// converts int64 to int -- Golang converts ints from interface always to Int64 type
		case reflect.Int64:
			fv = reflect.Indirect(fv)
			fv = fv.Convert(reflect.TypeOf(int(1)))
			dataValue.Set(reflect.ValueOf(fv.Interface()))

		case reflect.Slice:
			// all slices from Row are type []uint8 ... converted always to []float64 type at the moment since we are dealing mostly with numbers
			fv = reflect.Indirect(fv)
			asString := string(fv.Bytes()[1 : fv.Len()-1]) // convert reflect value to bytes, drop the { } -bytes from ends and convert to a string
			items := strings.Split(asString, ",")          // split the string into items == floats

			var floats []float64
			for _, s := range items {
				f, err := strconv.ParseFloat(s, 64)
				if err != nil {
					logging.Warn("Couldn't convert string %v to float64", s)
					continue
				}
				floats = append(floats, f)
			}
			dataValue.Set(reflect.ValueOf(floats))

		}

	}
}
