package kvalparse

import "github.com/pkg/errors"

var errInvalidFunction = errors.New("Attempting to parse invalid function")
var errZeroBuckets = errors.New("Zero buckets: No buckets specified in input query")
var errInsertRegex = errors.New("Invalid Pattern use: Can't have regex on insert")
var errKeyGetRegex = errors.New("Known Value: No need to GET a known value")
var errKeyLisRegex = errors.New("Known Value: No need to LIS a known value")
var errUnknownUnknown = errors.New("Unknown unknown: Cannot seek unknown key and value")
var errNoNameRename = errors.New("Rename: Missing newname parameter")
var errCompileRegex = errors.New("Invalid regex: Cannot compile regular expression")
var errParsedNoNewTokens = errors.New("Invalid query: Parsed without finding any new tokens")
var errIllegalToken = errors.New("Illegal token in query string")
