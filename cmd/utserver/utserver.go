package utserver

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"text/template"

	. "gitlab.com/zendrulat123/groundup/cmd/templates"
	"gitlab.com/zendrulat123/groundup/cmd/ut"
)

//p-path f-file s-script
//tie the viper config vars to params
func CreateServer(p string, f string, s string) {
	/* generate code */
	if err := os.MkdirAll(p, os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println("Whoops, could not create directory(ies) because", err)
	}

	if err := os.MkdirAll(p+"/templates", os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println("Whoops, could not create directory(ies) because", err)
	}

	mfile, err := os.Create(p + "/templates/" + f)
	if isError(err) {
		fmt.Println("error -", err, mfile)
	}
	sfile, err := os.Create(p + "/main.go")
	if isError(err) {
		fmt.Println("error -", err, sfile)
	}
	/* write to the files */
	tm := template.Must(template.New("queue").Parse(Maintemp))
	err = tm.Execute(sfile, nil)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	// //run the program
	// mydir, err := os.Getwd()

	// fmt.Println(mydir)

	err, out, errout := ut.Shellout("cd app && go mod init " + p + "&& go mod tidy && go mod vendor && go install && go build")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(out)
	fmt.Println("--- errs ---")
	fmt.Println(errout)

	err, outs, errouts := ut.Shellout("cd app && go run main.go")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(outs)
	fmt.Println("--- errs ---")
	fmt.Println(errouts)

	openbrowser("http://localhost:3000/")
}
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

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
