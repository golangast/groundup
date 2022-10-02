package getlib

import (
	"fmt"
	"log"

	. "github.com/golangast/groundup/src/dashboard/dbsql/conn"
)

func GetLib(Libtag string) string {

	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		lib string
	)
	i := 0 //used to get how many scans

	rows, err := data.Query("SELECT lib FROM library WHERE libtag = ?", Libtag)

	if err != nil {
		log.Fatal(err)
	}

	if rows.Next() {

		err := rows.Scan(&lib)

		if err != nil {
			log.Fatal(err)
		}
		i++
		fmt.Println("scan ", i)
		fmt.Printf("%v\n", lib)
	} else {

		fmt.Println("No lib found")
	}

	//close everything
	defer rows.Close()
	defer data.Close()
	return lib

}
func GetAllLib() []Library {

	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		id        string
		lib       string
		tag       string
		libraries []Library
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select * from library")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &lib, &tag)
		ErrorCheck(err)

		i++
		fmt.Println("scan ", i)
		L := Library{ID: id, Lib: lib, Tag: tag}
		libraries = append(libraries, L)

	}
	//close everything
	defer rows.Close()
	defer data.Close()
	return libraries

}

type Library struct {
	ID  string `param:"id" query:"id" form:"id"`
	Lib string `param:"lib" query:"lib" form:"lib"`
	Tag string `param:"tag" query:"tag" form:"tag"`
}
