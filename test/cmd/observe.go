/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows"
)

// observeCmd represents the observe command
var observeCmd = &cobra.Command{
	Use:   "observe",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		exe, path, pid, _ := CheckExeExists("app", "app.exe")
		wexe, wpath, wpid, _ := CheckExeExists("cmd", "watcher.exe")
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		// For info on each, see: https://golang.org/pkg/runtime/#MemStats
		fmt.Printf("---APP Running---\n")
		fmt.Printf("exe:'%s' path:'%s', pid:'%d'\n", exe, path, pid)
		fmt.Printf("---Watcher Running---\n")
		fmt.Printf("exe:'%s' path:'%s', pid:'%d'\n", wexe, wpath, wpid)
		fmt.Printf("---System Stats---\n")
		fmt.Printf("alloc:'%d' Mib, TotalAlloc:'%d' Mib, Sys:'%d' Mib,  NumGC:'%d' Mib\n", m.Alloc, m.TotalAlloc, m.Sys, m.NumGC)

		WatchSignals()

	},
}

func init() {
	rootCmd.AddCommand(observeCmd)

}

func CheckExeExists(exepath string, appexe string) (string, string, int, *os.Process) {
	path, err := exec.LookPath(exepath)
	if err != nil {
		fmt.Printf("didn't find '%s' executable\n", exepath)
	}

	pid, e := processID(appexe)
	if e != nil {
		fmt.Println("cant find pid for app.exe", e)
	}
	proc, err := os.FindProcess(int(pid))
	if err == nil {
		proc.Signal(syscall.SIGHUP)
	}

	return exepath, path, proc.Pid, proc
}
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
	// return 0, fmt.Errorf("%q not found", name)
}
