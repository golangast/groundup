package savedbtables

import (
	"fmt"
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"

	. "github.com/golangast/groundup/internal/dbsql/conn"
)

type formdata struct {
	ftable  string
	ffields []string
	fvalues []string
}

func Dbinsert(table string, fields []string, values []string) {
	t := strings.Replace(table, "[", "", -1)
	tt := strings.Replace(t, "]", "", -1)
	// fmt.Println("table:", table, " fields:", fields, " values:", values)
	//build the query
	query := fmt.Sprintf("insert into %s(%s) values (%s)",
		tt,
		strings.Join(fields, ", "),
		strings.TrimSuffix(strings.Repeat("?,", len(fields)), ","),
	)
	spew.Dump(query)
	spew.Dump(strings.Join(values, `", `))
	s := make([]interface{}, len(t))
	for i, v := range values {
		s[i] = v
	}
	//TODO right here test if values are making it to database https://stackoverflow.com/questions/27689058/convert-string-to-interface
	data, err := AppDbConnection()
	ErrorCheck(err)
	_, err = data.Exec(query, s...)
	if err != nil {
		log.Fatal(err)
	}

}
