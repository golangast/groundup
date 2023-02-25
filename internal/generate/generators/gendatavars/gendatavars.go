package gendatavars

import (
	"fmt"
	"log"
	"strings"

	. "github.com/golangast/groundup/internal/generate/templates/db/datavars"
	. "github.com/golangast/groundup/pkg/utility/general"
	. "github.com/golangast/groundup/pkg/utility/generate"
)

func Gendatavars(p, urls, datavar string) {

	//TODO generate in app: routes, handlers, db calls
	//generate route
	if !FindText("../app/app.go", `e.GET("/d/:`+urls+`", `+urls+`"`) {
		UpdateText("../app/app.go", "//#routes", `e.GET("/d/:`+urls+`", `+urls+`)`+"\n"+`//#routes`)
	}

	//make folder
	Makefolder(p)
	Makefolder(p + "/db/datavars")

	gdatafile := Makefile(p + "/db/datavars/datavars.go")
	d := make(map[string]string)

	d["t"] = datavar

	//write it to the file
	Writetemplate(Datavarstemp, gdatafile, d)
	nospaceroutes := strings.ReplaceAll(urls, " ", "")
	nospaceroutesnoslash := strings.ReplaceAll(nospaceroutes, "/", "")
	//make the handler
	if !FindText("../app/app.go", `func `+strings.ToUpper(urls)+`(c echo.Context)`) {
		UpdateText("../app/app.go", "//#handler", `func `+strings.ToUpper(datavar)+`(c echo.Context) error {

			//#getdatavar`+datavar+`

			return c.Render(http.StatusOK, "`+nospaceroutesnoslash+`.html", map[string]interface{}{
				"data":data,
			})
		
		}`+"\n"+`//#handler`)
	}

	//pull down dependencies
	PullDowndb("app")
	Pulldowneverythingbase("app")
	err, out, errout := Shellout(`cd .. && cd app && go mod tidy && go mod vendor`)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(out)
	fmt.Println("--- errs ---")
	fmt.Println(errout)
}
