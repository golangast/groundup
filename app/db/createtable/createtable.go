
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

statementsavedata, err := data.Exec("CREATE TABLE IF NOT EXISTS fff (id integer NOT NULL primary KEY AUTOINCREMENT,  aas string  NOT NULL,   ffs int  NOT NULL ); ")
if err != nil{
	fmt.Println(err)
	}
	data.Close()
	fmt.Println(statementsavedata)
	}