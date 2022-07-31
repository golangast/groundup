package gettabledata

import (
	"fmt"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
)

func Gettabledata() []DBFields {
	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		id        int
		stable    string
		f0        string
		f1        string
		f2        string
		f3        string
		f4        string
		f5        string
		f6        string
		f7        string
		f8        string
		f9        string
		f10       string
		f11       string
		t0        string
		t1        string
		t2        string
		t3        string
		t4        string
		t5        string
		t6        string
		t7        string
		t8        string
		t9        string
		t10       string
		t11       string
		DBFieldss []DBFields //used to store all users
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select * from savedtable")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &stable, &f0, &f1, &f2, &f3, &f4, &f5, &f6, &f7, &f8, &f9, &f10, &f11, &t0, &t1, &t2, &t3, &t4, &t5, &t6, &t7, &t8, &t9, &t10, &t11)
		ErrorCheck(err)

		i++
		fmt.Println("scan ", i)

		//store into memory
		u := DBFields{ID: id, Stable: stable, F0: f0, F1: f1, F2: f2, F3: f3, F4: f4, F5: f5, F6: f6, F7: f7, F8: f8, F9: f9, F10: f10, F11: f11, T0: t0, T1: t1, T2: t2, T3: t3, T4: t4, T5: t5, T6: t6, T7: t7, T8: t8, T9: t9, T10: t10, T11: t11}
		DBFieldss = append(DBFieldss, u)

	}
	//close everything
	defer rows.Close()
	defer data.Close()
	return DBFieldss

}

type DBFields struct {
	ID     int    `param:"id" query:"id" header:"id" form:"id" json:"id" xml:"id"`
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

	T0  string `param:"t0" query:"t0" header:"t0" form:"t0" json:"t0" xml:"t0"`
	T1  string `param:"t1" query:"t1" header:"t1" form:"t1" json:"t1" xml:"t1"`
	T2  string `param:"t2" query:"t2" header:"t2" form:"t2" json:"t2" xml:"t2"`
	T3  string `param:"t3" query:"t3" header:"t3" form:"t3" json:"t3" xml:"t3"`
	T4  string `param:"t4" query:"t4" header:"t4" form:"t4" json:"t4" xml:"t4"`
	T5  string `param:"t5" query:"t5" header:"t5" form:"t5" json:"t5" xml:"t5"`
	T6  string `param:"t6" query:"t6" header:"t6" form:"t6" json:"t6" xml:"t6"`
	T7  string `param:"t7" query:"t7" header:"t7" form:"t7" json:"t7" xml:"t7"`
	T8  string `param:"t8" query:"t8" header:"t8" form:"t8" json:"t8" xml:"t8"`
	T9  string `param:"t9" query:"t9" header:"t9" form:"t9" json:"t9" xml:"t9"`
	T10 string `param:"t10" query:"t10" header:"t10" form:"t10" json:"t10" xml:"t10"`
	T11 string `param:"t11" query:"t11" header:"t11" form:"t11" json:"t11" xml:"t11"`
}
