package handlerutil

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	. "github.com/golangast/groundup/dashboard/dbsql/pagecreation/deletebyurl"
	. "github.com/golangast/groundup/dashboard/dbsql/pagecreation/getlib"
	. "github.com/golangast/groundup/dashboard/dbsql/pagecreation/getpage"

	"github.com/golangast/groundup/dashboard/ut"

	//"golang.org/x/sys/windows"
	ps "github.com/mitchellh/go-ps"
)

func Kill(pid int) error {
	kill := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(pid))

	kill.Stderr = os.Stderr
	kill.Stdout = os.Stdout
	fmt.Println(kill.Stderr, kill.Stdout)
	kill.Run()
	return nil
}
func Startprod() *exec.Cmd {

	err, outs, errouts, cmd := ut.Startprograms("cd app && go run app.go")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(outs)
	fmt.Println("--- errs ---")
	fmt.Println(errouts)
	openbrowser("http://localhost:3000/")
	return cmd
}
func Startapp() *exec.Cmd {
	fmt.Println("restarting app")

	err, outs, errouts, cmd := ut.Startprograms("cd app && go run app.go")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(outs)
	fmt.Println("--- errs ---")
	fmt.Println(errouts)
	//openbrowser("http://localhost:3000/")
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
	err, outs, errouts := ut.Shellout("pwd && cd app && go mod tidy && go mod vendor && go install && go build")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(outs)
	fmt.Println("--- errs ---")
	fmt.Println(errouts)
}
func KillProcessByName(procname string) {
	_, pid, _, err := ProcessID(procname)
	if err != nil {
		fmt.Println(err)
	}
	kill := exec.Command("taskkill", "/im", procname, "/T", "/F")
	err = kill.Run()
	if err != nil {
		fmt.Println(err)
	}

	proc, err := os.FindProcess(int(pid))
	if err != nil {
		log.Println(err)
	}
	// Kill the process
	if proc != nil {
		proc.Kill()
		exec.Command("taskkill", "/f", "/t", "/pid", strconv.Itoa(int(pid))).Run()

	}

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
	// Build fileName from fullPath
	fileURL, err := url.Parse(urlthirdparty)
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
	resp, err := client.Get(urlthirdparty)
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

func AddLibtoFile(path, lib, title string) {

	l := GetLib(lib)

	o := "<!-- ### -->"

	var n = `<script scr="` + l + `" ></script> ` + "\n" + o

	fmt.Println(path, o, n)

	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err, "file not found ", path)
		//if not found then print list of files and delete in database
		files, err := ioutil.ReadDir("app")
		if err != nil {
			fmt.Println(err)
		}
		for _, file := range files {
			fmt.Println(file.Name(), file.IsDir())
		}
		Deletebyurl(title)
		return
	}

	output := bytes.Replace(input, []byte(o), []byte(n), -1)

	fmt.Println("file: ", path, " old: ", o, " new: ", n)

	if err = ioutil.WriteFile(path, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func CheckExeExists(exepath string, appexe string) (string, string, string) {
	path, err := exec.LookPath(exepath)
	if err != nil {
		fmt.Printf("didn't find '%s' executable\n", exepath)
	}

	pid, errs := Getpid(appexe)
	if errs != nil {
		fmt.Println("cant find pid for app.exe", errs)
	}

	spid := fmt.Sprint(pid)
	return exepath, path, spid
}

// func GetProcData(exepath string, appexe string) (string, string, string, string, string, string, string, error) {
// 	path, err := exec.LookPath(exepath)
// 	if err != nil {
// 		fmt.Printf("didn't find '%s' executable\n", exepath)
// 	}
// 	pid, size, parent, threads, usage, e := processInfo(appexe)
// 	if e != nil {
// 		fmt.Println(e)
// 	}
// 	return exepath, path, pid, size, parent, threads, usage, nil
// }
func WatchSignals() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			//handle SIGINT
			fmt.Println("interrupt:", os.Interrupt)
		case syscall.SIGTERM:
			//handle SIGTERM
			fmt.Println("syscall.SIGTERM:", syscall.SIGTERM)
		}
	}()
}
func logSignal(p *os.Process, sig os.Signal) error {
	log.Printf("sending signal %s to PID %d", sig, p.Pid)
	err := p.Signal(sig)
	if err != nil {
		log.Print(err)
	}
	return err
}

