package createlib

import (
	"net/http"

	"github.com/labstack/echo/v4"
	kdb "github.com/zendrulat123/groundup/dashboard/db/kval"
)

var counter int

func CreateLib(c echo.Context) error {
	//get form data
	lib := c.FormValue("lib")
	//create bucket
	counter++
	kdb.Insertkeyvalue("lib", string(counter), lib)

	//redirect
	c.Redirect(http.StatusFound, "/home")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})
}
