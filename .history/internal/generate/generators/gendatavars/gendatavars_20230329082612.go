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

	//writes import to app.go
	if !FindText("../app/app.go", `. "app/db/datavars"`) {
		UpdateText("../app/app.go", "//#import", `. "app/db/datavars"`+"\n"+`//#import`)
	}
	//generate route
	if !FindText("../app/app.go", `e.GET("/d/:`+urls+`", `+strings.ToUpper(urls)+`"`) {
		UpdateText("../app/app.go", "//#routes", `e.GET("/d/:`+urls+`", `+strings.ToUpper(urls)+`)`+"\n"+`//#routes`)
	}
	//#TODO make routes and handlers their own folders https://github.com/golangast/contribute/blob/main/cmd/gen.go
	//update router with route
	found, err = genutility.FindText("router/router.go", `contribute/handler/get/`+route)
	if err != nil {
		fmt.Print(err)
	}
	if !found {
		err := genutility.UpdateText("router/router.go", "//#import", `"contribute/handler/get/`+route+`"`+"\n"+`//#import`)
		if err != nil {
			fmt.Print(err)
		}
	}
	//make folder for handler
	err = genutility.Makefolder("handler/get/" + route)
	if err != nil {
		fmt.Print(err)
	}

	//make route file
	routefile, err := genutility.Makefile("handler/get/" + route + "/" + route + ".go")
	if err != nil {
		fmt.Print(err, routefile)
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
