package datavar

import (
	"database/sql"
	"fmt"
)

type Data struct {
	ID    string "param:'id' query:'id' form:'id'"
	Datas string "param:'datas' query:'datas' form:'datas'"
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

		datas  string
		Datass []Data //used to store all users
	)

	//get from database
	rows, err := data.Query("select * from datavar")
	if err != nil {
		fmt.Println(err)
	}

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &datas)
		if err != nil {
			fmt.Println(err)
		}

		//store into memory
		u := Data{ID: id, Datas: datas}
		Datass = append(Datass, u)

	}
	//close everything
	defer rows.Close()
	defer data.Close()
	return Datass

}
