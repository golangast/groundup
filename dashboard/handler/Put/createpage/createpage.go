package createpage

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/labstack/echo/v4"
	db "github.com/zendrulat123/groundup/dashboard/db"
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

		var Bodytemp = `
	{{ .header }}
	<!-- ### -->
	{{ .footer }}
	`
		tms := template.Must(template.New("queue").Parse(Bodytemp))
		//all of this is needed to parse {{define header}} and {{end}}
		m := make(map[string]string)
		header := fmt.Sprintf(`{{template "header"}}%s`, "")
		footer := fmt.Sprintf(`{{template "footer"}}%s`, "")
		m["header"] = header
		m["footer"] = footer
		err = tms.Execute(mfile, m)
		if err != nil {
			log.Print("execute: ", err)
		}
	}

	c.Redirect(http.StatusFound, "/home")

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})
}