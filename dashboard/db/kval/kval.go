package kval

import (
	"fmt"
	"os"
	"strings"

	//https://github.com/kval-access-language/kval-language-specification
	kval "github.com/kval-access-language/kval-boltdb"
)

func Insertinnterbkeyvalue(rb string, ib string, k string, v string) {
	kb, err := kval.Connect("db/bolt.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening bolt database: %#v", err)
		os.Exit(1)
	}
	defer kval.Disconnect(kb)

	//Lets do a test insert...
	_, err = kval.Query(kb, "INS "+rb+" >> "+ib+" >>>> "+k+" :: "+v+"")
	if err != nil {
		//work with your error
		fmt.Print(err)
	}
}

func Insertkeyvalue(b string, k string, v string) {
	kb, err := kval.Connect("db/bolt.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening bolt database: %#v", err)
		os.Exit(1)
	}
	defer kval.Disconnect(kb)

	//Lets do a test insert...
	_, err = kval.Query(kb, "INS "+b+" >>>> "+k+" :: "+v+"")
	if err != nil {
		//work with your error
		fmt.Print(err)
	}
}
func Insertkey(b string, k string) {
	kb, err := kval.Connect("db/bolt.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening bolt database: %#v", err)
		os.Exit(1)
	}
	defer kval.Disconnect(kb)

	//Lets do a test insert...
	_, err = kval.Query(kb, "INS "+b+" >>>> "+k+"")
	if err != nil {
		//work with your error
		fmt.Print(err)
	}
}
func Createb(b string) {
	kb, err := kval.Connect("db/bolt.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening bolt database: %#v", err)
		os.Exit(1)
	}
	defer kval.Disconnect(kb)

	//Lets do a test insert...
	_, err = kval.Query(kb, "INS "+b+"")
	if err != nil {
		//work with your error
		fmt.Print(err)
	}
}

func Addbinb(ob string, nb string) {

	kb, err := kval.Connect("db/bolt.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening bolt database: %#v", err)
		os.Exit(1)
	}
	defer kval.Disconnect(kb)

	//Lets do a test insert...
	_, err = kval.Query(kb, "INS "+ob+" >> "+nb+"")
	if err != nil {
		//work with your error6
		fmt.Print(err)
	}
}

func Getall(b string) ([]string, []string) {
	kb, err := kval.Connect("db/bolt.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening bolt database: %#v", err)
		os.Exit(1)
	}
	defer kval.Disconnect(kb)

	//Lets do a test insert...GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key
	res, err := kval.Query(kb, "GET "+b+"")
	if err != nil {
		//work with your error
		fmt.Print(err)
	}
	var libs []string
	var tags []string

	for k, v := range res.Result {
		libs = append(libs, string(v))
		tags = append(tags, string(k))

	}

	return libs, tags
}
func GetValue(b, key string) string {
	kb, err := kval.Connect("db/bolt.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening bolt database: %#v", err)
		os.Exit(1)
	}
	defer kval.Disconnect(kb)

	//Lets do a test insert...GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key
	res, err := kval.Query(kb, "GET "+b+"")
	if err != nil {
		//work with your error
		fmt.Print(err)
	}
	keys := strings.ReplaceAll(key, "[", "")
	keyss := strings.ReplaceAll(keys, "]", "")

	fmt.Println(res.Result[keyss])
	fmt.Println("GET " + b + " >>>> " + key + " :: Value")
	return res.Result[keyss]
}
