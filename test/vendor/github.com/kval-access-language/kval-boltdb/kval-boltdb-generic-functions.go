package kvalbolt

import (
	"github.com/boltdb/bolt"
	"github.com/kval-access-language/kval-scanner"
	"github.com/pkg/errors"
)

// Constant values for statdb to be able to externalize bucket stats
// for anyone using kvalresults structs without having to go back to bolt ptr
const (
	opcodeNormal       int = iota // opcodeNormal refers to LIS, GET, INS, and certain REN/DEL functions
	opcodeDelBucket               // delete functions that delete a bucket thus can't be handled as easily
	opcodeRenameBucket            // rename functions that rename a bucket thus can't be handled as easily
	opcodeRoot                    // we're working with Root bucket, e.g. can't stat Root/boltdb.Tx
	opcodeUnknown                 // cannot stat for an unknown function
)

//get boltdb bucket stats
func getbucketstats(kb Kvalboltdb, buckets []string) (bolt.BucketStats, error) {
	var bs bolt.BucketStats
	err := kb.DB.View(func(tx *bolt.Tx) error {
		bucket, err := gotobucket(tx, buckets)
		if err != nil {
			return err
		}
		bs = bucket.Stats()
		return err
	})
	return bs, err
}

// statdb allows kvalresult structures to be annotated with boltdb bucket stats
func statdb(kb Kvalboltdb, kr Kvalresult) (Kvalresult, error) {

	var err error
	kbuckets := kb.query.Buckets
	klen := len(kbuckets)

	if kb.query.Function == kvalscanner.LIS && !kr.Exists {
		// e.g. a LIS query for which a null result has been returned
		// our query is about the existence of something, stats can be
		// returned a better way than a search for non-existence
		kr.Stats = bolt.BucketStats{}
		return kr, nil
	}

	switch kr.opcode {
	case opcodeRoot:
		// cannot stat boltdb.Tx/Root bucket
		kr.Stats = bolt.BucketStats{}
		return kr, nil
	case opcodeNormal:
		kr.Stats, err = getbucketstats(kb, kbuckets)
		if err != nil {
			kr.Stats = bolt.BucketStats{}
			return kr, errors.Wrap(err, "Getting stats for regular operations resulted in error")
		}
	case opcodeDelBucket:
		if klen > 1 {
			//we need to navigate to the bucket we didn't delete...
			kbuckets := kbuckets[:klen-1]
			kr.Stats, err = getbucketstats(kb, kbuckets)
			if err != nil {
				kr.Stats = bolt.BucketStats{}
				return kr, errors.Wrap(err, "Getting stats for delete bucket functions resulted in error")
			}
		} else {
			//we're trying to stat from level zero of db, that is we've
			//deleted the only bucket we were interested in for this query
			//so no stats to be returned that are relvant to the query executed
			kr.Stats = bolt.BucketStats{}
		}
	case opcodeRenameBucket:
		//last bucket has to take on the new bucket name
		kbuckets[klen-1] = kb.query.Newname
		kr.Stats, err = getbucketstats(kb, kbuckets)
		if err != nil {
			kr.Stats = bolt.BucketStats{}
			return kr, errors.Wrap(err, "Getting stats for rename bucket functions resulted in error")
		}
	default:
		kr.Stats = bolt.BucketStats{}
		return kr, errors.New("Invalid opcode for statdb()")
	}
	return kr, nil
}
