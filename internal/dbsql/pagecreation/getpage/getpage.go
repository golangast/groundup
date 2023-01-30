package getpage

import (
	"log"

	. "github.com/golangast/groundup/internal/dbsql/conn"
)

func GetPage() ([]string, []string) {

	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		url    string
		title  string
		Urls   []string
		Titles []string
	)
	//i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select * from urls")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&url, &title)
		ErrorCheck(err)

		// i++
		// fmt.Println("scan ", i)

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
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		filename string
	)
	//i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select filename from urls where titles ==" + title + ";")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&filename)
		ErrorCheck(err)
		// i++
		// fmt.Println("scan ", i)

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
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		titles string
	)
	// i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("SELECT titles FROM urls WHERE urls = ?", urls)
	if err != nil {
		log.Fatal(err)
	}

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&titles)
		ErrorCheck(err)
		// i++
		// fmt.Println("scan ", i)

		//store into memory
		defer rows.Close()
		defer data.Close()
		return titles

	}
	//close everything
	return titles
}
