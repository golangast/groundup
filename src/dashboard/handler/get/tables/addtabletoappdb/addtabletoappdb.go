package addtabletoappdb

import (
	"fmt"
	"net/http"
	"reflect"

	//. "github.com/golangast/groundup/internal/dbsql/appdata/addappdata"
	"github.com/golangast/groundup/internal/dbsql/conn"
	. "github.com/golangast/groundup/internal/dbsql/gettabledata"
	. "github.com/golangast/groundup/internal/generate/generators/gendatabase/dashcreatetable"

	"github.com/labstack/echo/v4"
)

func AddTableToAppDB(c echo.Context) error {

	table := c.Param("table")

	fmt.Print(table)

	dbfields := GetOnetabledata(table)
	fmt.Print(dbfields)

	tabledata := GetTableNameAppDatas(dbfields)
	fmt.Print(tabledata)

	//create app table to recieve
	Generatedatabasefields(&dbfields, table)

	//Dbinsert(tabledata.Table, tabledata.Fields, tabledata.Types)

	c.Redirect(http.StatusFound, "/home")
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"data": "",
	})
}
func GetTableNameAppDatas(icb conn.DBFields) Data {
	dd := Data{}
	fmt.Print(icb)
	for i := 0; i < reflect.ValueOf(&icb).Elem().NumField(); i++ {
		//check if filed is empty
		if reflect.ValueOf(&icb).Elem().Field(i).IsValid() {
			//check the first letter of the field
			switch reflect.ValueOf(&icb).Elem().Type().Field(i).Name[0:1] {
			case "F": //fields
				dd.Fields = append(dd.Fields, fmt.Sprint(reflect.ValueOf(&icb).Elem().Field(i).Interface()))
			case "T": //types
				dd.Types = append(dd.Types, fmt.Sprint(reflect.ValueOf(&icb).Elem().Field(i).Interface()))
			case "S": //table name
				dd.Table = fmt.Sprint(reflect.ValueOf(&icb).Elem().Field(i).Interface())
			}
		}
	}
	return dd

}

type Data struct {
	Table      string
	Fields     []string
	Types      []string
	Fieldtypes []string
}
