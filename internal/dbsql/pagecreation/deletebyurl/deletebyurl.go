package deletebyurl

import (
	"errors"

	. "github.com/golangast/groundup/internal/dbsql/conn"
)

func Deletebyurl(urls string) error {
	data, err := DbConnection() //create db instance
	if err != nil {
		return err
	}
	slashurls := "/" + urls + "/"
	var ErrDeleteFailed = errors.New("delete failed")
	res, err := data.Exec("DELETE FROM urls WHERE urls = ?", slashurls)
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
