package createpage

import (
	"fmt"
	"net/http"
	"strconv"


	"github.com/labstack/echo/v4"
	"github.com/syndtr/goleveldb/leveldb" 
	ut "gitlab.com/zendrulat123/groundup/cmd/utserver"
	
)

var err error

type Pages struct {
	Title []string `json:"title,omitempty"`
	Url   []string `json:"url,omitempty"`
}

func CreatePage(c echo.Context) error {
	// p := new(Pages)
	// if err = c.Bind(p); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }
	url := c.FormValue("url")
	title := c.FormValue("title")
	db, err := leveldb.OpenFile("path/to/db", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	var counter int
	counter++
	s := strconv.Itoa(counter)

	err = db.Put([]byte(s), []byte(title+":"+url), nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	var titlesUrl []string

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		titlesUrl = append(titlesUrl, string(value))
		fmt.Println("key", string(key), string(value))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(titlesUrl)

	for _, titleurl := range titlesUrl {
		ut.GetUrlTitle(titleurl)
	}
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})
}
