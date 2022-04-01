package createpage

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	db "gitlab.com/zendrulat123/groundup/dashboard/db"
)

func CreatePage(c echo.Context) error {
	url := c.FormValue("url")
	title := c.FormValue("title")
	db.PutDB("urls", title, url)
	urltrim := strings.ReplaceAll(url, " ", "")
	urltrimslash := strings.ReplaceAll(urltrim, "/", "")
	if _, err := os.Stat("app/templates/" + urltrimslash + ".html"); errors.Is(err, os.ErrNotExist) {
		mfile, err := os.Create("app/templates/" + urltrimslash + ".html")
		if err != nil {
			fmt.Println("error -", err, mfile)
		}
	}
	c.Redirect(http.StatusFound, "/home")

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})
}
