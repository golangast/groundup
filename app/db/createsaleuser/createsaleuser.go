
package createsaleuser

import (

	. "app/db"
	"fmt"
)

func Createsaleuser(){
data, err := DbConnection() //create db instance
if err != nil {
	fmt.Println(err)
}

statementsavedata, err := data.Exec("CREATE TABLE IF NOT EXISTS saleuser (  name string  NOT NULL,   age int  NOT NULL,   email string  NOT NULL ); ")
if err != nil{
	fmt.Println(err)
	}
	data.Close()
	fmt.Println(statementsavedata)
	}