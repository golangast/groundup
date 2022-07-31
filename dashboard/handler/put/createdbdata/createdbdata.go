package createdbdata

import (
	"fmt"
	"net/http"

	. "github.com/golangast/groundup/dashboard/dbsql/datacreation/savedbtables"
	. "github.com/golangast/groundup/dashboard/generator/gen/gendatabase/dashcreatetable"

	"github.com/labstack/echo/v4"
)

func Createdbdata(c echo.Context) error {

	f := new(DBFields)

	if err := c.Bind(f); err != nil {
		fmt.Println(err)
	}

	Savedbtables(f)
	Gendatasave(f)

	c.Redirect(http.StatusFound, "/home")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"data": f,
	})
}
