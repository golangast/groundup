package deletebytitle

import (
	"fmt"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
)

func Deletebytitle(titles string) {
	data, err := DbConnection() //create db instance
	ErrorChecked(err)

	stmt, err := data.Prepare("delete from urls where titles=?")
	ErrorChecked(err)

	res, err := stmt.Exec(titles)
	ErrorChecked(err)

	// affected rows
	a, err := res.RowsAffected()
	ErrorChecked(err)

	fmt.Println(a)
}
func ErrorChecked(err error) {
	if err != nil {
		panic(err.Error())
	}
}
