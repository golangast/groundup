package addurltitle

import (
	"context"
	"log"
	"time"

	. "github.com/golangast/groundup/internal/dbsql/conn"
)

func AddUrlTitle(u Urls) error {

	data, err := DbConnection() //create db instance
	ErrorCheck(err)
	var exists bool //used for checking

	//create a context query so that you can know if it exists already.
	//if it does then you can stop the context of the request.u.Urls, u.Titles
	stmts := data.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM urls WHERE urls=?)", u.Urls)
	err = stmts.Scan(&exists)
	ErrorCheck(err)
	if !exists {
		query := "INSERT INTO `urls` (`urls`, `titles`, `lib`, `libtag`,`css`,`csstag`,`filename`,`datavars`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelfunc()
		stmt, err := data.PrepareContext(ctx, query)
		if err != nil {
			log.Printf("Error %s when preparing SQL statement", err)
			return err
		}
		defer stmt.Close()
		res, err := stmt.ExecContext(ctx, u.Urls, u.Titles, "", "", "", "", "", "")
		if err != nil {
			log.Printf("Error %s when inserting row into products table", err)
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return err
		}
		log.Printf("%d products created ", rows)
		return nil

	}

	data.Close()

	// insertResult, err := data.ExecContext(context.Background())
	// if err != nil {
	// 	fmt.Println("impossible insert :", err)

	// }
	// ids, err := insertResult.LastInsertId()
	// if err != nil {
	// 	fmt.Println("impossible to retrieve last inserted id:", err)
	// }
	// log.Printf("inserted id: %d", ids)
	return nil
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
