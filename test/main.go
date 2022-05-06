package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

//you do need gcc installed
func main() {
	urls := GetUrls()
	fmt.Println(urls)

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

func GetUrls() []Urls {

	db, err := dbConnection() //create db instance
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)

	}
	//variables used to store data from the query
	var (
		id    string
		urls  string
		urlss []Urls //used to store all users
	)

	rows, err := db.Query("select id, urls from urls;")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var i int
		if err = rows.Scan(&id, &urls); err != nil {
			fmt.Println(err)

		} else {
			i++
			fmt.Println("scan ", i)
		}

		//store into memory
		u := Urls{ID: id, Urls: urls}
		urlss = append(urlss, u)
		fmt.Println(urlss)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)

	}

	if err = db.Close(); err != nil {
		fmt.Println(err)

	}

	return urlss

}

//https://golangbot.com/mysql-create-table-insert-row/
func dbConnection() (*sql.DB, error) {
	//db urls   conn to db      database used
	db, err := sql.Open("sqlite", file)
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
} //end of connect

const file string = "./db/urls.db"
