package sys

import (
	"errors"
	"testing"
)

func TestNoErrHostname(t *testing.T) {
	tests := []struct {
		name     string
		hostname string
		err      error
		want     string
	}{
		{
			name:     "no error",
			hostname: "scrub",
			err:      nil,
			want:     "scrub",
		},
		{
			name:     "empty string",
			hostname: "",
			err:      nil,
			want:     "",
		},
		{
			name:     "error exists",
			hostname: "scrub",
			err:      errors.New("some error"),
			want:     errHostname,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := noErrHostname(test.hostname, test.err)
			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}
