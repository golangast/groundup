package routes

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/zendrulat123/groundup/dashboard/handler/Put/createpage"
	home "gitlab.com/zendrulat123/groundup/dashboard/handler/home"
)

func Routes(e *echo.Echo) {
	e.GET("/home", home.Home)
	e.GET("/config/:config", home.Home)
	e.GET("/server/:server", home.Home)
	e.GET("/startprod/:prodv", home.Home)
	e.GET("/startdev/:devv", home.Home)
	e.GET("/stop/:stopv", home.Home)
	e.GET("/reload/:reloadv", home.Home)
	e.GET("/routesconfig/:routesconfigv", home.Home)
	e.GET("/genroute/:genroutev", home.Home)
	e.GET("/db/:dbv", home.Home)
	e.GET("/show/:showv", home.Home)
	e.GET("/hotload/:hotloadv", home.Home)
	e.GET("/delete/:deletev/:title", home.Home)

	//create page
	e.POST("/page", createpage.CreatePage)
}
