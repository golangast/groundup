package getapptables

import (
	"fmt"

	. "github.com/golangast/groundup/internal/dbsql/conn"
)

type TableData struct {
	Name    []string
	Columns []string
	Values  []string
}

func Getapptables() []TableData {

	data, err := AppDbConnection()
	ErrorCheck(err)

	var (
		tables []string
		//i      int
		types string
		name  string
		TDS   []TableData
	)

	//SELECT type FROM sqlite_master where type='table'  AND name='urls
	//get from database SELECT * FROM sqlite_schema WHERE type ='table' AND name NOT LIKE 'sqlite_%';
	rows, err := data.Query("SELECT type, name FROM sqlite_master where type='table'")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {

		err := rows.Scan(&types, &name)
		ErrorCheck(err)

		// i++
		// fmt.Println("scan ", i, types, name)

		//store into memory
		tables = append(tables, name)

	}
	//after table names have been appended
	//grab their data.
	for _, table := range tables {
		TD := TableData{}
		rows, _ := data.Query("SELECT * FROM " + table + ";")
		columns, _ := rows.Columns()
		count := len(columns)
		values := make([]any, count)
		valuePtr := make([]any, count)
		var v any
		var prevtable string
		for rows.Next() {

			if prevtable != table {
				//fmt.Print(" table:", table, " columns:", columns)
				//scan needs any type so turn columns into []any
				for i, _ := range columns {
					valuePtr[i] = &values[i]
				}
				prevtable = table
				rows.Scan(valuePtr...)
				//go through the columns
				for a, _ := range columns {

					val := values[a]

					b, ok := val.([]byte)

					if ok {
						v = string(b)
					} else {
						v = val
					}

					//fmt.Println(col+" --", v)
					//put them into TD data
					TD.Values = append(TD.Values, fmt.Sprint(v))

				}
			}

		}
		TD.Columns = columns
		TD.Name = append(TD.Name, table)
		if values[0] == nil {

			//TD.Columns = append(TD.Columns, columns...)
		}
		TDS = append(TDS, TD)

	}

	//close everything
	defer rows.Close()
	defer data.Close()
	return TDS

}
