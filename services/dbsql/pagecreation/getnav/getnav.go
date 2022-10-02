package getnav

import (
	"fmt"

	. "github.com/golangast/groundup/internal/dbsql/conn"
)

func Getnav() []Nav {

	data, err := DbConnection() //create db instance
	ErrorCheck(err)
	//variables used to store data from the query
	var (
		id   string
		a    string
		b    string
		c    string
		d    string
		e    string
		f    string
		g    string
		h    string
		i    string
		j    string
		navs []Nav
	)
	ii := 0 //used to get how many scans
	//get from database
	rows, err := data.Query("select * from nav")
	ErrorCheck(err)
	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &a, &b, &c, &d, &e, &f, &g, &h, &i, &j)
		ErrorCheck(err)

		ii++
		fmt.Println("scan ", i)

		//store into memory
		n := Nav{ID: id, A: a, B: b, C: c, D: d, E: e, F: f, G: g, H: h, I: i, J: j}
		navs = append(navs, n)
	}
	//close everything
	defer rows.Close()
	defer data.Close()
	return navs
}

type Nav struct {
	ID string `param:"id" query:"id" form:"id"`
	A  string `param:"a" query:"a" form:"a"`
	B  string `param:"b" query:"b" form:"b"`
	C  string `param:"c" query:"c" form:"c"`
	D  string `param:"d" query:"d" form:"d"`
	E  string `param:"e" query:"e" form:"e"`
	F  string `param:"f" query:"f" form:"f"`
	G  string `param:"g" query:"g" form:"g"`
	H  string `param:"h" query:"h" form:"h"`
	I  string `param:"i" query:"i" form:"i"`
	J  string `param:"j" query:"j" form:"j"`
}
