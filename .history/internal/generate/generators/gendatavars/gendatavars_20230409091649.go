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
	/*
		makes folders
	*/
	Makefolder(p)
	Makefolder(p + "/db/datavars")
	Makefolder(p + "routes")
	Makefolder(p + "handler")
	/*
		makes files
	*/
	//router
	routerfile, err := Makefile("../app/router/router.go" + datavar + "/" + datavar + ".go")
	ErrorCheck(err)
	//handler
	handlerfile, err := Makefile("../app/handler/get/" + datavar + "/" + datavar + ".go")
	ErrorCheck(err)

	//datavars
	gdatafile, err := Makefile(p + "/db/datavars/datavars.go")
	ErrorCheck(err)

	/*
		write to files
	*/
	//add import
	err = Gentextcomment("../app/app.go", `. "app/router"`, "//#import")
	if err != nil {
		fmt.Print(err)
	}
	//add handler
	err = Gentextcomment("../app/app.go", `e.GET("/d/:`+urls+`", `+strings.ToUpper(urls)+`"`, "//#routes")
	if err != nil {
		fmt.Print(err)
	}
	//add import
	err = Gentextcomment("../app/router/router.go", `app/handler/get/`+datavar, "//#routes")
	if err != nil {
		fmt.Print(err)
	}

	d := make(map[string]string)
	d["t"] = datavar
	Writetemplate(Datavarstemp, gdatafile, d)
	nospaceroutes := strings.ReplaceAll(urls, " ", "")
	nospaceroutesnoslash := strings.ReplaceAll(nospaceroutes, "/", "")

	found, err := FindText("../app/app.go", `func `+strings.ToUpper(urls))
	if err != nil {
		fmt.Print(err)
	}
	if !found {
		err := UpdateText("../app/app.go", "//#handler", `func `+strings.ToUpper(urls)+`(c echo.Context) error {
			
			`+datavar+`:=Getvardata(`+`"`+datavar+`"`+`)
			//#getdatavar
			
			return c.Render(http.StatusOK, "`+nospaceroutesnoslash+`.html", map[string]interface{}{
				`+`"`+datavar+`"`+`:`+datavar+`,
				//#getdatavardata
			})
		
		}`+"\n"+`//#handler`)
		if err != nil {
			fmt.Print(err)
		}
	}

	//pull down dependencies
	PullDowndb("app")
	Pulldowneverythingbase("app")
	err, out, errout := Shellout(`cd .. && cd app && go mod tidy && go mod vendor && go build`)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(out)
	fmt.Println("--- errs ---")
	fmt.Println(errout)
}
func ErrorCheck(err error, s string) {
	if err != nil {
		fmt.Println(err.Error(), s)
	}
}
