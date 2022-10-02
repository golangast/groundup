package gendatabase

import (
	"fmt"
	"log"

	. "github.com/golangast/groundup/src/dashboard/generator/generatorutility"
	. "github.com/golangast/groundup/src/dashboard/ut"

	. "github.com/golangast/groundup/src/dashboard/generator/templates/db/dbconn"
	. "github.com/golangast/groundup/src/dashboard/generator/templates/db/getdata"
)

func Gendatabase(p string) {
	//make folder
	Makefolder(p)
	Makefolder(p + "/getdata")

	//make file/database
	Makefile(p + "/database.db")
	gdatafile := Makefile(p + "/getdata/getdata.go")
	dbfile := Makefile(p + "/db.go")

	//add database connections to server file
	AddDB("app/app.go", Grabdatatemp)

	//show data on html template file
	//Showdata("app/templates/newpage.html")

	//write it to the file
	Writetemplate(Dbconntemp, dbfile, nil)
	Writetemplate(Getdatatemp, gdatafile, nil)

	//pull down dependencies
	PullDowndb("app")
	Pulldowneverythingbase("app")
	err, out, errout := Shellout(`cd app && go mod tidy && go mod vendor`)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(out)
	fmt.Println("--- errs ---")
	fmt.Println(errout)

}
