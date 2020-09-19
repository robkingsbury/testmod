package testmod

import (
	"fmt"
)

type Fletch struct {
	times int
	name  string
}

type Opts struct {
	Times int
	Name  string
}

func New(o Opts) *Fletch {
	return &Fletch{
		times: o.Times,
		name:  o.Name,
	}
}

func (f *Fletch) String() string {
	return fmt.Sprintf("fletched %d times for %s",
		f.times, f.name)
}
