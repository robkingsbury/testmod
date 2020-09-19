package sys

import (
	"os"
)

var errHostname = "unknown"

// Hostname is a wrapper around os.Hostname that returns the string "unknown"
// if os.Hostname() returns an error.
func Hostname() string {
	return noErrHostname(os.Hostname())
}

func noErrHostname(hostname string, err error) string {
	if err != nil {
		return errHostname
	}
	return hostname
}
