package datavar

import (
	"database/sql"
	"fmt"

	. "github.com/golangast/groundup/internal/dbsql/conn"
)

func Getvardata(data *sql.DB, datavar string) []Data {

	dbf := DBField{}

	//variables used to store data from the query
	var (
		id string
		f0 string
		f1 string
		f2 string
		f3 string
		f4 string
		f5 string
		f6 string
		f7 string
		f8 string
		f9 string

		Datas []DBField //used to store all users
	)

	//get from database
	rows, err := data.Query("select * from datavar")
	if err != nil {
		fmt.Println(err)
	}

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &f0, &f1, &f2, &f3, &f4, &f5, &f6, &f7, &f8, &f9)
		if err != nil {
			fmt.Println(err)
		}

		//store into memory
		dbf = DBField{ID: id, F0: f0, F1: f1, F2: f2, F3: f3, F4: f4, F5: f5, F6: f6, F7: f7, F8: f8, F9: f9}
		Datas = append(Datas, dbf)

	}

	//close everything
	defer rows.Close()
	defer data.Close()
	return Datas

}

type DBField struct {
	ID     string `param:"id" query:"id" header:"id" form:"id" json:"id" xml:"id"`
	Stable string `param:"stable" query:"stable" header:"stable" form:"stable" json:"stable" xml:"stable"`
	F0     string `param:"f0" query:"f0" header:"f0" form:"f0" json:"f0" xml:"f0"`
	F1     string `param:"f1" query:"f1" header:"f1" form:"f1" json:"f1" xml:"f1"`
	F2     string `param:"f2" query:"f2" header:"f2" form:"f2" json:"f2" xml:"f2"`
	F3     string `param:"f3" query:"f3" header:"f3" form:"f3" json:"f3" xml:"f3"`
	F4     string `param:"f4" query:"f4" header:"f4" form:"f4" json:"f4" xml:"f4"`
	F5     string `param:"f5" query:"f5" header:"f5" form:"f5" json:"f5" xml:"f5"`
	F6     string `param:"f6" query:"f6" header:"f6" form:"f6" json:"f6" xml:"f6"`
	F7     string `param:"f7" query:"f7" header:"f7" form:"f7" json:"f7" xml:"f7"`
	F8     string `param:"f8" query:"f8" header:"f8" form:"f8" json:"f8" xml:"f8"`
	F9     string `param:"f9" query:"f9" header:"f9" form:"f9" json:"f9" xml:"f9"`
	F10    string `param:"f10" query:"f10" header:"f10" form:"f10" json:"f10" xml:"f10"`
	F11    string `param:"f11" query:"f11" header:"f11" form:"f11" json:"f11" xml:"f11"`
}
