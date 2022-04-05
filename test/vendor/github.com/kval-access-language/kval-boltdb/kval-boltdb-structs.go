package kvalbolt

import (
	"github.com/boltdb/bolt"
	"github.com/kval-access-language/kval-parse"
	"strings"
)

// Nestedbucket, const to help users validate Kvalblob struct
const Nestedbucket = "NestedBucket"

// Datam const to help users validate Kvalblob struct
const Data = "data"

// Base64, const to help users validate Kvalblob struct
const Base64 = "base64"

// Unexported const to validate Kvalblob "data:<mimetype>:<encoding type>:<data>"
const bloblen = 4

// Kvalboltdb represents a parsed query structure that we can pass aound in code
type Kvalboltdb struct {
	DB    *bolt.DB         // Pointer to BoltDB, users can access directly or via function call
	Fname string           // Filename used to create our DB with
	query kvalparse.KQuery // Parsed query string returned as struct for manipulation
}

// Kvalresult provides a mechanism for users to interact with results.
// It also allows for a wide-variety of results from single GET results, to
// retrieving all from a Bucket through various different mechanisms.
// Maplen will be one if a single result. Users should understand their data, but
// also be curious to the results they're getting from their various queries,
// therefore checking this length is a good approach.
type Kvalresult struct {
	Result map[string]string // Result map to access various types of result post-query
	Exists bool              // When using a LIS query this is set to true if we can find our data
	Stats  bolt.BucketStats  // BoltDB bucket stats relevant to the operation performed by the user
	opcode int               // Additional non-user metadata about the type of operation performed by the query
}

// Kvalblob struct to allow users to work with Base64 encoded blobs
type Kvalblob struct {
	Query    string // Where valid, the query used is stored here
	Datatype string // Datatype is stored here, expected: 'data'
	Mimetype string // Mimetype is recorded here and is up to users to insert correctly
	Encoding string // Encoding is stored here, expected:'base64'
	Data     string // Our data is stored here as a base64 string
}

func initKvalblob(query string, mimetype string, data string) Kvalblob {
	return Kvalblob{query, Data, mimetype, Base64, data}
}

func queryfromkvb(kvb Kvalblob) string {
	query := kvb.Query + " :: " + kvb.Datatype + ":" + kvb.Mimetype + ":" + kvb.Encoding + ":" + kvb.Data
	return query
}

func blobfromKvalresult(kv Kvalresult) (Kvalblob, error) {
	var kvb Kvalblob
	for k, v := range kv.Result {
		kvb.Query = k
		reslice := strings.Split(v, ":")
		if len(reslice) != 4 {
			return kvb, errBlobLen
		}
		kvb.Datatype = reslice[0]
		kvb.Mimetype = reslice[1]
		kvb.Encoding = reslice[2]
		kvb.Data = reslice[3]
	}
	return kvb, nil
}
