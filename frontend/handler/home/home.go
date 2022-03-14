package home

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
	. "gitlab.com/zendrulat123/groundup/cmd/ut"

	. "gitlab.com/zendrulat123/groundup/configutil/createserver"
	. "gitlab.com/zendrulat123/groundup/configutil/serverutil"
)

func Home(c echo.Context) error {
	var cmd *exec.Cmd
	con := c.Param("config")
	if con == "true" {
		CreateConfig()
	}
	s := c.Param("server")
	if s == "true" {
		Createservers()
	}
	prodv := c.Param("prodv")
	if prodv == "true" {
		cmd = Startprod()
	}
	devv := c.Param("devv")
	if devv == "true" {
		cmd = Startdev()
	}
	stopv := c.Param("stopv")
	if stopv == "true" {
		fmt.Println("gonna stop...")
		Stopping(cmd)
	}
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})

}
