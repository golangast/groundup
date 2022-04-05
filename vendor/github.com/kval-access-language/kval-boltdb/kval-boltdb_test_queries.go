package kvalbolt

import "github.com/boltdb/bolt"

//test invalid/non-implemented capabilities

var makeTea = "TEA bucket one >> bucket two >>>> cup :: saucer"

//---------------------------------------------------------------------------//

//test insert procedures
var createInitialState = []string{
	"ins bucket one >> bucket two >> bucket three >>>> test1 :: value1",
	"INS bucket one >> bucket two >> bucket three >>>> test2 :: value2",
	"INS bucket one >> bucket two >> bucket three >>>> test3 :: value3",
	"INS bucket one >> bucket two >>>> test4 :: value4",
	"INS bucket one >> bucket two >>>> test5 :: value5",
	"INS bucket one >>>> test6 :: value6",
	"INS code bucket >>>> code example :: GET bucket one >> bucket two >>>> key1 :: key2",
	"INS regex bucket >>>> regex example one :: middle regex string middle",
	"INS regex bucket >>>> regex example two :: regex string beginning beginning",
	"INS regex bucket >>>> regex example three :: end end regex string",
	"INS regex bucket >>>> regex example four :: regex shouldn't match",
	"INS regex bucket >>>> regex example five :: regex string",
	"INS regex bucket >> regex bucket two >>>> regex example six :: nil bucket test regex string",
}

var insGetBuckets1 = []string{"bucket one", "bucket two", "bucket three"}
var insGetBuckets2 = []string{"bucket one", "bucket two"}
var insGetBuckets3 = []string{"bucket one"}

// Utilise BoltDB Tree statistics.
// KeyN  int // number of keys/value pairs
// Depth int // number of levels in B+tree

var insResult1 = insresult{3, 1}
var insResult2 = insresult{6, 2}
var insResult3 = insresult{8, 3}

type insCheck struct {
	buckets []string
	counts  insresult
}

type insresult struct {
	keys  int
	depth int
}

var i1 = insCheck{insGetBuckets1, insResult1}
var i2 = insCheck{insGetBuckets2, insResult2}
var i3 = insCheck{insGetBuckets3, insResult3}

var insChecksAll = [...]insCheck{i1, i2, i3}

//---------------------------------------------------------------------------//

//test delete procedures

//good delete procedures - nil error expected...
var delkey = "DEL bucket one >> bucket two >> bucket three >>>> test1"         //delete key test1
var nullvalue = "DEL bucket one >> bucket two >> bucket three >>>> test3 :: _" //make value null without deleting key
var delkeys = "del bucket one >> bucket two >> bucket three >>>> _"            //del all keys from a bucket
var delbucket = "DEL bucket one >> bucket two"                                 //delete bucket two

var goodDelResults = [...]string{delkey, nullvalue, delkeys, delbucket}

//bad delete procedures - error of certain types are expected...
var delnonekey = "DEL zero bucket >>>> nonkey"
var delnonebucket = "DEL zero bucket"
var delnonebuckettwo = "DEL bucket one >> zero bucket two"

var delnonekeytwo = "DEL bucket one >>>> nonkey" //silent fail of non-existent key is all BoltDB does
var bucketNoneKey = []string{"bucket one"}

var badDelResults = map[string]error{
	delnonekey:       errNilBucket,
	delnonekeytwo:    nil, //we only get a silent fail, test may have little value, but it's here...
	delnonebucket:    bolt.ErrBucketNotFound,
	delnonebuckettwo: bolt.ErrBucketNotFound,
}

//---------------------------------------------------------------------------//

//test get procedures
var getTest1 = "GET bucket one >> bucket two >> bucket three >>>> test1"
var getTest2 = "GET bucket one >> bucket two >> bucket three >>>> test2"
var getBucketThree = "GET bucket one >> bucket two >> bucket three"
var getBucketOne = "GET bucket one"
var getCodeBucket = "GeT code bucket >>>> code example"
var getRoot = "GET _"

var getSoleResults = map[string]map[string]string{
	getTest1:       {"test1": "value1"},
	getTest2:       {"test2": "value2"},
	getBucketThree: {"test1": "value1", "test2": "value2", "test3": "value3"},
	getBucketOne:   {"bucket two": Nestedbucket, "test6": "value6"},

	// To clarify, the value is a string that represents a KVAL query, not a query proper
	getCodeBucket: {"code example": "GET bucket one >> bucket two >>>> key1 :: key2"},

	// We test this function twice because OCD...
	getRoot: {"bucket one": Nestedbucket, "code bucket": Nestedbucket, "regex bucket": Nestedbucket},
}

