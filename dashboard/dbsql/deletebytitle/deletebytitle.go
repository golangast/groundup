package deletebytitle

import (
	"fmt"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
)

func Deletebytitle(titles string) {
	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	stmt, err := data.Prepare("delete from urls where titles=?")
	ErrorCheck(err)

	res, err := stmt.Exec(titles)
	ErrorCheck(err)

	// affected rows
	a, err := res.RowsAffected()
	ErrorCheck(err)

	fmt.Println(a)
}
func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
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
