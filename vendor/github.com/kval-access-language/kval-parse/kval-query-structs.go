package kvalparse

import "github.com/kval-access-language/kval-scanner"

// KQuery struct used to breakdown a query string in a way that can
// be utilised by the calling code more easily.
type KQuery struct {
	Function kvalscanner.Token // Function being called from KVAL
	Buckets  []string          // The Bucket or Buckets the query wants to create/interrogate
	Key      string            // The Key the query wants to create/interrogate
	Value    string            // The Value the query wants to create/interrogate
	Newname  string            // If renaming a Bucket or Key, the new name is found here
	Regex    bool              // If regex is involved in the query anywhere, True here
}
