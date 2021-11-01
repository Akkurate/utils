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

func getFloat(unk interface{}) (float64, error) {
	v := reflect.ValueOf(unk)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return math.NaN(), fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	fv := v.Convert(floatType)
	return fv.Float(), nil
}

// Value Value
type Value struct {
	Raw interface{}
}

// AsString AsString
func (v *Value) AsString() string {
	return fmt.Sprint(v.Raw)
}

// AsFloat AsFloat
func (v *Value) AsFloat() float64 {
	if v.Raw == nil {
		return math.NaN()
	}
	f, _ := getFloat(v.Raw)
	return f
}

// AsInt AsInt
func (v *Value) AsInt() int {
	if v.Raw == nil {
		return 0
	}
	return v.Raw.(int)
}

func (v *Value) BinaryToInt() int {
	if v.Raw == nil {
		return 0
	}
	a := v.Raw.([]uint8)

	u32LE := binary.LittleEndian.Uint32(a)

	return int(u32LE)
}

//

// AsInt64 AsInt64
func (v *Value) AsInt64() int64 {
	if v.Raw == nil {
		return 0
	}
	return v.Raw.(int64)
}

// AsTime AsTime
func (v *Value) AsTime() time.Time {
	return v.Raw.(time.Time)
}

// AsBool AsBool
func (v *Value) AsBool() bool {
	return v.Raw.(bool)
}

// AsBool AsBool
func (v *Value) AsInterface() interface{} {
	return v.Raw
}

// AsBool AsBool
func (v *Value) AsByte() []byte {
	return v.Raw.([]byte)
}

// AsBool AsBool
func (v *Value) AsJSON(target interface{}) {
	b := v.Raw.([]byte)
	json.Unmarshal(b, target)
}

// Row Row
type Row struct {
	Data map[string]interface{}
}

// Get Get
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

// Get Get
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
