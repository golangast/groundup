package home

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
	. "gitlab.com/zendrulat123/groundup/cmd/ut"
	db "gitlab.com/zendrulat123/groundup/dashboard/db"
	"gitlab.com/zendrulat123/groundup/zegmarkup/utfsg"

	. "gitlab.com/zendrulat123/groundup/dashboard/configutil/createserver"
	. "gitlab.com/zendrulat123/groundup/dashboard/configutil/serverutil"
)

func Home(c echo.Context) error {
	var titles []string
	var urls []string
	const files = "databaseconfig/dbpersis.fsg"
	var cmd *exec.Cmd
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

	switch {
	case con == "true":
		fmt.Println("gonna config...")
		c.Redirect(http.StatusFound, "/home")
		CreateConfig()
	case s == "true":
		fmt.Println("gonna serv...")
		c.Redirect(http.StatusFound, "/home")
		Createservers()
	case prodv == "true":
		fmt.Println("gonna prod...")
		c.Redirect(http.StatusFound, "/home")
		cmd = Startprod()
	case devv == "true":
		fmt.Println("gonna dev...")
		c.Redirect(http.StatusFound, "/home")
		cmd = Startdev()
	case stopv == "true":
		fmt.Println("gonna stop...")
		c.Redirect(http.StatusFound, "/home")
		KillProcessByName("app.exe")
		KillProcessByName("app")
		Stopping(cmd)
	case reload == "true":
		fmt.Println("gonna reload...")
		c.Redirect(http.StatusFound, "/home")
		Reload()
	case routesconfig == "true":
		fmt.Println("gonna config...")
		c.Redirect(http.StatusFound, "/home")
		utfsg.Make("databaseconfig")
	case genroutev == "true":
		fmt.Println("gonna gen routes...")
		c.Redirect(http.StatusFound, "/home")
		titles, urls = utfsg.GenRoutes()
		fmt.Println("before-", titles, urls)
	case dbv == "true":
		c.Redirect(http.StatusFound, "/home")
		db.Tempfile()
		db.CreateBucket("urls", "home", "/home")
	case showv == "true":
		titles, urls = db.GetAllkv("urls")
	case hotloadv == "true":
		c.Redirect(http.StatusFound, "/home")
		KillProcessByName("app.exe")
		KillProcessByName("app")
		Stopping(cmd)
		cmd = Startprod()
	default:
		fmt.Println("none were used")
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"titles": titles,
		"urls":   urls,
	})

}
