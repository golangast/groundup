package kvalbolt

import "github.com/pkg/errors"

var errNilBucket = errors.New("Cannot GOTO bucket, bucket not found")
var errEmptyBucketSlice = errors.New("Cannot GOTO bucket, empty buckets slice provided")
var errNotImplemented = errors.New("KVAL Function not implemented")
var errNoKVInBucket = errors.New("No Keys: There are no key::value pairs in this bucket")

var errBlobKey = errors.New("No Key: attempting to add blob but key value is empty or '_'")
var errBlobVal = errors.New("Value added: attempting to add blob but have specified value")
var errBlobIns = errors.New("INS Only: Can only use INS to PUT blob")
var errBlobLen = errors.New("Blob data supplied is not a blob, or is wrapped incorrectly")
var errBlobMapLen = errors.New("Maplen not equal to one. May be multiple value context, or zero")

var errStat = errors.New("Cannot stat database. Stats struct will remain empty.")

//Other non-error error strings...
var errParse = "Query parse failed"
