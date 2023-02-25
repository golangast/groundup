
package createserviceuser

import (

	. "app/db"
	"fmt"
)

func Createserviceuser(){
data, err := DbConnection() //create db instance
if err != nil {
	fmt.Println(err)
}

statementsavedata, err := data.Exec("CREATE TABLE IF NOT EXISTS serviceuser (  age int  NOT NULL,   name string  NOT NULL,   email string  NOT NULL ); ")
if err != nil{
	fmt.Println(err)
	}
	data.Close()
	fmt.Println(statementsavedata)
	}