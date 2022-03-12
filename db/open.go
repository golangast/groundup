package db

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func Conn() *leveldb.DB {
	db, err := leveldb.OpenFile("path/to/db", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
