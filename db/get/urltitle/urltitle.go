package urltitle

import (
	"fmt"
	"strconv"

	ut "gitlab.com/zendrulat123/groundup/cmd/utserver"
	db "gitlab.com/zendrulat123/groundup/db/get"
)

var err error

func GetUrlTitle(title string, url string) ([]string, []string) {
	db.Conn()
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
	t, u := ut.GetUrlTitle(titlesUrl)
	fmt.Println(t, u)
	return t, u
}
