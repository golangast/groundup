package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	. "gitlab.com/zendrulat123/groundup/cmd/utserver"
	"gitlab.com/zendrulat123/groundup/frontend/handler/Put/createpage"
	home "gitlab.com/zendrulat123/groundup/frontend/handler/home"
)

var hh = "fff"

func Routes(e *echo.Echo) {
	e.GET("/home", home.Home)
	e.GET("/config/:config", home.Home)
	e.GET("/server/:server", home.Home)
	e.GET("/startprod/:prodv", home.Home)
	e.GET("/startdev/:devv", home.Home)
	e.GET("/stop/:stopv", home.Home)
	e.GET("/reload/:reloadv", home.Home)

	e.GET("/"+hh, func(c echo.Context) error {
		Createtemplatefile(hh + ".html")
		return c.Render(http.StatusOK, hh+".html", map[string]interface{}{})
	})
	//create page
	e.POST("/page", createpage.CreatePage)
}
