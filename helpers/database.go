package helpers

import (
	"encoding/json"
	"fmt"
	"log"
)

// Returns an array of all loot locations and values to plot on the map in iOS
func DumpDatabase(sqlString string) (string, error) {
	rows, err := ag.DB.Query(sqlString)
	if err != nil {
		log.Panic(err)
	}

	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		log.Panic(err)
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	jsonData, err := json.Marshal(tableData)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(jsonData))
	return string(jsonData), nil
}
