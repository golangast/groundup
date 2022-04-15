package home

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	. "github.com/zendrulat123/groundup/cmd/ut"
	db "github.com/zendrulat123/groundup/dashboard/db"
	kdb "github.com/zendrulat123/groundup/dashboard/db/kval"
	"github.com/zendrulat123/groundup/zegmarkup/utfsg"

	. "github.com/zendrulat123/groundup/dashboard/configutil/createserver"
	. "github.com/zendrulat123/groundup/dashboard/handler/home/handlerutil"
)

func Home(c echo.Context) error {
	var titles []string
	var urls []string
	var libs []string
	var libtag []string

	var cmd *exec.Cmd
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
	hotloadv := c.Param("hotloadv")
	deletev := c.Param("deletev")
	libtagv := c.Param("libtagv")
	libtagsv := c.Param("libtagsv")
	title := c.Param("titlev")

	//once params are grabbed then run methods
	switch {
	case con == "true": //create config
		fmt.Println("gonna config...")
		c.Redirect(http.StatusFound, "/home")
		CreateConfig()
	case s == "true": //create server
		fmt.Println("gonna serv...")
		c.Redirect(http.StatusFound, "/home")
		Createservers()
	case prodv == "true": //run production
		fmt.Println("gonna prod...")
		c.Redirect(http.StatusFound, "/home")
		Startprod()
	case devv == "true": //run dev
		fmt.Println("gonna dev...")
		c.Redirect(http.StatusFound, "/home")
		Startdev()
	case stopv == "true": //stop application
		fmt.Println("gonna stop...")
		c.Redirect(http.StatusFound, "/home")
		KillProcessByName("app.exe")
		KillProcessByName("app")
		Stopping(cmd)
	case reload == "true": //reload application
		fmt.Println("gonna reload...")
		c.Redirect(http.StatusFound, "/home")
		Reload()
	case routesconfig == "true": //create markup language for routes *not used any longer
		fmt.Println("gonna config...")
		c.Redirect(http.StatusFound, "/home")
		utfsg.Make("databaseconfig")
	case genroutev == "true": //generate routes
		fmt.Println("gonna gen routes...")
		c.Redirect(http.StatusFound, "/home")
		titles, urls = utfsg.GenRoutes()

		fmt.Println("before-", titles, urls)
	case dbv == "true": //generate database
		c.Redirect(http.StatusFound, "/home")
		db.Tempfile()
		db.CreateBucket("urls", "home", "/home")
	case showv == "true": //show routes
		titles, urls = db.GetAllkv("urls")
		libs, libtag = kdb.Getall("libs")
	case hotloadv == "true":
		c.Redirect(http.StatusFound, "/home")
		KillProcessByName("app.exe")
		KillProcessByName("app")
		Stopping(cmd)
		Hotreload()
	case deletev == "true": //delete routes
		titletrim := strings.ReplaceAll(title, " ", "")
		db.DeleteDB("urls", title)
		err := os.RemoveAll("app/templates/" + titletrim + ".html")
		if err != nil {
			log.Fatal(err)
		}
		c.Redirect(http.StatusFound, "/home")
	case libtagv == "true": //delete routes
		//get title
		titletrim := strings.ReplaceAll(title, " ", "")
		//get library
		lib := kdb.GetValue("libs", libtagsv)
		//add library to html file

		// paths, err := os.Getwd()
		// if err != nil {
		// 	log.Println(err)
		// }

		path := filepath.FromSlash(`app/templates/` + titletrim + `.html`)
		pp := strings.Replace(path, "\\", "/", -1)
		AddLibtoFile(pp, lib)
		c.Redirect(http.StatusFound, "/home")
	default:
		fmt.Println("none were used")
	}
	fmt.Println(libs)
	type Data struct {
		Titles []string
		Urls   []string
		Libs   []string
		Libtag []string
	}

	d := Data{Titles: titles, Urls: urls, Libs: libs, Libtag: libtag}

	return c.Render(http.StatusOK, "home.html", d)

}
