package addurltitle

import (
	"context"
	"log"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
)

func AddUrlTitle(u Urls) {

	data, err := DbConnection() //create db instance
	var exists bool             //used for checking

	//create a context query so that you can know if it exists already.
	//if it does then you can stop the context of the request.
	stmts := data.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM urls WHERE urls=?)", u.Urls)
	err = stmts.Scan(&exists)
	ErrorCheck(err)

	//prepare the statement to ensure no sql injection
	stmt, err := data.Prepare("INSERT INTO urls(urls, titles) VALUES(?, ?)")
	ErrorCheck(err)

	//actually make the execution of the query
	res, err := stmt.Exec(u.Urls, u.Titles)
	ErrorCheck(err)

	//get last id to double check
	lastId, err := res.LastInsertId()
	ErrorCheck(err)

	//get rows affected to double check
	rowCnt, err := res.RowsAffected()
	ErrorCheck(err)

	//print out what you actually did
	log.Printf("lastid = %d, affected = %d, titles = %d\n", lastId, rowCnt, u.Urls)
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