// const processEntrySize = 568

// func processInfo(name string) (string, string, string, string, string, error) {
// 	h, e := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
// 	if e != nil {
// 		fmt.Println(e)
// 	}
// 	p := windows.ProcessEntry32{Size: processEntrySize}
// 	for {
// 		e := windows.Process32Next(h, &p)
// 		if e != nil {
// 			fmt.Println("didnt find it ", e)
// 			return "", "", "", "", "", nil
// 		}

// 		if windows.UTF16ToString(p.ExeFile[:]) == name {
// 			spid := fmt.Sprint(p.ProcessID)
// 			ssize := fmt.Sprint(p.Size)
// 			sparent := fmt.Sprint(p.ParentProcessID)
// 			sthreads := fmt.Sprint(p.Threads)
// 			susage := fmt.Sprint(p.Usage)
// 			return spid, ssize, sparent, sthreads, susage, nil
// 		}

// 	}

// }

// func ProcessID(name string) (uint32, error) {
// 	h, e := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
// 	if e != nil {
// 		return 0, e
// 	}
// 	p := windows.ProcessEntry32{Size: processEntrySize}
// 	for {
// 		e := windows.Process32Next(h, &p)
// 		if e != nil {
// 			return 0, e
// 		}
// 		if windows.UTF16ToString(p.ExeFile[:]) == name {
// 			return p.ProcessID, nil

// 		}
// 	}

// }

func Getpid(name string) (string, error) {
	list, err := ps.Processes()
	if err != nil {
		panic(err)
	}
	for _, p := range list {
		log.Printf("Process %s with PID %d and PPID %d", p.Executable(), p.Pid(), p.PPid())

		if p.Executable() == name {
			return strconv.Itoa(p.Pid()), err
		}
	}
	return "", err
}

func ProcessID(name string) (string, uint32, uint32, error) {
	list, err := ps.Processes()
	if err != nil {
		panic(err)
	}
	for _, p := range list {
		log.Printf("Process %s with PID %d and PPID %d", p.Executable(), p.Pid(), p.PPid())

		if p.Executable() == name {
			return p.Executable(), uint32(p.Pid()), uint32(p.Pid()), err
		}
	}
	return "", 0, 0, err

}

func Getpidstring(name string) (string, string, string, error) {
	list, err := ps.Processes()
	if err != nil {
		panic(err)
	}
	for _, p := range list {
		log.Printf("Process %s with PID %d and PPID %d", p.Executable(), p.Pid(), p.PPid())

		if p.Executable() == name {
			return p.Executable(), strconv.Itoa(p.Pid()), strconv.Itoa(p.Pid()), err
		}
	}
	return "", "", "", err

}

// func Observe() (string, string, string, string, string, string, string, string, string, string, string) {
// 	exepath, path, pid, size, parent, threads, usage, err := GetProcData("app", "app.exe")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	var m runtime.MemStats
// 	runtime.ReadMemStats(&m)
// 	Alloc := fmt.Sprint(m.Alloc)
// 	TotalAlloc := fmt.Sprint(m.TotalAlloc)
// 	Sys := fmt.Sprint(m.Sys)
// 	NumGC := fmt.Sprint(m.NumGC)
// 	WatchSignals()
// 	return exepath, path, pid, size, parent, threads, usage, Alloc, TotalAlloc, Sys, NumGC
// }

func AddLibtoFilebyTitle(lib, title string) {
	var path, filename string

	if title == "footer" {
		path = "app/templates/footer.html"
	} else {
		filename = GetPageFile(title)
		path = "app/templates/" + filename
	}

	o := "<!-- ### -->"

	var n = `<script scr="` + lib + `" ></script> ` + "\n" + o

	fmt.Println("app/templates/", o, n)

	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte(o), []byte(n), -1)
	fmt.Println("file: ", path, " old: ", o, " new: ", n)
	if err = ioutil.WriteFile(path, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
