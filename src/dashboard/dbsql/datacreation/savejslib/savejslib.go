package savejslib

import (
	"context"
	"log"

	. "github.com/golangast/groundup/src/dashboard/dbsql/conn"
)

func Addsavedata(lib, libtag string) {
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
