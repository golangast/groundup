package home

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	. "github.com/golangast/groundup/dashboard/configutil/createconfig"
	. "github.com/golangast/groundup/dashboard/configutil/createserver"
	. "github.com/golangast/groundup/dashboard/dbsql/deletebytitle"
	. "github.com/golangast/groundup/dashboard/dbsql/getallcss"
	. "github.com/golangast/groundup/dashboard/dbsql/getpage"
	. "github.com/golangast/groundup/dashboard/ut"

	. "github.com/golangast/groundup/dashboard/dbsql/getallurls"

	. "github.com/golangast/groundup/dashboard/dbsql/dbutil"
	. "github.com/golangast/groundup/dashboard/dbsql/getlib"

	. "github.com/golangast/groundup/dashboard/handler/home/handlerutil"
	. "github.com/golangast/groundup/dashboard/watcher"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	//var u []geturls.Urls
	//var lib string
	var Stat Stats
	var err error
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
		GenerateTable()
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
		AddLibtoFile(pp, lib)
		if footer == "footer" {
			AddLibtoFilebyTitle(lib, footer)
		}

	case "observe": //*observe app process
		exe, path, pid, size, parent, threads, usage, alloc, totalAlloc, sys, numGC := Observe()
		Stat = Stats{Appexe: exe, Apppath: path, Apppid: pid, Appsize: size, Appparent: parent, Appthreads: threads, Appusage: usage, Alloc: alloc, Totalalloc: totalAlloc, Sys: sys, Numgc: numGC}
	}
	//load all the data.
	css := Getallcss()
	l := GetAllLib()
	u := GetUrls()
	d := Data{U: u, L: l, C: css, S: Stat}
	return c.Render(http.StatusOK, "home.html", d)
}

type Data struct {
	U []Urls
	L []Library
	C []CSS
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
