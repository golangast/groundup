package adddatavartopage

import (
	"context"
	"fmt"

	. "github.com/golangast/groundup/internal/dbsql/conn"
)

func Adddatavartopage(datavars, urls string) {

	//opening database
	data, err := DbConnection() //create db instance
	ErrorCheck(err)
	var exists bool
	stmts := data.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM urls WHERE datavars LIKE=*?*)", datavars)
	err = stmts.Scan(&exists)
	ErrorCheck(err)

	if !exists {

		stmt, e := data.Prepare("UPDATE urls SET datavars=? WHERE urls=?")
		ErrorCheck(e)

		// execute
		res, e := stmt.Exec(datavars, urls)
		ErrorCheck(e)

		a, e := res.RowsAffected()
		ErrorCheck(e)

		fmt.Println(a)

	}

	data.Close()

}
