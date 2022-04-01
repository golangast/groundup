package serverutil

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"

	"gitlab.com/zendrulat123/groundup/cmd/ut"
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
