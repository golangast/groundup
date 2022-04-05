package home

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/labstack/echo/v4"
	. "github.com/zendrulat123/groundup/cmd/ut"
	db "github.com/zendrulat123/groundup/dashboard/db"
	kdb "github.com/zendrulat123/groundup/dashboard/db/kval"
	"github.com/zendrulat123/groundup/zegmarkup/utfsg"

	. "github.com/zendrulat123/groundup/dashboard/configutil/createserver"
	. "github.com/zendrulat123/groundup/dashboard/configutil/serverutil"
)

func Home(c echo.Context) error {
	var titles []string
	var urls []string
	var libs []string

	const files = "databaseconfig/dbpersis.fsg"
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
		cmd = Startprod()
	case devv == "true": //run dev
		fmt.Println("gonna dev...")
		c.Redirect(http.StatusFound, "/home")
		cmd = Startdev()
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
		libs, _ = kdb.Getall("libs")
	case hotloadv == "true":
		c.Redirect(http.StatusFound, "/home")
		KillProcessByName("app.exe")
		KillProcessByName("app")
		Stopping(cmd)
		cmd = Hotreload()
	case deletev == "true": //delete routes
		titlev := c.Param("title")
		titletrim := strings.ReplaceAll(titlev, " ", "")
		db.DeleteDB("urls", titlev)
		err := os.RemoveAll("app/templates/" + titletrim + ".html")
		if err != nil {
			log.Fatal(err)
		}
		c.Redirect(http.StatusFound, "/home")

	default:
		fmt.Println("none were used")
	}
	fmt.Println(libs)
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"titles": titles,
		"urls":   urls,
		"libs":   libs,
	})

}
