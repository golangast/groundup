package watcherut

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"syscall"
	"time"

	. "github.com/golangast/groundup/dashboard/handler/home/handlerutil"

	"golang.org/x/sys/windows"
)

/*watcher basically watches files in the app directory
if the size changes then restart the app and the watcher*/

func Watching() {
	watchtime := 5 * time.Second
	path := "app/templates"
	var run bool
	//get folder size
	info, err := os.Lstat(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//check for sizes between times and then restart
	for {
		//check folder size between time
		sizes := diskUsage(path, info)
		time.Sleep(watchtime)
		sizetwo := diskUsage(path, info)
		//if not the same then kill process and restart
		if sizes != sizetwo {
			pid, e := processID("app.exe")
			if e != nil {
				fmt.Println("cant find pid", e)
				run = false
			} else {
				kill(int(pid))
				run = true
				if run {
					time.Sleep(2 * time.Second)
					Startapp()
				}
			}

		}

	}
}

func Startprogram(command string) (error, string, string, *exec.Cmd) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String(), cmd
}
func Startproduction() (*exec.Cmd, string, string) {

	err, outs, errouts, cmd := Startprogram("cd .. && cd app && go run app.go")
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(outs)
	fmt.Println("--- errs ---")
	fmt.Println(errouts)
	return cmd, outs, errouts
}

const processEntrySize = 568

func processID(name string) (uint32, error) {
	h, e := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if e != nil {
		return 0, e
	}
	p := windows.ProcessEntry32{Size: processEntrySize}
	for {
		e := windows.Process32Next(h, &p)
		if e != nil {
			return 0, e
		}
		if windows.UTF16ToString(p.ExeFile[:]) == name {
			return p.ProcessID, nil
		}
	}
}

func Restartapp(path string) {

	binary, err := exec.LookPath(path)
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}
	time.Sleep(1 * time.Second)
	args := []string{"go run app.go", ""}
	attr := syscall.ProcAttr{
		Env: os.Environ(),
	}
	pid, handle, err := syscall.StartProcess(binary, args, &attr)
	if err != nil {
		err = fmt.Errorf("Error in StartProcess: %s", err)
	} else {
		log.Println("Starting background server with pid=%d", pid, handle)
	}

}
func diskUsage(currentPath string, info os.FileInfo) int64 {
	size := info.Size()

	if !info.IsDir() {
		return size
	}

	dir, err := os.Open(currentPath)

	if err != nil {
		fmt.Println(err)
		return size
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.Name() == "." || file.Name() == ".." {
			continue
		}
		size += diskUsage(currentPath+"/"+file.Name(), file)
	}

	fmt.Printf("bytes:[%d] | path:[%s]\n", size, currentPath)

	return size
}

func restartapp(path string) bool {
	dir, err := os.Open(path)

	if err != nil {
		fmt.Println(err)

	}
	defer dir.Close()

	files, err := dir.Readdir(-1)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.Name() == "." || file.Name() == ".." {
			continue
		}
		fmt.Printf(path+"/"+file.Name(), file)

		if file.Name() == "app.exe" {
			Restartapp(path + "/" + file.Name())
			return true
		}

	}
	return true
}

func kill(pid int) error {
	kill := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(pid))

	kill.Stderr = os.Stderr
	kill.Stdout = os.Stdout
	fmt.Println(kill.Stderr, kill.Stdout)
	kill.Run()
	return nil
}
func CmdStart(args string) *exec.Cmd {

	err, outs, errouts, cmd := Startprogram(args)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(outs)
	fmt.Println("--- errs ---")
	fmt.Println(errouts)

	return cmd
}
