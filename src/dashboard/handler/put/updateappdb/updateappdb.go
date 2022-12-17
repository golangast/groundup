package Updateappdb

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
	. "github.com/golangast/groundup/internal/dbsql/appdata/addappdata"
	"github.com/labstack/echo/v4"
)

type Anys struct {
}

func Updateappdb(c echo.Context) error {
	var aa any
	table := c.Param("table")

	if err := c.Bind(aa); err != nil {
		fmt.Println(err)
	}
	var fields []string
	var values []string
	spew.Dump(c.Request().Form)
	for f, v := range c.Request().Form {
		if f != "id" {
			fields = append(fields, f)
			justString := strings.Join(v, "")
			values = append(values, justString)
		}

		fmt.Print("fv", fields, values)
	}
	fmt.Println("news", table, fields[1:], values[1:])
	//add data to apps tables
	Dbinsert(table, fields, values)

	c.Redirect(http.StatusFound, "/home")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})

}
