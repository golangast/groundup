package getdata

import (
	"database/sql"
	"fmt"
)

type Data struct {
	ID    string "param:'id' query:'id' form:'id'"
	Datas string "param:'datas' query:'datas' form:'datas'"
}

func Getalldata(data *sql.DB) []Data {

	//variables used to store data from the query
	var (
		id    string
		datas string
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select * from data")
	if err != nil {
		fmt.Println(err)
	}

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &datas)
		if err != nil {
			fmt.Println(err)
		}

		i++
		fmt.Println("scan ", i)

		//store into memory
		u := Data{ID: id, Datas: datas}
		Datass = append(Datass, u)

	}
	//close everything
	defer rows.Close()
	defer data.Close()
	return Datass
}
