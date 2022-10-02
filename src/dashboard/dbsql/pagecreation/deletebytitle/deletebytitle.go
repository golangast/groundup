package deletebytitle

import (
	"errors"

	. "github.com/golangast/groundup/src/dashboard/dbsql/conn"
)

func Deletebytitle(titles string) error {
	data, err := DbConnection() //create db instance
	if err != nil {
		return err
	}
	var ErrDeleteFailed = errors.New("delete failed")
	res, err := data.Exec("DELETE FROM urls WHERE titles = ?", titles)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}
