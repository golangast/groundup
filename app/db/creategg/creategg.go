
package creategg

import (

	. "app/db"
	"fmt"
)

func Creategg(){
data, err := DbConnection() //create db instance
if err != nil {
	fmt.Println(err)
}

statementsavedata, err := data.Exec("CREATE TABLE IF NOT EXISTS gg (id integer NOT NULL primary KEY AUTOINCREMENT,  names string  NOT NULL,   age int  NOT NULL ); ")
if err != nil{
	fmt.Println(err)
	}
	data.Close()
	fmt.Println(statementsavedata)
	}