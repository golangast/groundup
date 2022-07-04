package home

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	. "github.com/golangast/groundup/dashboard/dbsql/gettabledata"
	. "github.com/golangast/groundup/dashboard/dbsql/pagecreation/deletebytitle"
	. "github.com/golangast/groundup/dashboard/dbsql/pagecreation/deletebyurl"
	. "github.com/golangast/groundup/dashboard/dbsql/pagecreation/getallcss"
	. "github.com/golangast/groundup/dashboard/dbsql/pagecreation/getallurls"
	. "github.com/golangast/groundup/dashboard/dbsql/pagecreation/getlib"
	. "github.com/golangast/groundup/dashboard/dbsql/pagecreation/getpage"
	. "github.com/golangast/groundup/dashboard/generator/gen/genconfig"
	. "github.com/golangast/groundup/dashboard/generator/gen/gendatabase/createdatabase"

	. "github.com/golangast/groundup/dashboard/generator/gen/genserver"
	. "github.com/golangast/groundup/dashboard/handler/home/handlerutil"
	. "github.com/golangast/groundup/dashboard/ut"
	. "github.com/golangast/groundup/dashboard/watcher"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	var Stat Stats
	var err error
	var DBFields []DBFields
	//grab any route params
	m := c.Param("m")
	footer := c.Param("footer")
	libtagsv := c.Param("libtagsv")
	title := c.Param("titlev")
	//once params are grabbed then run methods
	switch m {
	case "config": //*generate config
		CreateConfig()
	case "server": //*generate server
		Createservers()
	case "production": //*run production
		Startprod()
	case "dev": //*run watcher
		Watching()
	case "stop": //*stop app
		KillProcessByName("app.exe")
		c.Redirect(http.StatusFound, "/show")
	case "reload": //*reload application
		Reload()
	case "routesconfig": //*create config
		Make("databaseconfig")
	case "db": //*generate database tables
		Gendatabase("app/db")
	case "show": //*show routes
	case "delete": //*delete routes
		titletrim := strings.ReplaceAll(title, " ", "")
		titles := GetPagetitle("/" + titletrim)
		err = Deletebytitle(titles)
		ErrorCheck(err)
		titletrimslash := strings.ReplaceAll(titletrim, "/", "")
		err := os.Remove("app/templates/" + titletrimslash + ".html")
		ErrorCheck(err)
	case "lib": //*add lib
		titletrim := strings.ReplaceAll(title, " ", "")
		lib := GetLib(libtagsv)
		path := filepath.FromSlash(`app/templates/` + titletrim + `.html`)
		pp := strings.Replace(path, "\\", "/", -1)
		AddLibtoFile(pp, lib, titletrim)
		if footer == "footer" {
			AddLibtoFilebyTitle(lib, footer)
		}
	case "observe": //*observe app process
		exe, path, pid, size, parent, threads, usage, alloc, totalAlloc, sys, numGC := Observe()
		Stat = Stats{Appexe: exe, Apppath: path, Apppid: pid, Appsize: size, Appparent: parent, Appthreads: threads, Appusage: usage, Alloc: alloc, Totalalloc: totalAlloc, Sys: sys, Numgc: numGC}

	case "showtable": //*show routes
		DBFields = Gettabledata()
	}

	//load all the data.
	css := Getallcss()
	l := GetAllLib()
	u := GetUrls()

	file_db_referentialintegrity(u)
	d := Data{U: u, L: l, C: css, F: DBFields, S: Stat}
	return c.Render(http.StatusOK, "home.html", d)
}

type Data struct {
	U []Urls
	L []Library
	C []CSS
	F []DBFields
	S Stats
}

type Stats struct {
	Appexe     string
	Apppath    string
	Apppid     string
	Appsize    string
	Appparent  string
	Appthreads string
	Appusage   string
	Alloc      string
	Totalalloc string
	Sys        string
	Numgc      string
}

func file_db_referentialintegrity(u []Urls) {
	for _, urls := range u {
		urlstrim := strings.ReplaceAll(urls.Urls, "/", "")
		fmt.Println(urlstrim)
		b := isDirectory(urlstrim + ".html")
		if !b {
			Deletebyurl(urlstrim)
		}
	}

}
func isDirectory(file string) bool {
	info, err := os.Stat("app/templates/" + file)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
