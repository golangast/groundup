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
	//TODO get the apps database data

	data, err := AppDbConnection()
	ErrorCheck(err)

	var (
		tables []string
		i      int

		types string
		name  string
		// tbl_name string
		// rootpage string
		// sql      string

		TDS []TableData
	)

	//SELECT type FROM sqlite_master where type='table'  AND name='urls
	//get from database SELECT * FROM sqlite_schema WHERE type ='table' AND name NOT LIKE 'sqlite_%';
	rows, err := data.Query("SELECT type, name FROM sqlite_master where type='table'")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {

		err := rows.Scan(&types, &name)
		ErrorCheck(err)

		i++
		fmt.Println("scan ", i, types, name)

		//store into memory
		tables = append(tables, name)

	}

	for _, table := range tables {
		TD := TableData{}
		rows, _ := data.Query("SELECT * FROM " + table + ";")
		columns, _ := rows.Columns()
		count := len(columns)
		values := make([]any, count)
		valuePtr := make([]any, count)

		for rows.Next() {

			for i, _ := range columns {
				valuePtr[i] = &values[i]
			}

			rows.Scan(valuePtr...)

			for i, col := range columns {

				var v any

				val := values[i]

				b, ok := val.([]byte)

				if ok {
					v = string(b)
				} else {
					v = val
				}

				fmt.Println(col+" --", v)

				TD.Columns = append(TD.Columns, col)
				TD.Values = append(TD.Values, fmt.Sprint(v))

			}

		}
		TD.Name = append(TD.Name, table)
		if values[0] == nil {

			TD.Columns = append(TD.Columns, columns...)
		}
		TDS = append(TDS, TD)

	}

	//close everything
	defer rows.Close()
	defer data.Close()
	return TDS

}
