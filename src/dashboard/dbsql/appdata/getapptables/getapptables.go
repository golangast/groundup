package getapptables

import (
	"fmt"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
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

		types    string
		name     string
		tbl_name string
		rootpage string
		sql      string
		TD       TableData
		TDS      []TableData
	)
	TD = TableData{}
	fmt.Print(TD)
	//SELECT type FROM sqlite_master where type='table'  AND name='urls
	//get from database SELECT * FROM sqlite_schema WHERE type ='table' AND name NOT LIKE 'sqlite_%';
	rows, err := data.Query("SELECT * FROM sqlite_master where type='table'  AND name='urls")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {

		err := rows.Scan(&types, &name, tbl_name, rootpage)
		ErrorCheck(err)

		i++
		fmt.Println("scan ", i, types, name, tbl_name, rootpage, sql)

		//store into memory
		tables = append(tables, name)

	}

	for _, table := range tables {
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
			TD.Name = append(TD.Name, table)
		}

	}
	TDS = append(TDS, TD)
	//close everything
	defer rows.Close()
	defer data.Close()
	return TDS

}
