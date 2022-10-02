package cliutility

import (
	"fmt"
	"log"
	"os"
	"text/template"
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
