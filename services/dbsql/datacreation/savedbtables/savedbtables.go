package savedbtables

import (
	"context"
	"log"

	. "github.com/golangast/groundup/internal/dbsql/conn"
)

//saves the table in the dashboard only
func Savedbtables(dbf *DBFields) {
	//opening database
	data, err := DbConnection() //create db instance
	ErrorCheck(err)
	var exists bool

	stmts := data.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM savedtable WHERE f0=?)", &dbf.F0)
	err = stmts.Scan(&exists)
	ErrorCheck(err)

	//prepare the statement to ensure no sql injection
	stmt, err := data.Prepare("INSERT INTO savedtable(stable, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, t0,t1,t2,t3,t4,t5,t6,t7,t8,t9,t10,t11) VALUES(?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?,?, ?, ?, ?, ?)")
	ErrorCheck(err)

	//actually make the execution of the query
	res, err := stmt.Exec(dbf.Stable, dbf.F0, dbf.F1, dbf.F2, dbf.F3, dbf.F4, dbf.F5, dbf.F6, dbf.F7, dbf.F8, dbf.F9, dbf.F10, dbf.F11, dbf.T0, dbf.T1, dbf.T2, dbf.T3, dbf.T4, dbf.T5, dbf.T6, dbf.T7, dbf.T8, dbf.T9, dbf.T10, dbf.T11)
	ErrorCheck(err)

	//get last id to double check
	lastId, err := res.LastInsertId()
	ErrorCheck(err)

	//get rows affected to double check
	rowCnt, err := res.RowsAffected()
	ErrorCheck(err)

	//print out what you actually did
	log.Printf("lastid = %d, affected = %d, tables = %s\n", lastId, rowCnt, dbf.Stable)
	defer data.Close()

}
