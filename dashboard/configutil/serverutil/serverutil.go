package serverutil

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/zendrulat123/groundup/cmd/ut"
)

func Hotreload() *exec.Cmd {
	err, outs, errouts, cmd := ut.Startprogram("cd app && go run app.go")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(outs)
	fmt.Println("--- errs ---")
	fmt.Println(errouts)
	return cmd
}
func Startprod() *exec.Cmd {

	err, outs, errouts, cmd := ut.Startprogram("cd app && go mod tidy && go mod vendor && go install && go run app.go")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(outs)
	fmt.Println("--- errs ---")
	fmt.Println(errouts)
	openbrowser("http://localhost:3000/")
	return cmd
}
func Startdev() *exec.Cmd {
	var cmd *exec.Cmd
	ticker := time.NewTicker(time.Second * time.Duration(9000))
	for t := range ticker.C {
		fmt.Println("Tick at", t)
		cmd = ut.Watch()
	}
	defer ticker.Stop()
	fmt.Println("Ticker stopped")
	openbrowser("http://localhost:3000/")
	return cmd
}
func Stopping(cmd *exec.Cmd) {
	fmt.Println("....stopping 3000")
	if runtime.GOOS == "windows" {
		if err := cmd.Process.Kill(); err != nil {
			log.Fatal("failed to kill process: ", err)
		}
		err, out, errout := ut.Shellout(`TASKKILL /IM app.exe`)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		fmt.Println(out)
		fmt.Println("--- errs ---")
		fmt.Println(errout)
	} else {
		err, out, errout := ut.Shellout(`TASKKILL -9 app.exe`)
		if err != nil {
			log.Printf("error: %v\n", err)
		}
		fmt.Println(out)
		fmt.Println("--- errs ---")
		fmt.Println(errout)
	}
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
func Reload() {
	err, outs, errouts := ut.Shellout("cd app && go mod tidy && go mod vendor && go install && go build")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(outs)
	fmt.Println("--- errs ---")
	fmt.Println(errouts)
}
func KillProcessByName(procname string) int {
	kill := exec.Command("taskkill", "/im", procname, "/T", "/F")
	err := kill.Run()
	if err != nil {
		return -1
	}
	return 0
}
func Addthirdparty(p string, lib string, templatefile string) string {
	var pt = ""
	switch {
	case p == "local":
		GetThirdPartyFile(lib)
	case p == "cdn":

	}

	return pt
}
func GetThirdPartyFile(urlthirdparty string) {
	var fullURLFile = urlthirdparty

	// Build fileName from fullPath
	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]
	CreateFolder("app/thirdparties")
	// Create blank file
	file, err := os.Create("app/thirdparties/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)
}
func CreateFolder(folder string) {
	if err := os.MkdirAll(folder, os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println("Whoops, could not create directory(ies) because", err)
	}

}
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
