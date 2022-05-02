package home

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	. "github.com/zendrulat123/groundup/cmd/ut"
	. "github.com/zendrulat123/groundup/cmd/watcherut"

	db "github.com/zendrulat123/groundup/dashboard/db"
	kdb "github.com/zendrulat123/groundup/dashboard/db/kval"
	"github.com/zendrulat123/groundup/zegmarkup/utfsg"

	. "github.com/zendrulat123/groundup/dashboard/configutil/createserver"
	. "github.com/zendrulat123/groundup/dashboard/handler/home/handlerutil"
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
	var titles []string
	var urls []string
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
		titles, urls = utfsg.GenRoutes()
		fmt.Println("before-", titles, urls)
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
		titles, urls = db.GetAllkv("urls")
		libs, libtag = kdb.Getall("libs")
	case deletev == "true": //delete routes
		/*
		*	Delete route
		 */
		titletrim := strings.ReplaceAll(title, " ", "")
		db.DeleteDB("urls", title)
		err := os.RemoveAll("app/templates/" + titletrim + ".html")
		if err != nil {
			log.Fatal(err)
		}
		c.Redirect(http.StatusFound, "/home")
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
	type Data struct {
		Titles []string
		Urls   []string
		Libs   []string
		Libtag []string
		S      Stats
	}

	d := Data{Titles: titles, Urls: urls, Libs: libs, Libtag: libtag, S: Stat}
	return c.Render(http.StatusOK, "home.html", d)
}
