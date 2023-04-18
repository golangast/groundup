package gendatabase

import (
	"fmt"
	"log"

	. "github.com/golangast/groundup/pkg/utility/general"
	. "github.com/golangast/groundup/pkg/utility/generate"

	. "github.com/golangast/groundup/internal/generate/templates/db/dbconn"
	. "github.com/golangast/groundup/internal/generate/templates/db/getdata"
)

func Gendatabase(p string) {
	//make folder
	Makefolder(p)
	Makefolder(p + "/getdata")

	//make file/database
	Makefile(p + "/database.db")
	genappgetdata, err := Makefile(p + "/getdata/getdata.go")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	dbfile, err := Makefile(p + "/db/db.go")
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	//add database connections to server file

	AddDB(p+"/app.go", Grabdatatemp)

	//show data on html template file
	//Showdata("app/templates/newpage.html")

	//write it to the file
	Writetemplate(Dbconntemp, dbfile, nil)
	Writetemplate(Getdatatemp, genappgetdata, nil)

	//pull down dependencies
	PullDowndb("app")
	Pulldowneverythingbase("app")
	err, out, errout := Shellout(`cd .. && cd app && go mod tidy && go mod vendor`)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(out)
	fmt.Println("--- errs ---")
	fmt.Println(errout)

}
