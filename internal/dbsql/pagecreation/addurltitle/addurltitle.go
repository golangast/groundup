package addurltitle

import (
	"context"
	"log"

	. "github.com/golangast/groundup/internal/dbsql/conn"
)

func AddUrlTitle(u Urls) {

	data, err := DbConnection() //create db instance
	ErrorCheck(err)
	var exists bool //used for checking

	//create a context query so that you can know if it exists already.
	//if it does then you can stop the context of the request.u.Urls, u.Titles
	stmts := data.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM urls WHERE urls=?)", u.Urls)
	err = stmts.Scan(&exists)
	ErrorCheck(err)
	if !exists {
		query := "INSERT INTO `urls` (`urls`, `titles`) VALUES (?, ?)"
		insertResult, err := data.ExecContext(context.Background(), query, u.Urls, u.Titles)
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
