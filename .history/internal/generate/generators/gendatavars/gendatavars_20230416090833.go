package gendatavars

import (
	"fmt"
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
	Makefolder(p + "db/datavars")
	Makefolder(p + "routes")
	Makefolder(p + "handler/get/" + datavar)
	/*
		makes files
	*/
	routerfile, err := Makefile("../app/routes/router.go")
	ErrorCheck(err)
	handlerfile, err = Makefile("../app/handler/get/" + datavar + "/" + datavar + ".go")
	ErrorCheck(err)
	gdatafile, err := Makefile(p + "/db/datavars/datavars.go")
	ErrorCheck(err)
	urldatavar := make(map[string]string)
	urldatavar["datavar"] = datavar
	urldatavar["urls"] = strings.ToUpper(urls)
	/* write to files*/
	err = Writetemplate(Routertemp, routerfile, nil)
	ErrorCheck(err)
	err = Writetemplate(handlerfile, handlerfile, urldatavar)
	ErrorCheck(err)
	err = Writetemplate(Routertemp, routerfile, nil)
	ErrorCheck(err)
	/*
		write to files
	*/
	//add import router in app.go
	err = Gentextcomment("../app/app.go", `. "app/router"`, "//#import")
	ErrorCheck(err)
	//add routes in routes.go
	err = Gentextcomment("../app/router/router.go", `e.GET("/`+urls+`", `+strings.ToUpper(urls)+`"`, "//#routes")
	ErrorCheck(err)
	//add import handler in routes.go
	err = Gentextcomment("../app/router/router.go", `app/handler/get/`+datavar, "//#routes")
	ErrorCheck(err)

	d := make(map[string]string)
	d["t"] = datavar
	Writetemplate(Datavarstemp, gdatafile, d)
	nospaceroutes := strings.ReplaceAll(urls, " ", "")
	nospaceroutesnoslash := strings.ReplaceAll(nospaceroutes, "/", "")

	//add import handler in routes.go
	err = Gentextcomment("../app/app.go", `func `+strings.ToUpper(urls), "//#routes")
	ErrorCheck(err)

	err = Gentextcomment("../app/handler/get/"+datavar+"/"+datavar+".go", `func `+strings.ToUpper(urls)+`(c echo.Context) error {
			
		`+datavar+`:=Getvardata(`+`"`+datavar+`"`+`)
		//#getdatavar
		
		return c.Render(http.StatusOK, "`+nospaceroutesnoslash+`.html", map[string]interface{}{
			`+`"`+datavar+`"`+`:`+datavar+`,
			//#getdatavardata
		})
	
	}`, "")
	ErrorCheck(err)

	//pull down dependencies
	PullDowndb("app")
	Pulldowneverythingbase("app")
	err, out, errout := Shellout(`cd .. && cd app && go mod tidy && go mod vendor && go build`)
	ErrorCheck(err)

	fmt.Println(out)
	fmt.Println("--- errs ---")
	fmt.Println(errout)
}
func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
