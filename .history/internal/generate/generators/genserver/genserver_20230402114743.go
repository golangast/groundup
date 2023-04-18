package genserver

import (
	"fmt"
	"os"

	. "github.com/golangast/groundup/internal/dbsql/conn"
	. "github.com/golangast/groundup/internal/generate/templates/body"
	. "github.com/golangast/groundup/internal/generate/templates/footer"
	. "github.com/golangast/groundup/internal/generate/templates/header"
	. "github.com/golangast/groundup/internal/generate/templates/server"
	. "github.com/golangast/groundup/pkg/utility/generate"
	"github.com/spf13/viper"
)

func Createservers() {
	db, err := DbConnection()
	ErrorCheck(err)
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS urls (id integer NOT NULL primary key, urls text NOT NULL, titles text NOT NULL, lib text NOT NULL libtag text NOT NULL, css text NOT NULL, csstag text NOT NULL, filename text NOT NULL);`)
	ErrorCheck(err)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS library (id integer NOT NULL primary key, lib text NOT NULL, libtag text NOT NULL )`)
	ErrorCheck(err)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS csstable (id integer NOT NULL primary key, css text NOT NULL, csstag text NOT NULL )`)
	ErrorCheck(err)
	db.Close()

	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config") // config file path
	viper.SetConfigName("persis")    // config file name without extension
	viper.AutomaticEnv()             // read value ENV variable
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	g := viper.GetString("app.app")  //app name
	p := viper.GetString("app.path") //app path
	f := viper.GetString("app.file")
	s := viper.GetString("app.script")
	sp := fmt.Sprintf("%v", p)
	sf := fmt.Sprintf("%v", f)
	ss := fmt.Sprintf("%v", s)
	sg := fmt.Sprintf("%v", g)
	GenServer(sp, sf, ss, sg)
}

//p-path f-file s-script
//tie the viper config vars to params
func GenServer(p string, f string, s string, g string) {
	//header map for {{define "header"}} {{end}}
	m := make(map[string]string)
	header := fmt.Sprintf(`{{define "header"}}%s`, "")
	end := fmt.Sprintf(`{{end}}%s`, "")
	m["header"] = header
	m["end"] = end

	//footer map for {{define "footer"}} {{end}}
	mf := make(map[string]string)
	footer := fmt.Sprintf(`{{define "footer"}}%s`, "")
	endf := fmt.Sprintf(`{{end}}%s`, "")
	mf["footer"] = footer
	mf["end"] = endf

	//footer/header map for {{template "footer" .}} {{end}}
	mb := make(map[string]string)
	headerb := fmt.Sprintf(`{{template "header" .}}%s`, "")
	footerb := fmt.Sprintf(`{{template "footer" .}}%s`, "")
	mb["footer"] = footerb
	mb["header"] = headerb

	/* create folders*/
	Makefolder(p)
	Makefolder(p + "/templates")

	/* create files*/
	bfile, err := Makefile(p + "/templates/home.html")
	hfile, err := Makefile(p + "/templates/header.html")
	ffile, err := Makefile(p + "/templates/footer.html")
	sfile, err := Makefile(p + "/" + g)
	dbfile, err := Makefile(p + "db/database.db")

	/* write to files*/
	Writetemplate(Servertemp, sfile, nil)
	ErrorCheck(err)

	Writetemplate(Headertemp, hfile, m)
	ErrorCheck(err)

	Writetemplate(Footertemp, hfile, mf)
	ErrorCheck(err)

	Writetemplate(Bodytemp, bfile, mb)
	ErrorCheck(err)

	Pulldowneverything(p) //pulls dependencies and loads it

	bfile.Close()
	hfile.Close()
	ffile.Close()
	sfile.Close()
	dbfile.Close()
}
