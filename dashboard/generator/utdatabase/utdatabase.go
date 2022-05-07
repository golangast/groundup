package utdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func Createdbfile() string {
	if err := os.MkdirAll("db", os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println("Whoops, could not create directory(ies) because", err)
	}
	f, err := os.Create("db/database.db")
	if err != nil {
		panic(err)
	}

	if err := f.Close(); err != nil {
		panic(err)
	}
	if err := os.Remove(f.Name()); err != nil {
		panic(err)
	}
	return f.Name()
}

const create string = `CREATE TABLE IF NOT EXISTS "urls" (
  id INTEGER NOT NULL PRIMARY KEY,
  urls TEXT,
  titles TEXT,
  lib TEXT,
  libtag TEXT,
  css TEXT,
  csstag TEXT,
  filename TEXT,
  );`

const file string = "./db/urls.db"

func GenerateTable(title string, urls string) {
	db, err := Conn()
	ErrorChecks(err)
	if _, err := db.Exec(create); err != nil {
		fmt.Println(err)

	}
	defer db.Close()
}

func AddUrlTitle(Urls string, Titles string) {

	db, err := Conn() //create db instance
	ErrorChecks(err)
	var exists bool //used for checking

	//create a context query so that you can know if it exists already.
	//if it does then you can stop the context of the request.
	stmts := db.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM urls WHERE urls=?)", Urls)
	err = stmts.Scan(&exists)
	ErrorChecks(err)

	//prepare the statement to ensure no sql injection
	stmt, err := db.Prepare("INSERT INTO urls(urls, titles) VALUES(?, ?)")
	ErrorChecks(err)

	//actually make the execution of the query
	res, err := stmt.Exec(Urls, Titles)
	ErrorChecks(err)

	//get last id to double check
	lastId, err := res.LastInsertId()
	ErrorChecks(err)

	//get rows affected to double check
	rowCnt, err := res.RowsAffected()
	ErrorChecks(err)

	//print out what you actually did
	log.Printf("lastid = %d, affected = %d, titles = %s\n", lastId, rowCnt, Urls)
	defer db.Close()

}

func ErrorChecks(err error) {
	if err != nil {
		panic(err.Error())
	}
}
func Conn() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)
	//check if it pings
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	log.Printf("Connected to DB %s successfully\n", file)
	return db, nil
}
