# KVAL-Parse

[![Build Status](https://travis-ci.org/kval-access-language/kval-parse.svg?branch=master)](https://travis-ci.org/kval-access-language/kval-parse)
[![GoDoc](https://godoc.org/github.com/kval-access-language/kval-parse?status.svg)](https://godoc.org/github.com/kval-access-language/kval-parse)
[![Go Report Card](https://goreportcard.com/badge/github.com/kval-access-language/kval-parse)](https://goreportcard.com/report/github.com/kval-access-language/kval-parse)

Parser written in Golang for a simple key value access language. 

### Key Value Access Language

I have created a modest specification for a key value access langauge. 
It allows for input and access of values to a key value store such as Golang's
[BoltDB](https://github.com/boltdb/). 

The language specification: https://github.com/kval-access-language/KVAL 

### Usage

Import the library and run the Parse function() e.g: 

    var query "GET Bucket One >> Bucket Two >>>> Requested Key"
    kq, err := kvalparse.Parse(query)
    if err != nil {
       return kr, err
    }

If we find an error we have an invalid query. Results are returned in a query
structure:

    type KQUERY struct { 
       Function Token
       Buckets []string  
       Key string
       Value string
       Newname string
       Regex bool
    }

### License

**[GPL Version 3](http://choosealicense.com/licenses/gpl-3.0/)**: https://github.com/kval-access-language/KVAL-BoltDB/blob/master/LICENSE
