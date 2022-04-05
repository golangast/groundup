/*
Package kvalbolt implements a BoltDB binding for KVAL (Key Value Access Language)

The binding provides more than just 'boilerplate' for working with BoltDB. It
implements a language thus enabling access from a higher point of abstraction.
Users of the binding can provide simple instructions and the binding will take
care of Bucket creation (and deletion), and the generation of keys and values,
plus their maintenance, no matter how deep into the database structure they
are needed, or indeed exist.

"Everything should be made as simple as possible, but no simpler." - Albert Einstein

The most up-to-date information on the project can be found on the GitHub.com
README: https://github.com/kval-access-language/kval-boltdb/blob/master/README.md

The most up-to-date specification for KVAL can be found here:
https://github.com/kval-access-language/kval
*/
package kvalbolt
