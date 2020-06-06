package cimg_go

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	want := "foobarhoge"
	in := []string{
		"foo",
		"bar",
		"hog  e",
	}
	var mb MyBuilder
	// Reserve spaces.
	mb.Grow(40)
	for _, s := range in {
		fmt.Fprint(&mb, s)
	}

	if got := mb.String(); got != want {
		t.Errorf("want %s, but got = %s\n", want, got)
	}
}
