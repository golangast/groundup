package handler 

var Handlertemp=`
func {{.urls}} (c echo.Context) error {
			
    `+{{.datavar}}+`:=Getvardata(`+`"`+{{.datavar}}+`"`+`)
    //#getdatavar
    
    return c.Render(http.StatusOK, {{.urls}}.html", map[string]interface{}{
        `+`"`+{{.datavar}}+`"`+`:{{.datavar}},
        //#getdatavardata
    })

}
`