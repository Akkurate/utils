package timescale

import "fmt"

func (ts *Timescale) FindByID(table string, id int) *Result {
	return ts.Query(fmt.Sprintf(`SELECT * FROM "%v" WHERE id = $id`, table), map[string]interface{}{
		"id": id,
	})
}

func (ts *Timescale) UpdateByID(table string, data map[string]interface{}, id int) {
	ts.Update(table, data, fmt.Sprintf("id=%v", id))
}
