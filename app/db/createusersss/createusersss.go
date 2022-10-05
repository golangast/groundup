
package createusersss

import (

	. "app/db"
	"fmt"
)

func Createusersss(){
data, err := DbConnection() //create db instance
if err != nil {
	fmt.Println(err)
}

statementsavedata, err := data.Exec("CREATE TABLE IF NOT EXISTS usersss (id integer NOT NULL primary KEY AUTOINCREMENT,  sss string  NOT NULL,   ggg int  NOT NULL ); ")
if err != nil{
	fmt.Println(err)
	}
	data.Close()
	fmt.Println(statementsavedata)
	}