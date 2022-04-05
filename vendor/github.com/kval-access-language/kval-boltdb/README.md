# KVAL-BoltDB

[![Build Status](https://travis-ci.org/kval-access-language/kval-boltdb.svg?branch=master)](https://travis-ci.org/kval-access-language/kval-boltdb)
[![GoDoc](https://godoc.org/github.com/kval-access-language/kval-boltdb?status.svg)](https://godoc.org/github.com/kval-access-language/kval-boltdb)
[![Go Report Card](https://goreportcard.com/badge/github.com/kval-access-language/kval-boltdb)](https://goreportcard.com/report/github.com/kval-access-language/kval-boltdb)

BoltDB bindings for [KVAL](https://github.com/kval-access-language/kval)

Package kvalbolt implements a BoltDB binding for KVAL (Key Value Access Language)
The binding provides more than just 'boilerplate' for working with BoltDB. It
implements a language thus enabling access from a higher point of abstraction.
Users of the binding can provide simple instructions and the binding will take
care of Bucket creation (and deletion), and the generation of keys and values,
plus their maintenance, no matter how deep into the database structure they
are needed, or indeed exist.

"Everything should be made as simple as possible, but no simpler." - Albert Einstein

The most up-to-date specification for KVAL can be found here:
https://github.com/kval-access-language/kval

My first blog post describing my thoughts a little better can be found here: 
http://exponentialdecay.co.uk/blog/key-value-access-language-kval-for-boltdb-and-golang/

### Key Value Access Language

I have created a modest specification for a key value access langauge. 
It allows for input and access of values to a key value store such as Golang's
[BoltDB](https://github.com/boltdb/). 

The language specification: https://github.com/kval-access-language/KVAL 

### Features 

* Single function entry-point:
    * res, err := Query(kb, "INS B1 >> B2 >> B3 >>>> KEY :: VAL") &nbsp; &nbsp; //(will create three buckets, plus k/v in one-go)
    * res, err := Query(kb, "GET B1 >> B2 >> B3 >>>> KEY") &nbsp; &nbsp; &nbsp; //(will retrieve that entry in one-go)
* Start using BoltDB immediately without writing boiler plate before you can code
* KVAL-Parse enables handling of Base64 binary BLOBS
* Regular Expression based searching for key names and values
* [KVAL Language](https://github.com/kval-access-language/KVAL) specifies easy bulk or singleton DEL and RENAME capabilities
* Language specification at highest abstraction, so other bindings for other DBs are hoped for (hint: [NOMS!](https://github.com/attic-labs/noms)) 
* Start working with BoltDB immediately! 

### Usage

Use is simple. There is one function which accepts a string formatted to KVAL's
specification:

    res, err = Query(kb, "GET Bucket One >> Bucket Two >>>> Requested Key")
    if err != nil {
       fmt.Fprintf(os.Stderr, "Error querying db: %v", err)
    } else {
       //Access our (result structure)[https://github.com/kval-access-language/kval-boltdb/blob/master/kval-bolt-structs.go#L16]: res.Result (a map[string]string)
    } 

For write operations we simply check for the existence of an error, else the
operation passed as expected: 

    res, err = Query(kb, "INS Bucket One >> Bucket Two >>>> Insert Key :: Insert Value")
    if err != nil {
       fmt.Fprintf(os.Stderr, "Error querying db: %v", err)
    }

### How easy is it? 

Once you've a connection to a database, call Query as many times as you like to 
work with your data. The most basic implementation, creating a DB, and inserting 
data looks as follows:

	package main

	import (
		"fmt"
		"os"
		kval "github.com/kval-access-language/kval-boltdb"
	)

	func main() {

		kb, err := kval.Connect("newdb.bolt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening bolt database: %#v", err)
			os.Exit(1)
		}
		defer kval.Disconnect(kb)

		//Lets do a test insert...
		res, err := kval.Query(kb, "INS test bucket one >> test bucket two >>>> key one :: value one")
		if err != nil {
			//work with your error
		}
		// else: start working with you res struct
	}

### Demo

Have a look at some of the bits and pieces implemented as part of this binding 
in the demo Go app here: https://github.com/kval-access-language/kval-boltdb-demo 

### How to contribute

I will be starting to use the code in my own work once the dust has settled from thie first tranche of work. As I do that,
and before then, I need the following:

* Comments on the KVAL specification, working towards a version 1.
* Code review
* Code testers - Get using it and report your issues! 
* Spread the word
* Let me know how it goes here via Issue, or on Twitter via [@beet_keeper](https://twitter.com/beet_keeper)

### License

**[GPL Version 3](http://choosealicense.com/licenses/gpl-3.0/)**: https://github.com/kval-access-language/KVAL-BoltDB/blob/master/LICENSE
