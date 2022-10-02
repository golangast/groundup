package createcsslib

import (
	"net/http"

	. "github.com/golangast/groundup/services/dbsql/pagecreation/addlibtag"
	"github.com/labstack/echo/v4"
)

func CreateCSSLib(c echo.Context) error {
	//get form data
	lib := c.FormValue("lib")
	libtag := c.FormValue("libtag")
	AddCSSlib(lib, libtag)

	//redirect
	c.Redirect(http.StatusFound, "/home")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})
}
