/*
The home handler basically has the first param in a
switch statement and that param is what is used
to call different functions of the program.
Instead of one handler per func, it is just the first param.
I did it this way because these functions need to run
on the same home.html file.
*/

package home

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	. "github.com/golangast/groundup/pkg/utility/general"
	. "github.com/golangast/groundup/pkg/utility/handler"
	. "github.com/golangast/groundup/services/dbsql/appdata/getapptablenames"
	. "github.com/golangast/groundup/services/dbsql/appdata/getapptables"
	"github.com/golangast/groundup/services/dbsql/conn"
	. "github.com/golangast/groundup/services/dbsql/deletetable"
	. "github.com/golangast/groundup/services/dbsql/gettabledata"
	. "github.com/golangast/groundup/services/dbsql/pagecreation/deletebytitle"
	. "github.com/golangast/groundup/services/dbsql/pagecreation/deletebyurl"
	. "github.com/golangast/groundup/services/dbsql/pagecreation/getallcss"
	. "github.com/golangast/groundup/services/dbsql/pagecreation/getallurls"
	. "github.com/golangast/groundup/services/dbsql/pagecreation/getlib"
	. "github.com/golangast/groundup/services/dbsql/pagecreation/getpage"
	. "github.com/golangast/groundup/services/generate/generators/genconfig"
	. "github.com/golangast/groundup/services/generate/generators/gendatabase/createdatabase"
	. "github.com/golangast/groundup/services/generate/generators/genserver"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {

	var Stat Stats
	var err error
	var DBFields []conn.DBFields
	//grab any route params
	m := c.Param("m")
	footer := c.Param("footer")
	libtagsv := c.Param("libtagsv")
	title := c.Param("titlev")
	table := c.Param("table")

	//once params are grabbed then run methods
	switch m {
	case "config": //*generate config
		CreateConfig()
	case "server": //*generate server
		Createservers()
	case "production": //*run production
		Startprod()

	case "stop": //*stop app
		if runtime.GOOS == "windows" {
			KillProcessByName("app.exe")
		} else {
			KillProcessByName("app")
		}
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

		var exe string
		var pid string
		var ppid string
		if runtime.GOOS == "windows" {
			exe, pid, ppid, err = Getpidstring("app.exe")
			ErrorCheck(err)
			Stat = Stats{Exe: exe, Pid: pid, Ppid: ppid}
		} else {
			exe, pid, ppid, err = Getpidstring("app")
			ErrorCheck(err)
			Stat = Stats{Exe: exe, Pid: pid, Ppid: ppid}
		}
	case "removetable": //*remove table process
		err := Deletetable(table)
		if err != nil {
			fmt.Println(err)
		}

	}

	/*
		This loads the data from the database
		like when you want to see all the js libs on the pages etc..
	*/
	tablenames := Getapptabledata()
	css := Getallcss()
	l := GetAllLib()
	u := GetUrls()
	file_db_referentialintegrity(u)
	DBFields = Gettabledata()
	alltabledata := Getapptables()
	//this is just to package the data and send it to the template
	d := Data{U: u, L: l, C: css, F: DBFields, S: Stat, tnames: tablenames, Alltabledata: alltabledata}
	return c.Render(http.StatusOK, "home.html", d)
}

type Data struct {
	U            []Urls          //urls
	L            []Library       //js libs
	C            []CSS           //css
	F            []conn.DBFields //table fields
	S            Stats           //pids for the program
	tnames       []string
	Alltabledata []TableData
	//Alltabledata string
}

type Stats struct {
	Exe  string
	Pid  string
	Ppid string
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
