package createpage

import (
	"net/http"

	"github.com/labstack/echo/v4"
	db "gitlab.com/zendrulat123/groundup/db/get/urltitle"
)

func CreatePage(c echo.Context) error {
	url := c.FormValue("url")
	title := c.FormValue("title")
	t, u := db.GetUrlTitle(title, url)
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"titles": t,
		"urls":   u,
	})
}
