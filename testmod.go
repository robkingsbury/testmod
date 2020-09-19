// Package testmod does nothing useful other than a way for me to play with Go
// modules and github.
package testmod

import (
	"fmt"
)

// A Fletch remembers how many times a name has been, um, fletched.
type Fletch struct {
	times int
	name  string
}

// Opts is used to set parameters for the Fletch returned by New().
type Opts struct {
	// Times is the number of times something has been fletched.
	Times int

	// Name is the identifier for the thing being fletched.
	Name string
}

// New returns an instance of Fletch with parameters set by the given Opts.
func New(o Opts) *Fletch {
	return &Fletch{
		times: o.Times,
		name:  o.Name,
	}
}

// String implements the Stringer interface.
func (f *Fletch) String() string {
	return fmt.Sprintf("Fletched %d times for %s!", f.times, f.name)
}
