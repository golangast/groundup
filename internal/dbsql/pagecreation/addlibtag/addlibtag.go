package addurltitle

import (
	"context"
	"log"

	. "github.com/golangast/groundup/internal/dbsql/conn"
)

func UpdateUrls(lib, tag, titles string) {
	//opening database
	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//prepare statement so that no sql injection
	stmt, err := data.Prepare("update urls set lib=?, libtag=? where titles=?")
	ErrorCheck(err)

	//execute qeury
	_, err = stmt.Exec(lib, tag, titles)
	ErrorCheck(err)

}

func Addlib(lib, libtag string) {
	//opening database
	data, err := DbConnection() //create db instance
	ErrorCheck(err)
	var exists bool
	stmts := data.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM library WHERE lib=?)", lib)
	err = stmts.Scan(&exists)
	ErrorCheck(err)

	if !exists {
		query := "INSERT INTO `library` (`lib`, `libtag`) VALUES (?, ?)"
		insertResult, err := data.ExecContext(context.Background(), query, lib, libtag)
		if err != nil {
			log.Fatalf("impossible insert : %s", err)
		}
		ids, err := insertResult.LastInsertId()
		if err != nil {
			log.Fatalf("impossible to retrieve last inserted id: %s", err)
		}
		log.Printf("inserted id: %d", ids)
	}

	data.Close()

}
func AddCSSlib(css, csstag string) {
	//opening database
	data, err := DbConnection() //create db instance
	ErrorCheck(err)
	var exists bool
	stmts := data.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM csstable WHERE css=?)", css)
	err = stmts.Scan(&exists)
	ErrorCheck(err)

	if !exists {
		query := "INSERT INTO `csstable` (`css`, `csstag`) VALUES (?, ?)"
		insertResult, err := data.ExecContext(context.Background(), query, css, csstag)
		if err != nil {
			log.Fatalf("impossible insert : %s", err)
		}
		ids, err := insertResult.LastInsertId()
		if err != nil {
			log.Fatalf("impossible to retrieve last inserted id: %s", err)
		}
		log.Printf("inserted id: %d", ids)
	}

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
