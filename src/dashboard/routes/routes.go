package routes

import (
	add "github.com/golangast/groundup/src/dashboard/handler/put/addlibtofooter"
	"github.com/golangast/groundup/src/dashboard/handler/put/createlib"

	adds "github.com/golangast/groundup/src/dashboard/handler/put/addlibtopage"
	"github.com/golangast/groundup/src/dashboard/handler/put/createcsslib"
	. "github.com/golangast/groundup/src/dashboard/handler/put/updateappdb"

	. "github.com/golangast/groundup/src/dashboard/handler/put/createdbdata"
	"github.com/golangast/groundup/src/dashboard/handler/put/createpage"

	home "github.com/golangast/groundup/src/dashboard/handler/home"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/", home.Home)
	e.GET("/:m", home.Home)
	e.GET("/:m/:titlev", home.Home)
	e.GET("/:m/:libtagsv/:titlev", home.Home)
	e.GET("/:m/remove/:table", home.Home)

	e.GET("/d/:m/:titlev", home.Home)

	// 	//create page
	e.POST("/page", createpage.CreatePage)
	e.POST("/lib", createlib.CreateLib)
	e.POST("/libcss", createcsslib.CreateCSSLib)
	e.POST("/chooselib", adds.Addlibtoppage)
	e.POST("/dbdata", Createdbdata)
	e.POST("/addfooterlib", add.Addlibtofooter)
	e.POST("/updateappdb", Updateappdb)

}
