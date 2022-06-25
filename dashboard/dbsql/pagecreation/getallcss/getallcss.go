package getallcss

import (
	"fmt"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
)

func Getallcss() []CSS {

	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		id     string
		css    string
		csstag string
		CSSS   []CSS //used to store all users
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select * from csstable")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &css, &csstag)
		ErrorCheck(err)

		i++
		fmt.Println("scan ", i)

		//store into memory
		u := CSS{ID: id, Css: css, Csstag: csstag}
		CSSS = append(CSSS, u)

	}
	//close everything
	defer rows.Close()
	defer data.Close()
	return CSSS

}

type CSS struct {
	ID     string `param:"id" query:"id" form:"id"`
	Css    string `param:"css" query:"css" form:"css"`
	Csstag string `param:"csstag" query:"csstag" form:"csstag"`
}
