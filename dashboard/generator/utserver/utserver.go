package utserver

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	. "github.com/golangast/groundup/dashboard/generator/templates/body"
	. "github.com/golangast/groundup/dashboard/generator/templates/db/dbconn"
	. "github.com/golangast/groundup/dashboard/generator/templates/footer"
	. "github.com/golangast/groundup/dashboard/generator/templates/header"
	. "github.com/golangast/groundup/dashboard/generator/templates/server"

	"github.com/golangast/groundup/dashboard/ut"
)

//p-path f-file s-script
//tie the viper config vars to params
func CreateServer(p string, f string, s string, g string) {

	/* create folders*/
	if err := os.MkdirAll(p, os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("~~~~could not create"+p, err)
	} else {
		fmt.Println("Directory " + p + " successfully created with sticky bits and full permissions")
	}
	if err := os.MkdirAll(p+"/templates", os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("~~~~could not create"+p+"/templates", err)
	} else {
		fmt.Println("Directory " + p + "/templates successfully created with sticky bits and full permissions")
	}
	if err := os.MkdirAll("db", os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println("Whoops, could not create directory(ies) because", err)
	}

	/* create files*/
	bfile, err := os.Create(p + "/templates/home.html")
	if isError(err) {
		fmt.Println("error -", err, bfile)
	}
	hfile, err := os.Create(p + "/templates/header.html")
	if isError(err) {
		fmt.Println("error -", err, hfile)
	}
	ffile, err := os.Create(p + "/templates/footer.html")
	if isError(err) {
		fmt.Println("error -", err, ffile)
	}
	sfile, err := os.Create(p + "/" + g)
	if isError(err) {
		fmt.Println("error -", err, sfile)
	}
	fdb, err := os.Create("db/database.db")
	if isError(err) {
		fmt.Println("error -", err, fdb)
	}

	/* generate code in files*/
	tms := template.Must(template.New("queue").Parse(Servertemp))
	err = tms.Execute(sfile, nil)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	tmh := template.Must(template.New("queue").Parse(Headertemp))
	//all of this is needed to parse {{define header}} and {{end}}
	m := make(map[string]string)
	header := fmt.Sprintf(`{{define "header"}}%s`, "")
	end := fmt.Sprintf(`{{end}}%s`, "")
	m["header"] = header
	m["end"] = end
	err = tmh.Execute(hfile, m)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	tmf := template.Must(template.New("queue").Parse(Footertemp))
	//all of this is needed to parse {{define footer}} and {{end}}
	mf := make(map[string]string)
	footer := fmt.Sprintf(`{{define "footer"}}%s`, "")
	endf := fmt.Sprintf(`{{end}}%s`, "")
	mf["footer"] = footer
	mf["end"] = endf
	err = tmf.Execute(ffile, mf)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	tmb := template.Must(template.New("queue").Parse(Bodytemp))
	//all of this is needed to parse {{template footer .}} and {{template header .}}
	mb := make(map[string]string)
	headerb := fmt.Sprintf(`{{template "header" .}}%s`, "")
	footerb := fmt.Sprintf(`{{template "footer" .}}%s`, "")
	mb["footer"] = footerb
	mb["header"] = headerb
	err = tmb.Execute(bfile, mb)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	dbmb := template.Must(template.New("queue").Parse(Dbconntemp))
	err = dbmb.Execute(fdb, mb)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	err, out, errout := ut.Shellout("cd app && go mod init " + p + "&& go mod tidy && go mod vendor && go install && go build")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(out)
	fmt.Println("--- errs ---")
	fmt.Println(errout)

	err, outs, errouts := ut.Shellout("cd app && go install github.com/labstack/echo/v4 && go install github.com/labstack/echo/v4/middleware && go get github.com/labstack/gommon/log && go mod tidy && go mod vendor && go install && go build")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(outs)
	fmt.Println("--- errs ---")
	fmt.Println(errouts)
	bfile.Close()
	hfile.Close()
	ffile.Close()
	sfile.Close()
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
func GetUrlTitle(prop []string) ([]string, []string) {
	var title []string
	var url []string
	for _, ss := range prop {
		title = append(title, TrimColan(ss))
		url = append(url, TrimColanright(ss))

	}
	return title, url
}
func TrimDot(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[:idx]
	}
	return s
}

func TrimDotright(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[idx:]
	}
	return s
}
func TrimColan(s string) string {
	if idx := strings.Index(s, ":"); idx != -1 {
		return s[:idx]
	}
	return s
}

func TrimColanright(s string) string {
	if idx := strings.Index(s, ":"); idx != -1 {
		id := strings.Replace(s[idx:], ":", "", 1)
		return id
	}
	return s
}

func Createtemplatefile(f string) {
	mfile, err := os.Create("templates/" + f)
	if isError(err) {
		fmt.Println("error -", err, mfile)
	}
}
