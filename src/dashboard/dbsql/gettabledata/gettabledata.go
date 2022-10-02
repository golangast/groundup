package gettabledata

import (
	"fmt"

	. "github.com/golangast/groundup/src/dashboard/dbsql/conn"
)

func Gettabledata() []DBFields {
	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		id        int
		stable    string
		f0        string
		f1        string
		f2        string
		f3        string
		f4        string
		f5        string
		f6        string
		f7        string
		f8        string
		f9        string
		f10       string
		f11       string
		t0        string
		t1        string
		t2        string
		t3        string
		t4        string
		t5        string
		t6        string
		t7        string
		t8        string
		t9        string
		t10       string
		t11       string
		DBFieldss []DBFields //used to store all users
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select * from savedtable")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &stable, &f0, &f1, &f2, &f3, &f4, &f5, &f6, &f7, &f8, &f9, &f10, &f11, &t0, &t1, &t2, &t3, &t4, &t5, &t6, &t7, &t8, &t9, &t10, &t11)
		ErrorCheck(err)

		i++
		fmt.Println("scan ", i)

		//store into memory
		u := DBFields{ID: id, Stable: stable, F0: f0, F1: f1, F2: f2, F3: f3, F4: f4, F5: f5, F6: f6, F7: f7, F8: f8, F9: f9, F10: f10, F11: f11, T0: t0, T1: t1, T2: t2, T3: t3, T4: t4, T5: t5, T6: t6, T7: t7, T8: t8, T9: t9, T10: t10, T11: t11}
		DBFieldss = append(DBFieldss, u)

	}
	//close everything
	defer rows.Close()
	defer data.Close()

	//Createfieldtable(DBFieldss)

	return DBFieldss

}

//to get one table
func GetOnetabledata(p string) DBFields {
	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		id     int
		stable string
		f0     string
		f1     string
		f2     string
		f3     string
		f4     string
		f5     string
		f6     string
		f7     string
		f8     string
		f9     string
		f10    string
		f11    string
		t0     string
		t1     string
		t2     string
		t3     string
		t4     string
		t5     string
		t6     string
		t7     string
		t8     string
		t9     string
		t10    string
		t11    string
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("SELECT * FROM savedtable WHERE stable = ?", p)
	ErrorCheck(err)
	var u DBFields
	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &stable, &f0, &f1, &f2, &f3, &f4, &f5, &f6, &f7, &f8, &f9, &f10, &f11, &t0, &t1, &t2, &t3, &t4, &t5, &t6, &t7, &t8, &t9, &t10, &t11)
		ErrorCheck(err)

		i++
		fmt.Println("scan ", i)

		//store into memory
		u = DBFields{ID: id, Stable: stable, F0: f0, F1: f1, F2: f2, F3: f3, F4: f4, F5: f5, F6: f6, F7: f7, F8: f8, F9: f9, F10: f10, F11: f11, T0: t0, T1: t1, T2: t2, T3: t3, T4: t4, T5: t5, T6: t6, T7: t7, T8: t8, T9: t9, T10: t10, T11: t11}

	}
	//close everything
	defer rows.Close()
	defer data.Close()
	return u

}
