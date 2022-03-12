package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/zendrulat123/groundup/frontend/handler/Put/createpage"
	home "gitlab.com/zendrulat123/groundup/frontend/handler/home"
)

var hh = "news"

func Routes(e *echo.Echo) {
	e.GET("/home", home.Home)
	e.GET("/config/:config", home.Home)
	e.GET("/server/:server", home.Home)
	e.GET(hh, func(c echo.Context) error {
		return c.Render(http.StatusOK, hh, map[string]interface{}{})
	})
	//create page
	e.POST("/page", createpage.CreatePage)

}
