package datavars

var Datavarstemp = `
package datavars

import (
	"database/sql"
	"fmt"
)

https://go.dev/play/p/uYjz_77wX9w
type DBFields struct {
	ID     int   ` + "`" + `param:"id" query:"id" header:"id" form:"id" json:"id" xml:"id" ` + "`"
	Done string  ` + "`" +`param:"done" query:"done" header:"done" form:"done" json:"done" xml:"done"`
	Dtwo     string `param:"dtwo" query:"dtwo" header:"dtwo" form:"dtwo" json:"dtwo" xml:"dtwo"`
	Dthree     string `param:"dthree" query:"dthree" header:"dthree" form:"dthree" json:"dthree" xml:"dthree"`
	Dfour     string `param:"dfour" query:"dfour" header:"dfour" form:"dfour" json:"dfour" xml:"dfour"`
	Dfive     string `param:"dfive" query:"dfive" header:"dfive" form:"dfive" json:"dfive" xml:"dfive"`
	Dsix     string `param:"dsix" query:"dsix" header:"dsix" form:"dsix" json:"dsix" xml:"dsix"`
	Dseven     string `param:"dseven" query:"dseven" header:"dseven" form:"dseven" json:"dseven" xml:"dseven"`
	Deight     string `param:"deight" query:"deight" header:"deight" form:"deight" json:"deight" xml:"deight"`
	Dnine     string `param:"dnine" query:"dnine" header:"dnine" form:"dnine" json:"dnine" xml:"dnine"`
	Dten     string `param:"dten" query:"dten" header:"dten" form:"dten" json:"dten" xml:"dten"`
	Deleven     string `param:"deleven" query:"deleven" header:"deleven" form:"deleven" json:"deleven" xml:"deleven"`
}

func Getvardata() []Data {
//variables used to store data from the query
var (
	done   string
	dtwo    string
	dthree    string
	dfour    string
	dfive    string
	dsix    string
	dseven    string
	deight    string
	dnine    string
	dten    string
	deleven    string
	Datas  []Data //used to store all users
)

//get from database
rows, err := data.Query("select * from {{.t}}")
if err != nil {
	fmt.Println(err)
}

//cycle through the rows to collect all the data
for rows.Next() {
	err := rows.Scan(&id, &done, &dtwo, &dthree, &dfour, &dfive, &dsix, &dseven, &deight, &dnine, &dten, &deleven)
	if err != nil {
		fmt.Println(err)
	}
	//store into memory
	u := DBFields{ID: id, Done: done,  Dtwo: dtwo,  Dthree: dthree,  Dfour: dfour,  Dfive: dfive,  Dsix: dsix,  Dseven: dseven,  Deight: deight,  Dnine: dnine,  Dten: dten,  Deleven: deleven}
	Datas = append(Datas, u)

}
//close everything
rows.Close()
data.Close()
return Datas
}
`
