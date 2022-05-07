package getpage

import (
	"fmt"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
)

func GetPage() ([]string, []string) {

	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		url    string
		title  string
		Urls   []string
		Titles []string
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select * from urls")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&url, &title)
		ErrorCheck(err)

		i++
		fmt.Println("scan ", i)

		//store into memory
		Urls = append(Urls, url)
		Titles = append(Titles, url)

	}
	//close everything
	defer rows.Close()
	defer data.Close()
	return Urls, Titles

}

type Urls struct {
	ID       string `param:"id" query:"id" form:"id"`
	Urls     string `param:"urls" query:"urls" form:"urls"`
	Titles   string `param:"titles" query:"titles" form:"titles"`
	Lib      string `param:"lib" query:"lib" form:"lib"`
	Libtag   string `param:"libtag" query:"libtag" form:"libtag"`
	Css      string `param:"css" query:"css" form:"css"`
	Csstag   string `param:"csstag" query:"csstag" form:"csstag"`
	Filename string `param:"filename" query:"filename" form:"filename"`
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
