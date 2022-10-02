package getapptabledata

import (
	"fmt"

	. "github.com/golangast/groundup/internal/dbsql/conn"
)

//just get the table names
func Getapptabledata() []string {
	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		tablenames []string
		stable     string
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select * from savedtable")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&stable)
		ErrorCheck(err)

		i++
		fmt.Println("scan ", i)

		//store into memory
		u := DBFields{Stable: stable}
		tablenames = append(tablenames, u.Stable)

	}
	//close everything
	defer rows.Close()
	defer data.Close()
	return tablenames

}
