/*
Package kvalscanner implements a basic lexical scanner for KVAL (Key Value
Access Language). The language has been created for more fluid access to
Key-Value-like databases such as the popular BoltDb.

Parsers can access a new Scanner() to check that a user provided string is
valid according to the language's specification. For Golang, however, this
work is already done for you at http://github.com/kval-access-language/kval-parse

The most up-to-date specification for KVAL can be found here:
https://github.com/kval-access-language/kval
*/
package kvalscanner
