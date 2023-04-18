package router

var Routertemp = `
func Routes(e *echo.Echo) {
	e.GET("/", Home)
	e.GET("/route/:routes", List)

	//#routes
}
`
