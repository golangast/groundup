package createdbdata

import (
	"fmt"
	"net/http"

	. "github.com/golangast/groundup/internal/dbsql/conn"
	. "github.com/golangast/groundup/pkg/utility/handler"

	. "github.com/golangast/groundup/internal/dbsql/datacreation/savedbtables"
	. "github.com/golangast/groundup/internal/generate/generators/gendatabase/dashcreatetable"

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

	Reload()
	Startprod()

	//Generatestruct(*f, f.Stable)
	c.Redirect(http.StatusFound, "/home")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"data": f,
	})
}
