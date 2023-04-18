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
	//add import
	err := Gentextcomment("../app/app.go", `. "app/db/datavars"`, "//#import")
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

	//make route file
	routefile, err := Makefile("handler/get/" + route + "/" + route + ".go")
	if err != nil {
		fmt.Print(err, routefile)
	}

	//make folder
	Makefolder(p)
	Makefolder(p + "/db/datavars")

	gdatafile, err := Makefile(p + "/db/datavars/datavars.go")
	if err != nil {
		fmt.Print(err, gdatafile)
	}
	d := make(map[string]string)

	d["t"] = datavar

	//write it to the file
	Writetemplate(Datavarstemp, gdatafile, d)
	nospaceroutes := strings.ReplaceAll(urls, " ", "")
	nospaceroutesnoslash := strings.ReplaceAll(nospaceroutes, "/", "")
	//make the handler
	if !FindText("../app/app.go", `func `+strings.ToUpper(urls)) {
		UpdateText("../app/app.go", "//#handler", `func `+strings.ToUpper(urls)+`(c echo.Context) error {
			
			`+datavar+`:=Getvardata(`+`"`+datavar+`"`+`)
			//#getdatavar
			
			return c.Render(http.StatusOK, "`+nospaceroutesnoslash+`.html", map[string]interface{}{
				`+`"`+datavar+`"`+`:`+datavar+`,
				//#getdatavardata
			})
		
		}`+"\n"+`//#handler`)
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
