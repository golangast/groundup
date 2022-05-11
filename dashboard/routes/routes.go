package routes

import (
	"github.com/golangast/groundup/dashboard/handler/put/createcsslib"
	"github.com/golangast/groundup/dashboard/handler/put/createpage"

	home "github.com/golangast/groundup/dashboard/handler/home"
	"github.com/golangast/groundup/dashboard/handler/put/addlibtofooter"
	addLibFile "github.com/golangast/groundup/dashboard/handler/put/addlibtopage"
	"github.com/golangast/groundup/dashboard/handler/put/createlib"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/", home.Home)
	e.GET("/:m", home.Home)
	e.GET("/:m/:titlev", home.Home)
	e.GET("/:m/:libtagsv/:titlev", home.Home)
	e.GET("/d/:m/:titlev", home.Home)

	//create page
	e.POST("/page", createpage.CreatePage)
	e.POST("/lib", createlib.CreateLib)
	e.POST("/libcss", createcsslib.CreateCSSLib)
	e.POST("/chooselib", addLibFile.AddLibFile)

	e.POST("/addfooterlib", addlibtofooter.Addlibtofooter)

}
