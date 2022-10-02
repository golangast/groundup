package deletetable

import (
	"fmt"

	. "github.com/golangast/groundup/internal/dbsql/conn"
)

func Deletetable(table string) error {
	//dashboard deletion of the app
	data, err := DbConnection() //create db instance
	ErrorCheck(err)
	res, err := data.Exec("DELETE FROM savedtable WHERE stable =$1", table)
	ErrorCheck(err)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		fmt.Println("rows affected were 0!!")
	}
	//app deletion of the table
	dataapp, err := AppDbConnection() //create db instance
	ErrorCheck(err)
	resapp, err := dataapp.Exec("DROP TABLE " + table)
	ErrorCheck(err)

	rowsAffectedapp, err := resapp.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffectedapp == 0 {
		fmt.Println("rows affected were 0!!")
	}

	return nil
}
