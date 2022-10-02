package cliutility

import (
	"fmt"
	"log"
	"os"
	"text/template"

	. "github.com/golangast/groundup/src/dashboard/generator/generatorutility"
	. "github.com/golangast/groundup/src/dashboard/ut"

	. "github.com/golangast/groundup/src/dashboard/generator/templates/db/dbconn"
	. "github.com/golangast/groundup/src/dashboard/generator/templates/db/getdata"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

//creates the config folder/file
func CreateConfig() {

	if err := os.MkdirAll("config", os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println("Whoops, could not create directory(ies) because", err)
	}

	mfile, err := os.Create("config/persis.yaml")
	if isError(err) {
		fmt.Println("error -", err, mfile)
	}
	var Configbase = `
app:
 app: "app.go"
 path: "app"
 file: "home.html"
 script: "jquery"`
	/* write to the files */
	tm := template.Must(template.New("queue").Parse(Configbase))
	err = tm.Execute(mfile, nil)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	defer mfile.Close()
}

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
