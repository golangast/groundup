package addurltitle

import (
	"fmt"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
)

func UpdateUser(u Urls) {
	//opening database
	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//prepare statement so that no sql injection
	stmt, err := data.Prepare("update urls set lib=?, libtag=? where titles=?")
	ErrorCheck(err)

	//execute qeury
	res, err := stmt.Exec(u.Lib, u.Libtag, u.Titles)
	ErrorCheck(err)

	//used to print rows
	a, err := res.RowsAffected()
	ErrorCheck(err)
	fmt.Println(a, u.Lib)

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
