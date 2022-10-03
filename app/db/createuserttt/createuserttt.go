
package createuserttt

import (

	. "app/db"
	"fmt"
)

func Createuserttt(){
data, err := DbConnection() //create db instance
if err != nil {
	fmt.Println(err)
}

statementsavedata, err := data.Exec("CREATE TABLE IF NOT EXISTS userttt (id integer NOT NULL primary KEY AUTOINCREMENT,  tster string  NOT NULL,   ttt int  NOT NULL ); ")
if err != nil{
	fmt.Println(err)
	}
	data.Close()
	fmt.Println(statementsavedata)
	}