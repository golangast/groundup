package addlibtofooter

import (
	"net/http"

	. "github.com/golangast/groundup/internal/dbsql/pagecreation/getlib"
	. "github.com/golangast/groundup/pkg/utility/handler"

	"github.com/labstack/echo/v4"
)

func Addlibtofooter(c echo.Context) error {
	//get form data
	libtag := c.FormValue("libtag")
	//get library
	l := GetLib(libtag)
	//add it to the footer
	AddLibtoFilebyTitle(l, "footer")
	//redirect
	c.Redirect(http.StatusFound, "/home")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})
}
