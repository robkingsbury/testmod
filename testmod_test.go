package testmod

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	o := Opts{
		Name:  "Doug",
		Times: 23,
	}
	f := New(o)
	want := "Fletched 23 times for Doug!"

	if got := fmt.Sprintf("%s", f); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
