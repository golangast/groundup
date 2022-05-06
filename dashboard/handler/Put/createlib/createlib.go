package createlib

import (
	"net/http"

	kdb "github.com/golangast/groundup/dashboard/db/kval"
	"github.com/labstack/echo/v4"
)

func CreateLib(c echo.Context) error {
	//get form data
	lib := c.FormValue("lib")
	libtag := c.FormValue("libtag")

	kdb.Insertkeyvalue("lib", libtag, lib)

	//redirect
	c.Redirect(http.StatusFound, "/home")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})
}
