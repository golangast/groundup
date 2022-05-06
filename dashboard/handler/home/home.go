package home

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	. "github.com/golangast/groundup/cmd/ut"
	. "github.com/golangast/groundup/cmd/watcherut"
	"github.com/labstack/echo/v4"

	db "github.com/golangast/groundup/dashboard/db"
	kdb "github.com/golangast/groundup/dashboard/db/kval"
	. "github.com/golangast/groundup/dashboard/dbsql/getallurls"
	"github.com/golangast/groundup/zegmarkup/utfsg"

	. "github.com/golangast/groundup/dashboard/configutil/createserver"
	. "github.com/golangast/groundup/dashboard/handler/home/handlerutil"
)

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

func Home(c echo.Context) error {
	var U []Urls
	var libs []string
	var libtag []string
	var Stat Stats
	//grab any route params
	con := c.Param("config")
	s := c.Param("server")
	prodv := c.Param("prodv")
	devv := c.Param("devv")
	stopv := c.Param("stopv")
	reload := c.Param("reloadv")
	genroutev := c.Param("genroutev")
	routesconfig := c.Param("routesconfigv")
	dbv := c.Param("dbv")
	showv := c.Param("showv")
	deletev := c.Param("deletev")
	libtagv := c.Param("libtagv")
	libtagsv := c.Param("libtagsv")
	title := c.Param("titlev")
	obsv := c.Param("obsv")
	//once params are grabbed then run methods
	switch {
	case con == "true":
		/*
		*	generate config
		 */

		fmt.Println("gonna config...")
		c.Redirect(http.StatusFound, "/home")
		CreateConfig()
	case s == "true":
		/*
		*	generate server
		 */
		fmt.Println("gonna serv...")
		c.Redirect(http.StatusFound, "/home")
		Createservers()
	case prodv == "true":
		/*
		* run production
		 */
		fmt.Println("gonna prod...")
		c.Redirect(http.StatusFound, "/home")
		Startprod()

	case devv == "true":
		/*
		*	run watcher
		 */
		fmt.Println("gonna watch...")
		c.Redirect(http.StatusFound, "/home")
		Watching()

	case stopv == "true": //stop application
		/*
		*	stop app
		 */
		c.Redirect(http.StatusFound, "/home")
		fmt.Println("gonna stop app...")
		c.Redirect(http.StatusFound, "/home")
		KillProcessByName("app.exe")
	case reload == "true": //reload application
		/*
		*	reload app
		 */
		fmt.Println("gonna reload...")
		c.Redirect(http.StatusFound, "/home")
		Reload()
	case routesconfig == "true": //create markup language for routes *not used any longer
		/*
		*	create config
		 */
		fmt.Println("gonna config...")
		c.Redirect(http.StatusFound, "/home")
		utfsg.Make("databaseconfig")
	case genroutev == "true": //generate routes
		/*
		*	generate routes
		 */
		fmt.Println("gonna gen routes...")
		c.Redirect(http.StatusFound, "/home")
		U = GetUrls()
		fmt.Println("before-", U)
	case dbv == "true": //generate database
		/*
		*	generate database
		 */
		c.Redirect(http.StatusFound, "/home")
		db.Tempfile()
		db.CreateBucket("urls", "home", "/home")
	case showv == "true": //show routes
		/*
		*	show routes
		 */
		U = GetUrls()
		fmt.Println("before-", U)
	case deletev == "true": //delete routes
		/*
		*	Delete route
		 */
		c.Redirect(http.StatusFound, "/home")
		titletrim := strings.ReplaceAll(title, " ", "")
		err := os.Remove("app/templates/" + titletrim + ".html")
		if err != nil {
			log.Fatal(err)
		}

		db.DeleteDB("urls", title)

	case libtagv == "true": //add lib
		/*
		*	add library
		 */
		//get title
		titletrim := strings.ReplaceAll(title, " ", "")
		//get library
		lib := kdb.GetValue("libs", libtagsv)
		path := filepath.FromSlash(`app/templates/` + titletrim + `.html`)
		pp := strings.Replace(path, "\\", "/", -1)
		AddLibtoFile(pp, lib)
		c.Redirect(http.StatusFound, "/home")

	case obsv == "true":
		/*
		*	observe app process
		 */
		exe, path, pid, size, parent, threads, usage, alloc, totalAlloc, sys, numGC := Observe()
		Stat = Stats{Appexe: exe, Apppath: path, Apppid: pid, Appsize: size, Appparent: parent, Appthreads: threads, Appusage: usage, Alloc: alloc, Totalalloc: totalAlloc, Sys: sys, Numgc: numGC}
	default:
		fmt.Println("none were used")
	}
	// type Data struct {
	// 	Titles []string
	// 	Urls   []string
	// 	Libs   []string
	// 	Libtag []string
	// 	S      Stats
	// }
	type Data struct {
		Urls   []Urls
		Libs   []string
		Libtag []string
		S      Stats
	}
	d := Data{Urls: U, Libs: libs, Libtag: libtag, S: Stat}

	return c.Render(http.StatusOK, "home.html", d)
}
