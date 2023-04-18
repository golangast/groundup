
package createggg

import (

	. "app/db"
	"fmt"
)

func Createggg(){
data, err := DbConnection() //create db instance
if err != nil {
	fmt.Println(err)
}

statementsavedata, err := data.Exec("CREATE TABLE IF NOT EXISTS ggg (  ggg string  NOT NULL,   ddd string  NOT NULL,   eee string  NOT NULL,   rrrr string  NOT NULL,   wwww string  NOT NULL ); ")
if err != nil{
	fmt.Println(err)
	}
	data.Close()
	fmt.Println(statementsavedata)
	}