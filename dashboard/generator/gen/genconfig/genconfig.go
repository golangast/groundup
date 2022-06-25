package genconfig

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

//makes config file for my new format
func Make(dir string) {
	if err := os.MkdirAll(dir, os.ModeSticky|os.ModePerm); err != nil {
		fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
	} else {
		fmt.Println("Whoops, could not create directory(ies) because", err)
	}

	mfile, err := os.Create(dir + "/dbpersis.fsg")
	if err != nil {
		fmt.Println("error -", err, mfile)
	}

	var dbpersis = `
	urls{
		home:"/home",
		aboutme:"/aboutme",
		}
 `
	/* write to the files */
	tm := template.Must(template.New("queue").Parse(dbpersis))
	err = tm.Execute(mfile, nil)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

}

//creates a title markup
func Writetitle(f string, e string) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(f, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write some text line-by-line to file.
	err = AppendStringToFile(f, e)
	if err != nil {
		log.Fatal(err)
	}
	err = AppendStringToFile(f, " \n")
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

//adds string to file
func AppendStringToFile(path, text string) error {
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
func UpdateTexts(f string, o string, n string) {
	fmt.Println(f, o, n)
	input, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte(o), []byte(n), -1)

	fmt.Println("file: ", f, " old: ", o, " new: ", n)

	if err = ioutil.WriteFile(f, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetValue(files string, title string, value string) string {

	fileOpen, _ := os.Open(files)
	defer fileOpen.Close()
	var ss string
	scanner := bufio.NewScanner(fileOpen)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), title+"{") {
			scanner.Scan()
			if strings.Contains(scanner.Text(), value) {
				ss = TrimColonRight(scanner.Text())

			}
		}
	}
	return ss
}
func GetValues(files string, title string, value string) []string {
	var urls []string
	fileOpen, _ := os.Open(files)
	defer fileOpen.Close()
	var ss string
	scanner := bufio.NewScanner(fileOpen)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), title+"{") {
			scanner.Scan()
			if strings.Contains(scanner.Text(), value) {
				ss = TrimColonRight(scanner.Text())
				urls = append(urls, ss)
			}
		}
	}
	return urls
}
func TrimColonRight(s string) string {
	if idx := strings.Index(s, ":"); idx != -1 {
		return s[idx+1:]
	}
	return s
}
func DeleteLine(files string, value string) {
	input, err := ioutil.ReadFile(files)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, value) {
			lines[i] = ""
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(files, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
func UpdateKey(files string, title string, key string, value string) {
	input, err := ioutil.ReadFile(files)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, title) {
			if strings.Contains(lines[i+1], key) {
				values := TrimColonRight(lines[i+1])
				v := value + ":" + values
				lines[i+1] = v
			}
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(files, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
func AddRoute(files string, title string, key string, value string) {
	input, err := ioutil.ReadFile(files)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, title) {
			if line != "" {
				lines[i] = lines[i] + "\n"
				output := strings.Join(lines, "\n")
				err = ioutil.WriteFile(files, []byte(output), 0644)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}

	liness := strings.Split(string(input), "\n")
	for i, line := range liness {
		if strings.Contains(line, title) {
			v := key + ":" + value + ","
			lines[i+1] = v
			outputs := strings.Join(lines, "\n")
			err = ioutil.WriteFile(files, []byte(outputs), 0644)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

}
