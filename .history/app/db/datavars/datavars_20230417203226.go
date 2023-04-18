package datavar

import (
	"database/sql"
	"fmt"
	"reflect"
)

type Data struct {
	ID    string
	F0    string
	F1    string
	F2    string
	F3    string
	F4    string
	F5    string
	F6    string
	F7    string
	F8    string
	F9    string
	Datas string
}

func Getvardata(data *sql.DB, datavar string) []Data {

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

		Datas []Data //used to store all users
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
		u := Data{ID: id, F0: f0, F1: f1, F2: f2, F3: f3, F4: f4, F5: f5, F6: f6, F7: f7, F8: f8, F9: f9}
		Datas = append(Datas, u)

	}

	//https://go.dev/play/p/0HrA-jrPZqG
	d := Data{}
	s := reflect.ValueOf(icb).Elem()
	//get table fields, types, and name
	for i := 0; i < s.NumField(); i++ {
		if s.Field(i).Interface() != "" {
			switch s.Type().Field(i).Name[0:1] {
			case "F": //fields
				d.Fields = append(d.Fields, s.Field(i).Interface().(string))
			case "T": //types
				d.Types = append(d.Types, s.Field(i).Interface().(string))
			case "S": //table name
				d.Table = s.Field(i).Interface().(string)
			}
		}
	}

	//store them in a slice
	for ii := 0; ii < len(d.Fields); ii++ {
		stringft := " " + d.Fields[ii] + " " + d.Types[ii] + " "
		d.Fieldtypes = append(d.Fieldtypes, stringft)
	}

	//save table and generate template in file
	WritetemplateData(DBcreates, file, d)

	//close everything
	defer rows.Close()
	defer data.Close()
	return Datas

}
