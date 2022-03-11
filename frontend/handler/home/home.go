package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
	. "gitlab.com/zendrulat123/groundup/cmd/ut"
	
	. "gitlab.com/zendrulat123/groundup/configutil/createserver"
)

func Home(c echo.Context) error {

	con := c.Param("config")
	if con == "true" {
		CreateConfig()
	}
	s := c.Param("server")
	if s == "true" {
		Createservers()
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})

}
