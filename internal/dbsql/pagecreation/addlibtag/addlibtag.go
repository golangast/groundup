package addurltitle

import (
	"context"
	"fmt"
	"log"

	. "github.com/golangast/groundup/internal/dbsql/conn"
)

func UpdateUrls(lib, tag, titles string) {
	//opening database
	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//prepare statement so that no sql injection
	stmt, err := data.Prepare("update urls set lib=?, libtag=? where titles=?")
	ErrorCheck(err)

	//execute qeury
	res, err := stmt.Exec(lib, tag, titles)
	ErrorCheck(err)

	//used to print rows
	a, err := res.RowsAffected()
	ErrorCheck(err)
	fmt.Println(a, lib)

}

func Addlib(lib, libtag string) {
	//opening database
	data, err := DbConnection() //create db instance
	ErrorCheck(err)
	var exists bool
	stmts := data.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM library WHERE lib=?)", lib)
	err = stmts.Scan(&exists)
	ErrorCheck(err)

	//prepare the statement to ensure no sql injection
	stmt, err := data.Prepare("INSERT INTO library(lib, libtag) VALUES(?, ?)")
	ErrorCheck(err)

	//actually make the execution of the query
	res, err := stmt.Exec(lib, libtag)
	ErrorCheck(err)

	//get last id to double check
	lastId, err := res.LastInsertId()
	ErrorCheck(err)

	//get rows affected to double check
	rowCnt, err := res.RowsAffected()
	ErrorCheck(err)

	//print out what you actually did
	log.Printf("lastid = %d, affected = %d, titles = %s\n", lastId, rowCnt, lib)
	defer data.Close()

}
func AddCSSlib(css, csstag string) {
	//opening database
	data, err := DbConnection() //create db instance
	ErrorCheck(err)
	var exists bool
	stmts := data.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM csstable WHERE css=?)", css)
	err = stmts.Scan(&exists)
	ErrorCheck(err)

	//prepare the statement to ensure no sql injection
	stmt, err := data.Prepare("INSERT INTO csstable(css, csstag) VALUES(?, ?)")
	ErrorCheck(err)

	//actually make the execution of the query
	res, err := stmt.Exec(css, csstag)
	ErrorCheck(err)

	//get last id to double check
	lastId, err := res.LastInsertId()
	ErrorCheck(err)

	//get rows affected to double check
	rowCnt, err := res.RowsAffected()
	ErrorCheck(err)

	//print out what you actually did
	log.Printf("lastid = %d, affected = %d, titles = %s\n", lastId, rowCnt, css)
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
