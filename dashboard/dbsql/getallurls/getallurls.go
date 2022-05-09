package geturls

import (
	"fmt"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
)

func GetUrls() []Urls {

	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		id       string
		urls     string
		titles   string
		lib      string
		libtag   string
		css      string
		csstag   string
		filename string
		urlss    []Urls //used to store all users
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select * from urls")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &urls, &titles, &lib, &libtag, &css, &csstag, &filename)
		ErrorCheck(err)

		i++
		fmt.Println("scan ", i)

		//store into memory
		u := Urls{ID: id, Urls: urls, Titles: titles, Lib: lib, Libtag: libtag, Css: css, Csstag: csstag, Filename: filename}
		urlss = append(urlss, u)

	}
	//close everything
	defer rows.Close()
	defer data.Close()
	return urlss

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
		fmt.Println(err.Error())
	}
}
