package geturls

import (
	. "github.com/golangast/groundup/internal/dbsql/conn"
)

func GetUrls() []Urls {

	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		id       string
		urls     string
		titles   string
		lib      string
		libtag   string
		css      string
		csstag   string
		filename string
		datavars string
		urlss    []Urls //used to store all users
	)

	//get from database
	rows, err := data.Query("select * from urls")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &urls, &titles, &lib, &libtag, &css, &csstag, &filename, &datavars)
		ErrorCheck(err)

		//store into memory
		u := Urls{ID: id, Urls: urls, Titles: titles, Lib: lib, Libtag: libtag, Css: css, Csstag: csstag, Filename: filename, Datavars: datavars}
		urlss = append(urlss, u)

	}
	//close everything
	rows.Close()
	data.Close()
	return urlss

}

type Urls struct {
	ID       string `param:"id" query:"id,omitempty" form:"id"`
	Urls     string `param:"urls" query:"urls,omitempty" form:"urls,omitempty"`
	Titles   string `param:"titles" query:"titles,omitempty" form:"titles,omitempty"`
	Lib      string `param:"lib" query:"lib,omitempty" form:"lib,omitempty"`
	Libtag   string `param:"libtag" query:"libtag,omitempty" form:"libtag,omitempty"`
	Css      string `param:"css" query:"css,omitempty" form:"css,omitempty"`
	Csstag   string `param:"csstag" query:"csstag,omitempty" form:"csstag,omitempty"`
	Filename string `param:"filename" query:"filename,omitempty" form:"filename,omitempty"`
	Datavars string `param:"datavars" query:"datavars,omitempty" form:"datavars,omitempty"`
}
