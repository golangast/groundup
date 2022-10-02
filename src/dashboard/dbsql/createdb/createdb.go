package createdb

import (
	"fmt"
	"os"

	. "github.com/golangast/groundup/src/dashboard/dbsql/conn"
)

func CreateDB() {
	exists, err := Exists("db/urls.db")
	if err != nil {
		fmt.Println(err)
	}

	if !exists {
		if err := os.MkdirAll("db", os.ModeSticky|os.ModePerm); err != nil {
			fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
		} else {
			fmt.Println("Whoops, could not create directory(ies) because", err)
		}
		_, err := os.Create("db/urls.db")
		if err != nil {
			fmt.Println(err)
		}
	}

	data, err := DbConnection() //create db instance
	if err != nil {
		fmt.Println(err)
	}
	// statementsavedata, err := data.Exec("CREATE TABLE IF NOT EXISTS savedb (id integer NOT NULL primary KEY(id AUTOINCREMENT,table string NOT NULL,fields string NOT NULL,types string NOT NULL) ")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	statementsavetable, err := data.Exec("CREATE TABLE IF NOT EXISTS savedtable (id integer NOT NULL primary KEY AUTOINCREMENT,stable string NOT NULL,f0 string NOT NULL,f1 string NOT NULL,f2 string NOT NULL,f3 string NOT NULL,f4 string NOT NULL,f5 string NOT NULL,f6 string NOT NULL,f7 string NOT NULL,f8 string NOT NULL,f9 string NOT NULL,f10 string NOT NULL,f11 string NOT NULL, t0 string NOT NULL,t1 string NOT NULL,t2 string NOT NULL,t3 string NOT NULL,t4 string NOT NULL,t5 string NOT NULL,t6 string NOT NULL,t7 string NOT NULL,t8 string NOT NULL,t9 string NOT NULL,t10 string NOT NULL,t11 string NOT NULL); ")
	if err != nil {
		fmt.Println(err)
	}

	data.Close()
	fmt.Println(statementsavetable)
}
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
