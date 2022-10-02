package addappdatabase

import (
	. "github.com/golangast/groundup/internal/dbsql/conn"
)

func Addappdatabase() {
	//opening database
	// data, err := AppDbConnection() //create db instance
	// ErrorCheck(err)
	// var exists bool

	// stmts := data.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM csstable WHERE css=?)", css)
	// err = stmts.Scan(&exists)
	// ErrorCheck(err)

	// //prepare the statement to ensure no sql injection
	// stmt, err := data.Prepare("INSERT INTO csstable(css, csstag) VALUES(?, ?)")
	// ErrorCheck(err)

	// //actually make the execution of the query
	// res, err := stmt.Exec(css, csstag)
	// ErrorCheck(err)

	// //get last id to double check
	// lastId, err := res.LastInsertId()
	// ErrorCheck(err)

	// //get rows affected to double check
	// rowCnt, err := res.RowsAffected()
	// ErrorCheck(err)

	// //print out what you actually did
	// log.Printf("lastid = %d, affected = %d, titles = %s\n", lastId, rowCnt, css)
	// defer data.Close()

}
