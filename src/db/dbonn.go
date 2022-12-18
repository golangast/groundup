package dbconn

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

//go:embed urls.db
var Dashboarddb embed.FS

const file string = "./src/db/database.db"

func DbConnection() (*sql.DB, error) {
	//db urls   conn to db      database used
	ddb, err := getAllFilenamesdb(&Dashboarddb)
	if err != nil {
		fmt.Print(err)
	}
	db, err := sql.Open("sqlite", ddb)
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
func getAllFilenamesdb(efs *embed.FS) (files string, err error) {
	if err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	}); err != nil {
		return "", err
	}

	return files, nil
}