//problems testing this, so place in own routine
var getPrimeBucket = "GET _"
var getRootResults = map[string]map[string]string{
	//retrieve the top-level bucket of the database
	getPrimeBucket: {"bucket one": Nestedbucket, "code bucket": Nestedbucket, "regex bucket": Nestedbucket},
}

//---------------------------------------------------------------------------//

//test get regex procedures
//GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> {PAT}
//GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> _ :: Value
//GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> _ :: {PAT}
var getRegexTest1 = "GET bucket one >> bucket two >> bucket three >>>> {^test\\d+$}"
var getRegexTest2 = "GET bucket one >> bucket two >> bucket three >>>> _ :: value3"
var getRegexTest3 = "GET regex bucket >>>> _ :: {regex string}"

var regexRes1 = map[string]string{"test1": "value1", "test2": "value2", "test3": "value3"}
var regexRes2 = map[string]string{"test3": "value3"}
var regexRes3 = map[string]string{"regex example one": "middle regex string middle", "regex example two": "regex string beginning beginning", "regex example three": "end end regex string", "regex example five": "regex string"}

var getRegexResults = map[string]map[string]string{
	//var get_sole_results = map[string]map[string]string {
	getRegexTest1: regexRes1,
	getRegexTest2: regexRes2,
	getRegexTest3: regexRes3,
}

//---------------------------------------------------------------------------//

//example Kvalresults
//a: Kvalresult{map[string]string{"test1": "value1", "test2": "value2", "test3": "value3"}, false},
//b: Kvalresult{map[string]string{"bucket two": NESTEDBUCKET, "test6": "value6"}, false},

//test list procedures
var lisBucketTwo = "LIS bucket one >> bucket two"
var lisTest1 = "LIS bucket one >> bucket two >> bucket three >>>> test1"
var lisUnknownKey = "LIS bucket one >> bucket two >> bucket three >>>> nokey"
var lisUnknownBucket = "LIS ins1 >> ins2 >> no-bucket"

var lisResults = map[string]bool{
	lisBucketTwo:     true,
	lisTest1:         true,
	lisUnknownKey:    false,
	lisUnknownBucket: false,
}

//---------------------------------------------------------------------------//

//test rename procedures
var renameState = []string{
	"INS ren1 >> ren2 >> ren3 >>>> r1 :: v1",         //2
	"INS ren1 >> ren2 >> ren3 >>>> r2 :: v2",         //3
	"INS ren1 >> ren2 >> ren3 >> ren4",               //4
	"INS ren1 >> ren2 >> ren3 >> ren4 >>>> r3 :: v3", //5
	"INS ren1 >> ren2 >> ren3 >> ren4 >>>> r4 :: v4", //6
	"INS ren1 >> ren2 >> ren3 >> ren4 >>>> r5 :: v5", //7
	"INS ren1 >> ren2 >> ren3 >>>> r6 :: v6",         //8
	"INS ren1 >> ren2 >>>> r6 :: v6",                 //9
	"INS ren1 >> ren2 >>>> r7 :: v6",                 //10
	"INS ren1 >> ren2 >>>> r8 :: v6",                 //11
	"INS ren1 >> ren2 >>>> r1 :: v1",                 //12
	"INS ren1 >> renamekey >>>> key :: value",        //key to rename...
}

var r1 = "ren_key"
var r2 = "ren_bucket"

//though few, this should prove our capability adequately...
var renameTests = map[string]string{
	r1: "REN ren1 >> renamekey >>>> key => newkey", //rename key
	r2: "REN ren1 >> ren2 => rnew",                 //rename bucket
}

//FALSE :: TRUE if rename has worked, we'll see true for second value
var renLis1 = [2]string{"LIS ren1 >> renamekey >>>> key", "LIS ren1 >> renamekey >>>> newkey"}
var renLis2 = [2]string{"LIS ren1 >> ren2", "LIS ren1 >> rnew"}

//grab stats for our rename functions to make sure the test works by using
//mix of list queries and our statsdb capabilities, testing two code features
var renOldList = "LIS ren1 >> ren2"
var renNewList = "LIS ren1 >> rnew"

//---------------------------------------------------------------------------//
