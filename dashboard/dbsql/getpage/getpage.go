package getpage

import (
	"fmt"
	"log"

	. "github.com/golangast/groundup/dashboard/dbsql/conn"
)

func GetPage() ([]string, []string) {

	data, err := DbConnection() //create db instance
	ErrorCheckss(err)

	//variables used to store data from the query
	var (
		url    string
		title  string
		Urls   []string
		Titles []string
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select * from urls")
	ErrorCheckss(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&url, &title)
		ErrorCheckss(err)

		i++
		fmt.Println("scan ", i)

		//store into memory
		Urls = append(Urls, url)
		Titles = append(Titles, url)

	}
	//close everything
	defer rows.Close()
	defer data.Close()
	return Urls, Titles

}

func GetPageFile(title string) string {

	data, err := DbConnection() //create db instance
	ErrorCheckss(err)

	//variables used to store data from the query
	var (
		filename string
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select filename from urls where titles ==" + title + ";")
	ErrorCheckss(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&filename)
		ErrorCheckss(err)
		i++
		fmt.Println("scan ", i)

		//store into memory
		defer rows.Close()
		defer data.Close()
		return filename

	}
	//close everything
	return filename
}
func GetPagetitle(urls string) string {

	data, err := DbConnection() //create db instance
	ErrorCheckss(err)

	//variables used to store data from the query
	var (
		titles string
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("SELECT titles FROM urls WHERE urls = ?", urls)
	if err != nil {
		log.Fatal(err)
	}

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&titles)
		ErrorCheckss(err)
		i++
		fmt.Println("scan ", i)

		//store into memory
		defer rows.Close()
		defer data.Close()
		return titles

	}
	//close everything
	return titles
}

func ErrorCheckss(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
