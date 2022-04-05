package kvalparse

import "github.com/kval-access-language/kval-scanner"

/*
type KQUERY struct {
   function kvalscanner.Token
   buckets []string
   key string
   value string
   newname string
   regex bool
}
*/

var (
	kq01 = KQuery{kvalscanner.INS, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "Value", "", false}
	kq02 = KQuery{kvalscanner.INS, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
	kq03 = KQuery{kvalscanner.GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
	kq04 = KQuery{kvalscanner.GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "", false}
	kq05 = KQuery{kvalscanner.GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "PAT", "", "", true}
	kq06 = KQuery{kvalscanner.GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "_", "Value", "", false}
	kq07 = KQuery{kvalscanner.GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "_", "PAT", "", true}
	kq08 = KQuery{kvalscanner.LIS, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "", false}
	kq09 = KQuery{kvalscanner.LIS, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
	kq0a = KQuery{kvalscanner.DEL, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
	kq0b = KQuery{kvalscanner.DEL, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "", false}
	kq0c = KQuery{kvalscanner.DEL, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "_", "", false}
	kq0d = KQuery{kvalscanner.REN, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "New Key", false}
	kq0e = KQuery{kvalscanner.REN, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "Third Bucket", false}
	//non-reference-queries
	kq0f = KQuery{kvalscanner.INS, []string{"Prime Bucket"}, "", "", "", false}
	kq10 = KQuery{kvalscanner.INS, []string{"Prime Bucket"}, "key", "", "", false}
	kq11 = KQuery{kvalscanner.INS, []string{"Prime Bucket"}, "key", "value", "", false}
	kq12 = KQuery{kvalscanner.INS, []string{"Prime Bucket"}, "key", "hyphen-value", "", false}
	kq13 = KQuery{kvalscanner.GET, []string{"Prime Bucket"}, "_", "PATT WITH THREE SPACES", "", true}
	kq14 = KQuery{kvalscanner.GET, []string{"Prime Bucket"}, "PATT WITH THREE SPACES", "", "", true}
	kq15 = KQuery{kvalscanner.INS, []string{"Prime Bucket"}, "key", "value with space", "", false}
	kq16 = KQuery{kvalscanner.INS, []string{"link index"}, "internet archive latest", "http://web.archive.org/web/20170328100131/http://www.bbc.co.uk/news/", "", false}
	kq17 = KQuery{kvalscanner.INS, []string{"link index"}, "internet archive response code", "200", "", false}
	kq18 = KQuery{kvalscanner.GET, []string{"Prime Bucket"}, "", "", "", false}
	kq19 = KQuery{kvalscanner.GET, []string{"_"}, "", "", "", false}
	kq20 = KQuery{kvalscanner.INS, []string{"Prime Bucket"}, "key", "value\r\nvalue", "", false}
)

//Queries that should work according to the KVAL specification
var goodQueryMap = map[string]string{
	"kq01_insert_value":               "ins Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key :: Value",
	"kq02_insert_value":               "INS Prime Bucket >> Secondary Bucket >> Tertiary Bucket",
	"kq03_get_bucket_contents":        "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket",
	"kq04_get_value":                  "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key",
	"kq05_get_value_from_key_pattern": "get Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> {PAT}",
	"kq06_get_key_from_value":         "GeT Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> _ :: Value",
	"kq07_get_key_from_value_pattern": "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> _ :: {PAT}",
	"kq08_does_key_exist":             "LIS Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key",
	"kq09_does_bucket_exist":          "LIS Prime Bucket >> Secondary Bucket >> Tertiary Bucket",
	"kq0a_delete_bucket":              "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket",
	"kq0b_delete_key":                 "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key",
	"kq0c_delete_value":               "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key :: _ ",
	"kq0d_rename_key":                 "REN Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key => New Key",
	"kq0e_rename_bucket":              "REN Prime Bucket >> Secondary Bucket >> Tertiary Bucket => Third Bucket",
	"kq0f_one_bucket":                 "INS Prime Bucket",
	"kq10_one_bucket_key":             "INS Prime Bucket >>>> key",
	"kq11_one_bucket_key_value":       "INS Prime Bucket >>>> key :: value",
	"kq12_one_bucket_hyphen_value":    "INS Prime Bucket >>>> key :: hyphen-value",
	"kq13_regex_spaces_value":         "GET Prime Bucket >>>> _ :: {PATT WITH THREE SPACES}",
	"kq14_regex_spaces_key":           "GET Prime Bucket >>>> {PATT WITH THREE SPACES}",
	"kq15_value_spaces":               "INS Prime Bucket >>>> key :: value with space",
	"kq16_value_hyperlink":            "INS link index >>>> internet archive latest :: http://web.archive.org/web/20170328100131/http://www.bbc.co.uk/news/",
	"kq17_value_number":               "INS link index >>>> internet archive response code :: 200",
	"kq18_first_bucket":               "GET Prime Bucket",
	"kq19_prime_bucket":               "GET _",
	"kq20_one_bucket_key_value":       "INS Prime Bucket >>>> key :: value\r\nvalue",
}

var goodQueryExpected = map[string]KQuery{
	"kq01_insert_value":               kq01,
	"kq02_insert_value":               kq02,
	"kq03_get_bucket_contents":        kq03,
	"kq04_get_value":                  kq04,
	"kq05_get_value_from_key_pattern": kq05,
	"kq06_get_key_from_value":         kq06,
	"kq07_get_key_from_value_pattern": kq07,
	"kq08_does_key_exist":             kq08,
	"kq09_does_bucket_exist":          kq09,
	"kq0a_delete_bucket":              kq0a,
	"kq0b_delete_key":                 kq0b,
	"kq0c_delete_value":               kq0c,
	"kq0d_rename_key":                 kq0d,
	"kq0e_rename_bucket":              kq0e,
	"kq0f_one_bucket":                 kq0f,
	"kq10_one_bucket_key":             kq10,
	"kq11_one_bucket_key_value":       kq11,
	"kq12_one_bucket_hyphen_value":    kq12,
	"kq13_regex_spaces_value":         kq13,
	"kq14_regex_spaces_key":           kq14,
	"kq15_value_spaces":               kq15,
	"kq16_value_hyperlink":            kq16,
	"kq17_value_number":               kq17,
	"kq18_first_bucket":               kq18,
	"kq19_prime_bucket":               kq19,
	"kq20_one_bucket_key_value":       kq20,
}

var badQueryMap = map[string]string{
	"badkq01_no_buckets":  "INS",
	"badkq02_ins_regex":   "INS Prime Bucket >>>> {PATT}",
	"badkq03_ins_regex":   "INS Prime Bucket >>>> key :: {PATT}",
	"badkq04_ins_regex":   "INS Prime Bucket >>>> {PATT} :: {PATT}",
	"badkq05_get_val":     "GET Prime Bucket >>>> known :: unknown", //if we know value, we don't need get
	"badkq06_lis_val":     "LIS Prime Bucket >>>> known :: unknown", //validate for yourself, for many reasons!
	"badkq07_get_unknown": "GET Prime Bucket >>>> _",
	"badkq08_lis_unknown": "LIS Prime Bucket >>>> _",
	"badkq09_ren_bucket":  "REN Prime Bucket => ",
	"badkq0a_ren_key":     "REN Prime Buckey >>>> Key => ",
}

var badQueryExpected = map[string]error{
	"badkq01_no_buckets":  errZeroBuckets,
	"badkq02_ins_regex":   errInsertRegex,
	"badkq03_ins_regex":   errInsertRegex,
	"badkq04_ins_regex":   errInsertRegex,
	"badkq05_get_val":     errKeyGetRegex,
	"badkq06_lis_val":     errKeyLisRegex,
	"badkq07_get_unknown": errUnknownUnknown,
	"badkq08_lis_unknown": errUnknownUnknown,
	"badkq09_ren_bucket":  errNoNameRename,
	"badkq0a_ren_key":     errNoNameRename,
}
