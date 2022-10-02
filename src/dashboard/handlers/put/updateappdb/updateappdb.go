package Updateappdb

import (
	"fmt"
	"net/http"
	"strings"

	//. "github.com/golangast/groundup/services/dbsql/conn"

	"github.com/labstack/echo/v4"
)

type Anys struct {
}

func Updateappdb(c echo.Context) error {
	var aa any

	if err := c.Bind(aa); err != nil {
		fmt.Println(err)
	}
	var fields []string
	var values []string

	for f, v := range c.Request().Form {
		justString := strings.Join(v, "")
		fields = append(fields, f)
		values = append(values, justString)
		fmt.Print(fields, values)
	}

	//add data to apps tables

	c.Redirect(http.StatusFound, "/home")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})

}
