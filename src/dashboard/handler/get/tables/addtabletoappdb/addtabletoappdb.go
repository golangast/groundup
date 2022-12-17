package addtabletoappdb

import (
	"fmt"
	"net/http"

	//. "github.com/golangast/groundup/internal/dbsql/appdata/addappdata"
	. "github.com/golangast/groundup/internal/dbsql/gettabledata"
	. "github.com/golangast/groundup/pkg/utility/generate"

	"github.com/labstack/echo/v4"
)

func AddTableToAppDB(c echo.Context) error {

	table := c.Param("table")

	fmt.Print(table)

	dbfields := GetOnetabledata(table)
	fmt.Print(dbfields)

	tabledata := GetTableNameAppData(dbfields)
	fmt.Print(tabledata)
	//add table to appdb

	// Generatestruct(*f, f.Stable)

	c.Redirect(http.StatusFound, "/home")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"data": "",
	})
}
