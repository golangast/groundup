package handler 

var Handlertemp=`
func `+strings.ToUpper(urls)+`(c echo.Context) error {
			
    `+datavar+`:=Getvardata(`+`"`+datavar+`"`+`)
    //#getdatavar
    
    return c.Render(http.StatusOK, "`+nospaceroutesnoslash+`.html", map[string]interface{}{
        `+`"`+datavar+`"`+`:`+datavar+`,
        //#getdatavardata
    })

}
`