package testmod

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	f := New(2)
	want := "fletched 2 times"

	if got := fmt.Sprintf("%s", f); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
