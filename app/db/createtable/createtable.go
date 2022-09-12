
package createtable

import (

	. "app/db"
	"fmt"
)

func Createdb(){
data, err := DbConnection() //create db instance
if err != nil {
	fmt.Println(err)
}

statementsavedata, err := data.Exec("CREATE TABLE IF NOT EXISTS customer (id integer NOT NULL primary KEY AUTOINCREMENT,  name string  NOT NULL,   names string  NOT NULL ); ")
if err != nil{
	fmt.Println(err)
	}
	data.Close()
	fmt.Println(statementsavedata)
	}