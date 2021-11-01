package timescale

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"strings"
	"sync"
	"time"

	"github.com/Akkurate/utils/logging"
	"github.com/Akkurate/utils/system"
	"github.com/jmoiron/sqlx"

	// should be imported
	_ "github.com/lib/pq"
)

// Timescale Timescale
type Timescale struct {
	uri string
	DB  *sqlx.DB
	sync.Mutex
}

var debug = system.GetEnvOrDefault("QUERY_DEBUG", "") == "1"

// Result Result
type Result struct {
	Rows    []*Row
	IsEmpty bool
	First   *Row
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

func getLevelRecord(num int) string {
	if num < 0 {
		return "NULL"
	}
	return fmt.Sprintf("%v", num)
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

func (ts *Timescale) FindByID(table string, id int) *Result {
	return ts.Query(fmt.Sprintf(`SELECT * FROM "%v" WHERE id = $id`, table), map[string]interface{}{
		"id": id,
	})
}

// Query Query
func (ts *Timescale) Query(query string, data map[string]interface{}) *Result {

	var list []*Row
	for key, value := range data {
		query = strings.Replace(query, fmt.Sprintf("$%v", key), CastType(value), -1)
	}
	rows, err := ts.DB.Queryx(query)
	if debug {
		logging.Info("<cyan>[PG]</>: %v", query)
	}
	if err != nil {
		logging.Error("%v", err)
		return nil
	}
	for rows.Next() {
		m := make(map[string]interface{})
		rows.MapScan(m)
		res := &Row{
			Data: m,
		}
		list = append(list, res)
	}
	res := &Result{
		Rows:    list,
		IsEmpty: len(list) == 0,
	}
	if !res.IsEmpty {
		res.First = list[0]
	}
	return res
}

// ExecAndCommit ExecAndCommit
func (ts *Timescale) ExecAndCommit(query string, args ...interface{}) error {
	var argList []interface{}
	for _, arg := range args {
		argList = append(argList, escapeString(fmt.Sprintf("%v", arg)))
	}
	q := fmt.Sprintf(query, argList...)
	tx := ts.DB.MustBegin()
	if debug {
		logging.Info("%v", query)
	}
	tx.MustExec(q)
	return tx.Commit()
}

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

// Insert Insert
func (ts *Timescale) Insert(table string, data map[string]interface{}) {
	var keys []string
	var values []string
	for key, value := range data {
		keys = append(keys, fmt.Sprintf(`"%v"`, key))
		values = append(values, CastType(value))
	}
	q := fmt.Sprintf(`INSERT INTO %v (%v) VALUES (%v)`, table, strings.Join(keys, ","), strings.Join(values, ","))
	if debug {
		logging.Info("%v", q)
	}

	tx := ts.DB.MustBegin()
	tx.MustExec(q)

	tx.Commit()
}

func (ts *Timescale) Update(table string, data map[string]interface{}, where string) {
	var pairs []string

	for key, value := range data {
		pairs = append(pairs, fmt.Sprintf("%v = %v", key, CastType(value)))
	}
	q := fmt.Sprintf(`UPDATE "%v" SET %v WHERE %v`, table, strings.Join(pairs, ", "), where)
	if debug {
		logging.Info("%v", q)
	}

	tx := ts.DB.MustBegin()
	tx.MustExec(q)

	tx.Commit()
}

func (ts *Timescale) UpdateByID(table string, data map[string]interface{}, id int) {
	ts.Update(table, data, fmt.Sprintf("id=%v", id))
}

func (ts *Timescale) InsertStruct(table string, data interface{}) {
	thatMap := structToMap(data)
	if debug {
		logging.Info("%v", thatMap)
	}
	ts.Insert(table, thatMap)

}
func (ts *Timescale) BulkInsert(table string, data []map[string]interface{}) {
	var fields []string

	if len(data) > 0 {
		first := data[0]
		var sqlFields []string
		for key := range first {
			fields = append(fields, key)
			sqlFields = append(sqlFields, fmt.Sprintf(`"%v"`, key))
		}
		var sqlValues []string
		for _, vals := range data {
			var a []string
			for _, key := range fields {
				a = append(a, CastType(vals[key]))
			}
			sqlValues = append(sqlValues, fmt.Sprintf(`(%v)`, strings.Join(a, ",")))
		}
		q := fmt.Sprintf(`INSERT INTO "%v" (%v) VALUES`, table, strings.Join(sqlFields, ", "))
		finalQuery := q + "\n" + strings.Join(sqlValues, ",")

		if debug {
			logging.Info("%v", finalQuery)
		}
		if len(sqlValues) > 0 {
			tx := ts.DB.MustBegin()
			tx.MustExec(finalQuery)
			tx.Commit()
		}
	}
}

// Close close connection
func (ts *Timescale) Close() {
	ts.DB.Close()
}

// New New
func New(config string) *Timescale {
	ts := &Timescale{
		uri: config,
	}
	ts.Connect()
	return ts
}

// Connect Connect
func (ts *Timescale) Connect() {

	a := strings.Index(ts.uri, "?sslmode=disable")
	if a == -1 {
		ts.uri = fmt.Sprintf("%v?sslmode=disable", ts.uri)
	}
	logging.Info("<cyan>Connecting to timescale</> <gray>%v</>", ts.uri)
	db, err := sqlx.Connect("postgres", ts.uri)
	if err != nil {
		log.Fatalln(err)
		return
	}
	logging.Info("<green>Connected to TimeScale!</>")
	ts.DB = db
}
