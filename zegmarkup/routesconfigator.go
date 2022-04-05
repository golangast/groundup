package zegmarkuputil

import (
	"fmt"

	"github.com/zendrulat123/groundup/zegmarkup/utfsg"
)

func Routesconfigator() {
	const files = "databaseconfig/dbpersis.fsg"
	utfsg.Make("databaseconfig")
	utfsg.Writetitle(files, "new")
	utfsg.UpdateText(files, "new", `new{
	 	path:"/new"
	 	}`)
	v := utfsg.GetValue(files, "new", "path")
	fmt.Println(v)
	//utfsg.DeleteLine(files, "path")
	utfsg.UpdateKey(files, "new", "path", "this")
}
