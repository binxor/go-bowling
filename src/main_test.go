package gobowling

import (
	"testing"
)

func TestScore(t *testing.T) {
	want := 5
	if got := Score("5"); got != want {
		t.Errorf("Score(5) = %q, want %q", got, want)
	}
}
