package kvalparse

var kvalBoltdbVersion = "0.0.0-KVAL-Working-Draft"

// Version will return an indication of which version of the KVAL language we
// are working from and the version of the library that you are implementing from
func Version() string {
	return kvalBoltdbVersion
}
