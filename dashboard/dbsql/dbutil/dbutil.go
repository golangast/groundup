package dbutil

import (
	"context"
	"fmt"
	"log"
	"os"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
)

func GenerateTable() {
	db, err := DbConnection()
	ErrorCheckdb(err)
	statementurls, err := db.Exec(`CREATE TABLE IF NOT EXISTS urls (id integer NOT NULL primary key, urls text NOT NULL, titles text NOT NULL, lib text NOT NULL libtag text NOT NULL, css text NOT NULL, csstag text NOT NULL, filename text NOT NULL);`)
	ErrorCheckdb(err)

	statementlibrary, err := db.Exec(`CREATE TABLE IF NOT EXISTS library (id integer NOT NULL primary key, lib text NOT NULL, libtag text NOT NULL )`)
	ErrorCheckdb(err)

	statementcss, err := db.Exec(`CREATE TABLE IF NOT EXISTS csstable (id integer NOT NULL primary key, css text NOT NULL, csstag text NOT NULL )`)
	ErrorCheckdb(err)
	db.Close()

	fmt.Println(statementurls, statementlibrary, statementcss)
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func AddUrlTitle(Urls string, Titles string) {

	db, err := DbConnection() //create db instance
	ErrorCheckdb(err)
	var exists bool //used for checking

	//create a context query so that you can know if it exists already.
	//if it does then you can stop the context of the request.
	stmts := db.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM urls WHERE urls=?)", Urls)
	err = stmts.Scan(&exists)
	ErrorCheckdb(err)

	//prepare the statement to ensure no sql injection
	stmt, err := db.Prepare("INSERT INTO urls(urls, titles) VALUES(?, ?)")
	ErrorCheckdb(err)

	//actually make the execution of the query
	res, err := stmt.Exec(Urls, Titles)
	ErrorCheckdb(err)

	//get last id to double check
	lastId, err := res.LastInsertId()
	ErrorCheckdb(err)

	//get rows affected to double check
	rowCnt, err := res.RowsAffected()
	ErrorCheckdb(err)

	//print out what you actually did
	log.Printf("lastid = %d, affected = %d, titles = %s\n", lastId, rowCnt, Urls)
	defer db.Close()

}

func ErrorCheckdb(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
