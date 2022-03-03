package timescale

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func CastType(iface interface{}) string {

	item := reflect.ValueOf(iface)
	rt := reflect.TypeOf(iface)
	switch rt.Kind() {
	case reflect.String:
		return fmt.Sprintf(`'%v'`, escapeString(fmt.Sprintf("%v", iface)))
	case reflect.Map:
		s, e := json.Marshal(iface)
		if e != nil {
			return "'{}'"
		}
		return fmt.Sprintf(`'%v'`, escapeString(string(s)))
	case reflect.Func:
		response := item.Call([]reflect.Value{})
		return fmt.Sprintf("%v", response[0])
	case reflect.Int:
	case reflect.Int64:
	case reflect.Float64:
		return fmt.Sprintf("%v", iface)
	case reflect.Slice:
		var snippets []string

		for i := 0; i < item.Len(); i++ {
			target := item.Index(i)
			snippets = append(snippets, CastType(target))
		}
		return strings.Join(snippets, ",")
	}

	switch iface.(type) {

	case time.Time:
		return fmt.Sprintf(`'%v'`, iface.(time.Time).Format("2006-01-02T15:04:05"))
	}
	return fmt.Sprintf(`%v`, iface)
}

func escapeString(value string) string {
	replace := map[string]string{
		//"\\":   "\\\\",
		"'":    `\'`,
		"\\0":  "\\\\0",
		"\n":   "\\n",
		"\r":   "\\r",
		"\x1a": "\\Z",
	}
	for b, a := range replace {
		value = strings.Replace(value, b, a, -1)
	}
	return value
}
func structToMap(item interface{}) map[string]interface{} {

	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		if tag != "" && tag != "-" {
			field := reflectValue.Field(i).Interface()
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = structToMap(field)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}

// ToSlice ToSlice
func (r *Result) ToSlice() []map[string]interface{} {

	list := make([]map[string]interface{}, 0)
	for _, row := range r.Rows {
		list = append(list, row.Data)
	}

	return list
}
