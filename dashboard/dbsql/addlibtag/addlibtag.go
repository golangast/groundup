package addurltitle

import (
	"context"
	"fmt"
	"log"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
)

func UpdateUrls(u Urls) {
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
func Addlib(u Urls) {
	//opening database
	data, err := DbConnection() //create db instance
	ErrorCheck(err)
	var exists bool
	stmts := data.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM urls WHERE lib=?)", u.Lib)
	err = stmts.Scan(&exists)
	ErrorCheck(err)

	//prepare the statement to ensure no sql injection
	stmt, err := data.Prepare("INSERT INTO urls(lib, libtag) VALUES(?, ?)")
	ErrorCheck(err)

	//actually make the execution of the query
	res, err := stmt.Exec(u.Lib, u.Libtag)
	ErrorCheck(err)

	//get last id to double check
	lastId, err := res.LastInsertId()
	ErrorCheck(err)

	//get rows affected to double check
	rowCnt, err := res.RowsAffected()
	ErrorCheck(err)

	//print out what you actually did
	log.Printf("lastid = %d, affected = %d, titles = %d\n", lastId, rowCnt, u.Lib)
	defer data.Close()

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
