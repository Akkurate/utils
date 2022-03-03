package timescale

import (
	"fmt"
	"strings"

	"github.com/Akkurate/utils/logging"
)

// Result struct
type Result struct {
	Rows    []*Row
	IsEmpty bool
	First   *Row
}

// Query to Result struct
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

// QueryByRow goes through the data row by row with MapScan and executes given function for each row.
func (ts *Timescale) QueryByRow(query string, data map[string]interface{}, fn func(row *Row)) {

	for key, value := range data {
		query = strings.Replace(query, fmt.Sprintf("$%v", key), CastType(value), -1)
	}

	rows, err := ts.DB.Queryx(query)

	if err != nil {
		logging.Error("%v", err)
		panic(err)
	}

	for rows.Next() {
		m := make(map[string]interface{})

		rows.MapScan(m)
		res := &Row{
			Data: m,
		}
		fn(res)
	}
}

// Select queries the data to given interface struct
func (ts *Timescale) Select(dest interface{}, query string) {
	err := ts.DB.Select(dest, query)
	if err != nil {
		logging.Warn("%v", err)
	}
}

// Insert single record
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

// Update single record
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

// Insert single struct
func (ts *Timescale) InsertStruct(table string, data interface{}) {
	thatMap := structToMap(data)
	if debug {
		logging.Info("%v", thatMap)
	}
	ts.Insert(table, thatMap)

}

// Insert multiple records
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

type Batch struct {
	Rows []string
}

// Batch processing for fast raw data writing to Timescale. Returns the number of records written and possible error.
func (ts *Timescale) ProcessBatch(b *Batch, copyCmd, splitChar string) (int, error) {

	tx, err := ts.DB.Begin()
	if err != nil {
		logging.Warn("BEGIN: %v", err)
		return 0, err
	}

	stmt, err := tx.Prepare(copyCmd)
	if err != nil {
		logging.Warn("PREPARE: %v", err)
		return 0, err
	}

	for _, line := range b.Rows {
		// For some reason this is only needed for tab splitting
		if splitChar == "\t" {
			sp := strings.Split(line, splitChar)
			args := make([]interface{}, len(sp))
			for i, v := range sp {
				args[i] = v
			}
			_, err = stmt.Exec(args...)
		} else {
			_, err = stmt.Exec(line)
		}

		if err != nil {
			logging.Warn("ROWS: %v", err)
			return 0, err
		}
	}

	err = stmt.Close()
	if err != nil {
		logging.Warn("STMT: %v", err)
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		logging.Warn("COMMIT: %v", err)
		return 0, err
	}
	return len(b.Rows), nil
}
