package adddatavar

import (
	"fmt"
	"net/http"

	. "github.com/golangast/groundup/internal/dbsql/datacreation/adddatavartopage"
	. "github.com/golangast/groundup/internal/generate/generators/gendatavars"

	"github.com/labstack/echo/v4"
)

func Adddatavar(c echo.Context) error {
	//get form data
	datavar := c.FormValue("datavar")
	urls := c.FormValue("urls")

	fmt.Print(datavar, urls)

	//add datavar to pagedb
	Adddatavartopage(datavar, urls)

	Gendatavars("../app", urls, datavar)

	//update database
	//redirect
	c.Redirect(http.StatusFound, "/home")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})
}