package maybe

import (
	"testing"
)

func TestOption(t *testing.T) {
	a := Some(3)
	b := None[int]()

	if a.UnwrapOrDefault() != 3 {
		t.Error("a should be some 3")
	}
	if b.UnwrapOrDefault() != 0 {
		t.Error("b should be none")
	}
}
