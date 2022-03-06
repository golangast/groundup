package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
	. "gitlab.com/zendrulat123/groundup/frontend"
)

func Home(c echo.Context) error {

	bools := c.Param("bools")
	if bools == "true" {
		Serv()
	}

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"bool": bools,
	})

}
