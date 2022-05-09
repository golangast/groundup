package routes

import (
	"github.com/golangast/groundup/dashboard/handler/Put/createpage"
	home "github.com/golangast/groundup/dashboard/handler/home"
	"github.com/golangast/groundup/dashboard/handler/put/createlib"
	"github.com/golangast/groundup/dashboard/handler/put/addlibtopage"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/", home.Home)
	e.GET("/:m", home.Home)
	e.GET("/:m/:titlev", home.Home)
	e.GET("/:m/:libtagsv/:titlev", home.Home)
	e.GET("/:m/:libtagsv/:footer", home.Home)

	//create page
	e.POST("/page", createpage.CreatePage)
	e.POST("/lib", createlib.CreateLib)
	e.POST("/chooselib", addLibFile.AddLibFile)

}
