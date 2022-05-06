package routes

import (
	"github.com/golangast/groundup/dashboard/handler/Put/createpage"
	home "github.com/golangast/groundup/dashboard/handler/home"
	"github.com/golangast/groundup/dashboard/handler/put/createlib"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/home", home.Home)
	e.GET("/config/:config", home.Home)
	e.GET("/server/:server", home.Home)
	e.GET("/startprod/:prodv", home.Home)
	e.GET("/startdev/:devv", home.Home)
	// e.GET("/stop/:stopv", home.Home)
	e.GET("/reload/:reloadv", home.Home)
	e.GET("/routesconfig/:routesconfigv", home.Home)
	e.GET("/genroute/:genroutev", home.Home)
	e.GET("/db/:dbv", home.Home)
	e.GET("/show/:showv", home.Home)
	e.GET("/showlibs/:showlibsv", home.Home)
	e.GET("/hotload/:hotloadv", home.Home)
	e.GET("/delete/:deletev/:titlev", home.Home)
	e.GET("/stopapp/:stopv", home.Home)
	e.GET("/stopwatcher/:killwv", home.Home)
	e.GET("/observe/:obsv", home.Home)

	e.GET("/addlib/:libtagv/:libtagsv/:titlev", home.Home)

	//create page
	e.POST("/page", createpage.CreatePage)
	e.POST("/lib", createlib.CreateLib)

}
