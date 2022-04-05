# KVAL-Scanner

[![Build Status](https://travis-ci.org/kval-access-language/kval-scanner.svg?branch=master)](https://travis-ci.org/kval-access-language/kval-scanner)
[![GoDoc](https://godoc.org/github.com/kval-access-language/kval-scanner?status.svg)](https://godoc.org/github.com/kval-access-language/kval-scanner)
[![Go Report Card](https://goreportcard.com/badge/github.com/kval-access-language/kval-scanner)](https://goreportcard.com/report/github.com/kval-access-language/kval-scanner)

Lexical scanner for KVAL (Key Value Access Language) for KVAL-Parser 

### Key Value Access Language

I have created a modest specification for a key value access langauge. 
It allows for input and access of values to a key value store such as Golang's
[BoltDB](https://github.com/boltdb/). 

The language specification: https://github.com/kval-access-language/KVAL 

### Usage

Usage is pretty low-level so will normally be done via [KVAL-Parse](https://github.com/kval-access-language/kval-parse) so please see there for examples.

Scanner library is maintained in its own repo so that changes between the two can be viewed as purely as possible.

### License

**[GPL Version 3](http://choosealicense.com/licenses/gpl-3.0/)**: https://github.com/kval-access-language/KVAL-BoltDB/blob/master/LICENSE
