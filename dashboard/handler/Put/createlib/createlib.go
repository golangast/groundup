package createlib

import (
	"net/http"

	. "github.com/golangast/groundup/dashboard/dbsql/addlibtag"
	"github.com/labstack/echo/v4"
)

func CreateLib(c echo.Context) error {
	//get form data
	lib := c.FormValue("lib")
	libtag := c.FormValue("libtag")
	u := Urls{Lib: lib, Libtag: libtag}
	Addlib(u)

	//redirect
	c.Redirect(http.StatusFound, "/home")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})
}
