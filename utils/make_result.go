package utils

import (
	"database/sql"
	"goApi/global"
)

func MakeRowsResult(rows *sql.Rows) []map[string]string {
	result := make([]map[string]string, 0)
	columns, err := rows.Columns()
	if err != nil {
		global.OPLOGGER.Error(err)
		return nil
	}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		for i := range columns {
			values[i] = new(*string)
		}
		err := rows.Scan(values...)
		if err != nil {
			global.OPLOGGER.Error(err)
			return nil
		}
		row := map[string]string{}
		for k, v := range columns {
			val := *(values[k].(**string))
			if val == nil {
				row[v] = ""
			} else {
				row[v] = *val
			}
		}
		result = append(result, row)
	}

	err = rows.Err()
	if err != nil {
		global.OPLOGGER.Error(err)
		return nil
	}
	return result
}
