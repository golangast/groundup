package createdb

var DBcreate = `
package create{{.Table}}

import (

	. "app/db"
	"fmt"
)

func Create{{.Table}}(){
data, err := DbConnection() //create db instance
if err != nil {
	fmt.Println(err)
}

statementsavedata, err := data.Exec("CREATE TABLE IF NOT EXISTS {{.Table}} (id integer NOT NULL primary KEY AUTOINCREMENT, {{- range $val := Iterate .Fieldtypes }} {{ $val }} {{- end }}); ")
if err != nil{
	fmt.Println(err)
	}
	data.Close()
	fmt.Println(statementsavedata)
	}`
