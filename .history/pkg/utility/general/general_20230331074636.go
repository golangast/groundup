package general

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

func AppendStringToFiles(path, text string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

//creates the config folder/file
func CreateConfig() {

	if err := os.MkdirAll("../config", os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println("Whoops, could not create directory(ies) because", err)
	}

	mfile, err := os.Create("../config/persis.yaml")
	if isError(err) {
		fmt.Println("error -", err, mfile)
	}
	//use the ../ to make sure it gets out of /src
	var Configbase = `
app:
 app: "app.go"
 path: "../app"
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

func DBcreate() {

	if err := os.MkdirAll("../databaseconfig", os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println("Whoops, could not create directory(ies) because", err)
	}

	mfile, err := os.Create("../databaseconfig/dbpersis.yaml")
	if isError(err) {
		fmt.Println("error -", err, mfile)
	}
	var dbpersis = `
urls:
 home: "/home/"
 home: "/homes/"
 `
	/* write to the files */
	tm := template.Must(template.New("queue").Parse(dbpersis))
	err = tm.Execute(mfile, nil)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

}

func Getfields() string {
	viper.SetConfigName("dbpersis") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../databaseconfig/") // config file path
	viper.AutomaticEnv()                      // read value ENV variable
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	home := viper.GetString("urls.home")
	fmt.Println(home)

	return home
}

func ConfigAddEnv(f string, E string) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(f, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(E + " \n")
	if err != nil {
		log.Fatal(err)
	}

	// Save file changes.
	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Updated Successfully.")
}

func Readconfig(f string, configvalues string) string {
	str := FindTextNReturn(f, configvalues)

	return str

}
func ConfigAddFile(f string, E string) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(f, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write some text line-by-line to file.
	err = AppendStringToFiles(f, E)
	if err != nil {
		log.Fatal(err)
	}
	err = AppendStringToFiles(f, " \n")
	if err != nil {
		log.Fatal(err)
	}
	// Save file changes.
	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Updated Successfully.")
}
func CreateServ(p string, filename string) *os.File {
	pleft := strings.Split(p, "/")[1]
	if err := os.MkdirAll(pleft+"/serv", os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println("Whoops, could not create directory(ies) because", pleft+"/serv")
	}
	//create server file
	f, err := os.Create(pleft + "/" + pleft + "/serv" + filename + ".go")
	if isError(err) {
		fmt.Println("error -", err)
	}
	return f
}
func Noescape(str string) template.HTML {
	return template.HTML(str)
}

//creates the main, server, template
func CreateBase(p string, filename string) (*os.File, *os.File) {
	//split p to get the directory
	pleft := strings.Split(p, "/")[1]
	if err := os.MkdirAll(pleft, os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println("Whoops, could not create directory(ies) because", err)
	}
	if err := os.MkdirAll(pleft+"/templates", os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println("Whoops, could not create directory(ies) because", err)
	}
	//create file
	var tfile *os.File

	//making main.go
	mfile, err := os.Create(pleft + "/main.go")
	if isError(err) {
		fmt.Println("error -", err)
	}

	//making template director
	tfile, err = os.Create(pleft + "/templates/index.html")
	if isError(err) {
		fmt.Println("error -", err)
	}

	return mfile, tfile

}

//create dir and file of the database/model
func CreateDB(p string, fn string) {
	//split p to get the directory
	pleft := strings.Split(p, "/")[1]
	//make directory
	if err := os.MkdirAll(pleft+"/db", os.ModeSticky|os.ModePerm); err != nil {
		//	fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println(pleft + "/db already created")
	}
	_, err := os.Create(pleft + "/db/" + "db" + fn + ".go")
	if isError(err) {
		fmt.Println("already created")
	}
}

//create dir and file of the database/model
func CreateRoute(p string, fn string) *os.File {
	//split p to get the directory
	pleft := strings.Split(p, "/")[1]
	//make directory
	if err := os.MkdirAll(pleft+"/route", os.ModeSticky|os.ModePerm); err != nil {
		//	fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println(pleft + "/route already created")
	}
	r, err := os.Create(pleft + "/route/" + "route" + fn + ".go")
	if isError(err) {
		fmt.Println("already created")
	}
	return r

}

//https://gist.github.com/mastef/05f46d3ab2f5ed6a6787
func Deletefile(t string) {
	e := os.Remove(t)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(e, "was deleted")
}
func Getcom(b *bufio.Reader) string {

	texts, _ := b.ReadString('\n')
	s := strings.TrimSpace(texts)
	s = strings.Replace(s, "\n", "", -1)
	return s
}
func ReadFile(p string) string {
	// Open file for reading.
	var file, err = os.OpenFile(p, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			fmt.Println(err)

		}
	}

	fmt.Println("Reading from file.")
	fmt.Println(string(text))
	s := string(text)
	return s
}

func Gentextcomment(file, updatetext, lookuptext string) error {
	found, err := FindText(file, updatetext)
	if err != nil {
		return err
	}
	if !found {
		err := UpdateText(file, lookuptext, updatetext+"\n"+lookuptext)
		if err != nil {
			return err
		}
	}
	return nil
}

func FindText(p, str string) (bool, error) {
	// Open file for reading.
	var file, err = os.OpenFile(p, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)
		if err != nil {
			return false, err
		}
		if strings.Contains(string(text), str) {
			return true, nil

		}
		// Break if finally arrived at end of file
		if err == io.EOF {
			if err != nil {
				return false, err
			}
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			if err != nil {
				return false, err
			}
			fmt.Println(err)

		}
	}

	file.Close()
	return false, nil
}
func FindTextNReturn(p, str string) string {
	// Open file for reading.
	var file, err = os.OpenFile(p, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	toplevel := TrimDot(str)
	property := TrimDotright(str)
	strs := strings.Replace(property, ".", " ", 1)
	// fmt.Println(str)
	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		if strings.Contains(string(text), toplevel) {
			//is the dot string and split it
			if strings.Contains(string(text), strs) {
				return string(text)
			}
		}
		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			fmt.Println(err)

		}
	}

	// fmt.Println("Reading from file.")
	fmt.Println(string(text))

	return ""
}
func WriteFile(f string) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(f, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Save file changes.
	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File Updated Successfully.")
}

func Replace(o string, n string) {
	err := os.Rename(o, n)
	if err != nil {
		log.Fatal(err)
	}
}

func Copy(o string, n string) {
	sourceFile, err := os.Open(o)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create(n)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesCopied)
}

//file logger
func FInfo(f os.FileInfo) {

	fmt.Println("File Name:", f.Name())        // Base name of the file
	fmt.Println("Size:", f.Size())             // Length in bytes for regular files
	fmt.Println("Permissions:", f.Mode())      // File mode bits
	fmt.Println("Last Modified:", f.ModTime()) // Last modification time
	fmt.Println("Is Directory: ", f.IsDir())   // Abbreviation for Mode().IsDir()
}

func AddZip(filename string, zipw *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Failed to open %s: %s", filename, err)
	}
	defer file.Close()

	wr, err := zipw.Create(filename)
	if err != nil {
		msg := "Failed to create entry for %s in zip file: %s"
		return fmt.Errorf(msg, filename, err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("Failed to write %s to zip: %s", filename, err)
	}

	return nil
}
func ZipUp(z string, a string, b string, c string) {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(z, flags, 0644)
	if err != nil {
		log.Fatalf("Failed to open zip for writing: %s", err)
	}
	defer file.Close()

	var files = []string{a, b, c}

	zipw := zip.NewWriter(file)
	defer zipw.Close()

	for _, filename := range files {
		if err := AddZip(filename, zipw); err != nil {
			log.Fatalf("Failed to add file %s to zip: %s", filename, err)
		}
	}
}
func ZipOpen(z string, d string) {
	zipReader, _ := zip.OpenReader(z)
	for _, file := range zipReader.Reader.File {

		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		targetDir := d
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		if file.FileInfo().IsDir() {
			log.Println("Directory Created:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			log.Println("File extracted:", file.Name)

			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func ScanWords(f string, o string, n string) {

	input, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte(o), []byte(n), -1)

	if err = ioutil.WriteFile(f, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//f is for file, o is for old text, n is for new text
func UpdateText(f string, o string, n string) error {
	input, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
	}

	output := bytes.Replace(input, []byte(o), []byte(n), -1)

	if err = ioutil.WriteFile(f, output, 0666); err != nil {
		fmt.Println(err)
	}

	return nil
}
func PWD() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}
func Tree() {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
func FTree(f string) {
	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range list {
		fmt.Println(name)
	}
}
func readFiles(dir string) ([]string, error) {
	fil, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer fil.Close()
	// return fil.Readdirnames(-1)
	// return fil.Readdirnames(1024) // doesn't leak
	return fil.Readdirnames(8 * 1024)
}
func TrimStringFromDot(s string) string {
	if idx := strings.Index(s, "."); idx != -1 {
		return s[:idx]
	}
	return s
}
func TrimStringFromDash(s string) string {
	if idx := strings.Index(s, "-"); idx != -1 {
		return s[idx:]
	}
	return s
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
func TrimDash(s string) string {
	if idx := strings.Index(s, "-"); idx != -1 {
		return s[:idx]
	}
	return s
}

func GetPropValue(prop string) []string {
	var str string
	var strprop string
	var totalsd []string

	s := strings.Split(prop, " ")
	for _, ss := range s {
		str = strings.Replace(TrimStringFromDot(ss), ".", " ", 1)
		strprop = strings.Replace(TrimStringFromDash(ss), "-", "", 1)
		totalsd = append(totalsd, str)
		totalsd = append(totalsd, strprop)
	}
	return totalsd
}
func GetPropDatatype(prop string) []string {
	var property []string
	var strright string

	s := strings.Split(prop, " ")
	for _, ss := range s {
		property = append(property, TrimDot(ss))
		strright = strings.Replace(TrimDotright(ss), ".", "", 1)
		property = append(property, TrimDash(strright))
	}
	return property
}
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
func Shellout(command string) (error, string, string) {
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
	return err, stdout.String(), stderr.String()
}
func ShelloutBash(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var cmd *exec.Cmd

	cmd = exec.Command("bash", "-c", command)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
func Startprograms(command string) (error, string, string, *exec.Cmd) {
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

//creates a slice of slices
func GetSubslice(s []string) [][]string {
	var j int
	var ss [][]string

	for i := 0; i < len(s); i += 2 {
		j += 2
		if j > len(s) {
			j = len(s)
		}

		ss = append(ss, s[i:j])
		fmt.Println(ss)
	}
	return ss
}

//takes a slice and incrementally adds values by 2
func SepProp(s []string) []string {
	result := make([]string, 0, len(s)/2)
	for i := 1; i < len(s); i += 2 {
		result = append(result, s[i-1]+" "+s[i])
	}
	return result
}

//takes a slice and incrementally adds values by 2 and adds commas
func SepCommaProp(s []string) []string {
	result := make([]string, 0, len(s)/2)
	values := make([]string, 0, len(s)/2)
	for i := 1; i < len(s); i += 2 {
		values = append(values, string('"')+s[i]+string('"'))
		result = append(result, s[i-1]+":"+",")
	}

	return result
}
func SeparateCommaProp(s []string) ([]string, []string) {
	prop := make([]string, 0, len(s)/2)
	values := make([]string, 0, len(s)/2)
	for i := 1; i < len(s); i += 2 {
		if IsNumeric(s[i]) {
			values = append(values, s[i]+",")
		} else {
			values = append(values, string('"')+s[i]+string('"')+",")
		}
		prop = append(prop, s[i-1]+":")
	}

	return prop, values
}
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func Watch() *exec.Cmd {

	files, _ := ioutil.ReadDir("/app")

	var newestFile string
	var newestTime int64 = 0
	for _, f := range files {
		fi, err := os.Stat("/app" + f.Name())
		if err != nil {
			fmt.Println(err)
		}
		currTime := fi.ModTime().Unix()
		if currTime > newestTime {
			newestTime = currTime
			newestFile = f.Name()
		}
	}
	fmt.Println("~~~~~~~~~~~~~~~", newestFile)
	curr_wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err, out, errout, cmd := Startprograms(`cd app && go run .`)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(filepath.Base(curr_wd))
	fmt.Println(out)
	fmt.Println("--- errs ---")
	fmt.Println(errout)
	return cmd
}

//get list of running processes
// func CheckProcessLife() {
// 	const processEntrySize = 568
// 	h, e := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
// 	if e != nil {
// 		panic(e)
// 	}
// 	p := windows.ProcessEntry32{Size: processEntrySize}
// 	for {
// 		e := windows.Process32Next(h, &p)
// 		if e != nil {
// 			break
// 		}
// 		s := windows.UTF16ToString(p.ExeFile[:])
// 		println(s)
// 	}
// }

// kill a process
func killProcessByName(procname string) int {
	kill := exec.Command("taskkill", "/im", procname, "/T", "/F")
	err := kill.Run()
	if err != nil {
		return -1
	}
	return 0
}
