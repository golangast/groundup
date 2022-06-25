
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

statementsavedata, err := data.Exec("CREATE TABLE IF NOT EXISTS usertest (id integer NOT NULL primary KEY AUTOINCREMENT,  toy1 string  NOT NULL,   toy2 int  NOT NULL ); ")
if err != nil{
	fmt.Println(err)
	}
	data.Close()
	fmt.Println(statementsavedata)
	}