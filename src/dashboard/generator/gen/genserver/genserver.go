package genserver

import (
	"fmt"
	"os"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
	. "github.com/golangast/groundup/dashboard/generator/generatorutility"
	. "github.com/golangast/groundup/dashboard/generator/templates/body"
	. "github.com/golangast/groundup/dashboard/generator/templates/footer"
	. "github.com/golangast/groundup/dashboard/generator/templates/header"
	. "github.com/golangast/groundup/dashboard/generator/templates/server"
	"github.com/spf13/viper"
)

func Createservers() {
	db, err := DbConnection()
	ErrorCheck(err)

	statementurls, err := db.Exec(`CREATE TABLE IF NOT EXISTS urls (id integer NOT NULL primary key, urls text NOT NULL, titles text NOT NULL, lib text NOT NULL libtag text NOT NULL, css text NOT NULL, csstag text NOT NULL, filename text NOT NULL);`)
	ErrorCheck(err)

	statementlibrary, err := db.Exec(`CREATE TABLE IF NOT EXISTS library (id integer NOT NULL primary key, lib text NOT NULL, libtag text NOT NULL )`)
	ErrorCheck(err)

	statementcss, err := db.Exec(`CREATE TABLE IF NOT EXISTS csstable (id integer NOT NULL primary key, css text NOT NULL, csstag text NOT NULL )`)
	ErrorCheck(err)

	db.Close()

	fmt.Println(statementurls, statementlibrary, statementcss)

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // config file path
	viper.SetConfigName("persis")   // config file name without extension
	viper.AutomaticEnv()            // read value ENV variable
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	conf := viper.ConfigFileUsed()
	fmt.Println(conf)
	// Declare var
	g := viper.GetString("app.app")
	p := viper.GetString("app.path")
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
	Makefolder(p + "db")
	// Makefolder(p + "dbsql")
	// Makefolder(p + "dbsql/nav")

	// Makefolder(p + "/templates/nav")

	/* create files*/
	bfile := Makefile(p + "/templates/home.html")
	hfile := Makefile(p + "/templates/header.html")
	ffile := Makefile(p + "/templates/footer.html")
	sfile := Makefile(p + "/" + g)
	dbfile := Makefile(p + "db/database.db")
	// navfile := Makefile(p + "dbsql/nav/nav.go")
	// fnav := Makefile("templates/nav/nav.html")

	/* write to files*/
	Writetemplate(Servertemp, sfile, nil)
	/*fix */
	// Writetemplate(Navtemp, navfile, nil)
	// Writetemplate(Navtemp, fnav, nil)
	Writetemplate(Headertemp, hfile, m)
	Writetemplate(Footertemp, hfile, mf)
	Writetemplate(Bodytemp, bfile, mb)

	Pulldowneverything(p) //pulls dependencies and loads it

	bfile.Close()
	hfile.Close()
	ffile.Close()
	sfile.Close()
	dbfile.Close()
	// fnav.Close()
}
