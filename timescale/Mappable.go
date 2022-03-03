package timescale

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"time"
)

var floatType = reflect.TypeOf(float64(0))
var floatsType = reflect.TypeOf([]float64{})

// Value struct contains the Raw data of unknown type
type Value struct {
	Raw interface{}
}

func getFloat(unk interface{}) (float64, error) {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return math.NaN(), fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	fv := v.Convert(floatType)
	return fv.Float(), nil
}

// Returns raw value as string
func (v *Value) AsString() string {
	return fmt.Sprint(v.Raw)
}

// Returns raw value as float64. Returns NaN if no value.
func (v *Value) AsFloat() float64 {
	if v.Raw == nil {
		return math.NaN()
	}
	f, _ := getFloat(v.Raw)
	return f
}

// Returns raw value as int. Returns 0 if no value.
func (v *Value) AsInt() int {
	if v.Raw == nil {
		return 0
	}
	return v.Raw.(int)
}

// Returns raw value as int64. Returns 0 if no value.
func (v *Value) AsInt64() int64 {
	if v.Raw == nil {
		return 0
	}
	return v.Raw.(int64)
}

// Returns raw binary value as int. Returns 0 if no value.
func (v *Value) BinaryToInt() int {
	if v.Raw == nil {
		return 0
	}
	a := v.Raw.([]uint8)

	u32LE := binary.LittleEndian.Uint32(a)

	return int(u32LE)
}

// Returns raw value as time.Time
func (v *Value) AsTime() time.Time {
	return v.Raw.(time.Time)
}

// Returns raw value as boolean.
func (v *Value) AsBool() bool {
	return v.Raw.(bool)
}

// Returns raw value as interface.
func (v *Value) AsInterface() interface{} {
	return v.Raw
}

// Returns raw value as byte.
func (v *Value) AsByte() []byte {
	return v.Raw.([]byte)
}

// Unmarshals JSON to target.
func (v *Value) AsJSON(target interface{}) {
	b := v.Raw.([]byte)
	json.Unmarshal(b, target)
}

func (v *Value) AsFloats() []float64 {
	unk := reflect.ValueOf(v.Raw)
	unk = reflect.Indirect(unk)
	if !unk.Type().ConvertibleTo(floatsType) {
		return nil
	}
	f := unk.Convert(floatsType)

	res := make([]float64, f.Len())
	for i := 0; i < f.Len(); i++ {
		res[i] = f.Index(i).Float()
	}
	return res
}
