package routes

import (
	"github.com/labstack/echo/v4"

	home "github.com/golangast/groundup/src/dashboard/handler/get/home"
	add "github.com/golangast/groundup/src/dashboard/handler/put/addlibtofooter"
	adds "github.com/golangast/groundup/src/dashboard/handler/put/addlibtopage"

	. "github.com/golangast/groundup/src/dashboard/handler/get/tables/addtabletoappdb"
	. "github.com/golangast/groundup/src/dashboard/handler/put/adddatavar"
	"github.com/golangast/groundup/src/dashboard/handler/put/createcsslib"
	. "github.com/golangast/groundup/src/dashboard/handler/put/createdbdata"
	"github.com/golangast/groundup/src/dashboard/handler/put/createlib"
	"github.com/golangast/groundup/src/dashboard/handler/put/createpage"
	. "github.com/golangast/groundup/src/dashboard/handler/put/updateappdb"
)

func Routes(e *echo.Echo) {
	e.GET("/", home.Home)
	e.GET("/:m", home.Home)
	e.GET("/:m/:titlev", home.Home)
	e.GET("/:m/:libtagsv/:titlev", home.Home)
	e.GET("/:m/remove/:table", home.Home)
	e.GET("/d/:m/:titlev", home.Home)
	e.GET("/addtabletoapp/:table", AddTableToAppDB) //table editing has to be get for params

	//create page
	e.POST("/page", createpage.CreatePage)
	e.POST("/lib", createlib.CreateLib)
	e.POST("/libcss", createcsslib.CreateCSSLib)
	e.POST("/chooselib", adds.Addlibtoppage)
	e.POST("/dbdata", Createdbdata)
	e.POST("/addfooterlib", add.Addlibtofooter)
	e.POST("/updateappdb/:table", Updateappdb)
	e.POST("/adddatavar", Adddatavar)

}
