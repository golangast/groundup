package getlib

import (
	"fmt"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
)

func GetLib(Libtag string) string {

	data, err := DbConnection() //create db instance
	ErrorCheckin(err)

	//variables used to store data from the query
	var (
		lib     string
		Library string
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select lib from urls WHERE libtag=" + Libtag + ";")
	ErrorCheckin(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&Libtag)
		ErrorCheckin(err)

		i++
		fmt.Println("scan ", i)

		//store into memory
		Library = lib

	}
	//close everything
	defer rows.Close()
	defer data.Close()
	return Library

}

func ErrorCheckin(err error) {
	if err != nil {
		panic(err.Error())
	}
}
