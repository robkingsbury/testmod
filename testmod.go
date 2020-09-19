package testmod

import (
	"fmt"
)

type Fletch struct {
	i int
}

func New(i int) *Fletch {
	return &Fletch{i: i}
}

func (f *Fletch) String() string {
	return fmt.Sprintf("fletched %d times", f.i)
}
