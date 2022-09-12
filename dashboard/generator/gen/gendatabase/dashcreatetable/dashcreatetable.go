package dashcreatetable

import (
	"os"
	"reflect"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
	. "github.com/golangast/groundup/dashboard/generator/generatorutility"
	. "github.com/golangast/groundup/dashboard/generator/templates/db/createdb"
)

//creates database and files in the app
func Gendatasave(icb *DBFields) {
	//update the app's files with db connection
	if !FindText("app/app.go", "Createdb()") {
		UpdateText("app/app.go", "//#dbcall", `Createdb() `+"\n"+`//#dbcall`)
	}
	if !FindText("app/app.go", "app/db/createtable") {
		UpdateText("app/app.go", "//#import", `. "app/db/createtable" `+"\n"+`//#import`)
	}
	var ct *os.File
	//make the files and folders
	if _, err := os.Stat("app/db/createtable"); os.IsNotExist(err) {
		Makefolder("app/db/createtable")
		ct = Makefile("app/db/createtable/createtable.go")
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

	//store them
	for ii := 0; ii < len(d.Fields); ii++ {
		stringft := " " + d.Fields[ii] + " " + d.Types[ii] + " "
		d.Fieldtypes = append(d.Fieldtypes, stringft)
	}

	//save table and generate template in file
	WritetemplateData(DBcreate, ct, d)

}

//creates database and files in the app
func Generatedatabasefields(icb *DBFields, name string) {
	//update the app's files with db connection

	if !FindText("app/app.go", cases.Title(language.Und, cases.NoLower).String(name)+"()") {
		UpdateText("app/app.go", "//#dbcall", cases.Title(language.Und, cases.NoLower).String(name)+`() `+"\n"+`//#dbcall`)
	}
	if !FindText("app/app.go", "app/db/create"+name) {
		UpdateText("app/app.go", "//#import", `. "app/db/create`+name+`"`+"\n"+`//#import`)
	}
	var ct *os.File
	//make the files and folders
	if _, err := os.Stat("app/db/create" + name); os.IsNotExist(err) {
		Makefolder("app/db/create" + name)
		ct = Makefile("app/db/create" + name + "/create" + name + ".go")
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

	//store them
	for ii := 0; ii < len(d.Fields); ii++ {
		stringft := " " + d.Fields[ii] + " " + d.Types[ii] + " "
		d.Fieldtypes = append(d.Fieldtypes, stringft)
	}

	//save table and generate template in file
	WritetemplateData(DBcreate, ct, d)

}
