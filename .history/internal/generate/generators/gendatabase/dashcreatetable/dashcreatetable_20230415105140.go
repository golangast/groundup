package dashcreatetable

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	. "github.com/golangast/groundup/internal/dbsql/conn"
	. "github.com/golangast/groundup/internal/generate/templates/db/createdb"
	. "github.com/golangast/groundup/pkg/utility/general"
	. "github.com/golangast/groundup/pkg/utility/generate"
)

// creates database and files in the app
func Generatedatabasefields(icb *DBFields, name string) {
	//update the app's files with db connection
	//writes function to app.go

	err := Gentextcomment("../app/app.go", "Create"+name+"()", "//#createdb")
	ErrorCheck(err)

	err = Gentextcomment("../app/app.go", `. "app/db/create/create`+name+`"`, "//#import")
	ErrorCheck(err)

	err = Gentextcomment("../app/app.go", "Get"+name+"()", "//#getdatavar"+name)
	ErrorCheck(err)

	file, err := Filefolder("../app/db/create/create", name)
	ErrorCheck(err)

	var ct *os.File //keep file open till generate
	//make the files and folders
	if _, err := os.Stat("../app/db/create/create" + name); errors.Is(err, os.ErrNotExist) {
		Makefolder("../app/db/create/create" + name)
		ct, err = Makefile("../app/db/create/create" + name + "/create" + name + ".go")
		ErrorCheck(err)
	} else {
		fmt.Println("file exists")
	}

	//https://go.dev/play/p/0HrA-jrPZqG
	d := Data{}
	s := reflect.ValueOf(icb).Elem()
	//get table fields, types, and name
	for i := 0; i < s.NumField(); i++ {
		if s.Field(i).Interface() != "" {
			switch s.Type().Field(i).Name[0:1] {
			case "F": //fields
				d.Fields = append(d.Fields, s.Field(i).Interface().(string))
			case "T": //types
				d.Types = append(d.Types, s.Field(i).Interface().(string))
			case "S": //table name
				d.Table = s.Field(i).Interface().(string)
			}
		}
	}

	//store them in a slice
	for ii := 0; ii < len(d.Fields); ii++ {
		stringft := " " + d.Fields[ii] + " " + d.Types[ii] + " "
		d.Fieldtypes = append(d.Fieldtypes, stringft)
	}

	//save table and generate template in file
	WritetemplateData(DBcreates, ct, d)

}