package createdbdata

import (
	"fmt"
	"net/http"

	//. "github.com/golangast/groundup/dashboard/dbsql/appdata/getappdata"
	. "github.com/golangast/groundup/dashboard/dbsql/conn"

	//. "github.com/golangast/groundup/dashboard/generator/generatorutility"

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
	Generatedatabasefields(f, f.Stable)

	//Getappdata(f.Stable)

	//Generatestruct(*f, f.Stable)
	c.Redirect(http.StatusFound, "/home")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"data": f,
	})
}
