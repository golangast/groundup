package createdbdata

import (
	"fmt"
	"net/http"

	//. "github.com/golangast/groundup/services/dbsql/appdata/getappdata"
	. "github.com/golangast/groundup/services/dbsql/conn"

	//. "github.com/golangast/groundup/pkg/generateutility"

	. "github.com/golangast/groundup/services/dbsql/datacreation/savedbtables"
	. "github.com/golangast/groundup/services/generate/generators/gendatabase/dashcreatetable"

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
