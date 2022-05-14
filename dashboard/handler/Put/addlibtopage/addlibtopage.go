package addLibFile

import (
	"net/http"
	"path/filepath"
	"strings"

	. "github.com/golangast/groundup/dashboard/dbsql/addlibtag"
	. "github.com/golangast/groundup/dashboard/dbsql/getlib"

	. "github.com/golangast/groundup/dashboard/handler/home/handlerutil"

	"github.com/labstack/echo/v4"
)

func AddLibFile(c echo.Context) error {
	//get form data
	titles := c.FormValue("titles")
	libtag := c.FormValue("libtag")
	//form the data
	titletrim := strings.ReplaceAll(titles, " ", "")
	path := filepath.FromSlash(`app/templates/` + titletrim + `.html`)
	pp := strings.Replace(path, "\\", "/", -1)
	//add lib to file
	AddLibtoFile(pp, libtag, titletrim)
	l := GetLib(libtag)
	//update database
	UpdateUrls(l, libtag, titles)
	//redirect
	c.Redirect(http.StatusFound, "/home")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})
}
