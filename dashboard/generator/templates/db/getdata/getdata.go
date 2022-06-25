package getdata

var Getdatatemp = `
package getdata

import (
	"database/sql"
	"fmt"
)

type Data struct {
	ID     string "param:'id' query:'id' form:'id'"
	Datas    string "param:'datas' query:'datas' form:'datas'"
}
func Getalldata(data *sql.DB) []Data {


//variables used to store data from the query
var (
	id     string
	datas    string
	Datass   []Data //used to store all users
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
`

var Grabdatatemp = `db, err := DbConnection()
	if err != nil {
		fmt.Println(err)
	}
	var exists bool    
	type Data struct {
		ID     string "param:'id' query:'id' form:'id'"
		Datas    string "param:'datas' query:'datas' form:'datas'"
	}

	d:=Data{}
	const create string = "CREATE TABLE IF NOT EXISTS data ( id INTEGER NOT NULL PRIMARY KEY, datas DATETIME NOT NULL);"

	if _, err := db.Exec(create); err != nil {
		fmt.Println(err)
	   }

	   stmts := db.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM data WHERE datas=?)", d.Datas)
	err = stmts.Scan(&exists)
	if err != nil {
		fmt.Println(err)
	}

	//prepare the statement to ensure no sql injection
	stmt, err := db.Prepare("INSERT INTO data(datas) VALUES(?)")
	if err != nil {
		fmt.Println(err)
	}

	//actually make the execution of the query
	_, err = stmt.Exec(d.Datas)
	if err != nil {
		fmt.Println(err)
	}

	data:=Getalldata(db)
	` + "\n"
