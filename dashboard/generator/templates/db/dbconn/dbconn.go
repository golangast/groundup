package dbconn

var Dbconntemp = `
package dbconn

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

const file string = "./db/database.db"

func DbConnection() (*sql.DB, error) {
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
`
