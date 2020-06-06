package cimg_go

import (
	"fmt"
	"strings"
	"testing"
)

func TestBuilder(t *testing.T) {
	want := "foobarhoge"
	in := []string{
		"foo",
		"bar",
		"hoge",
	}
	var sb strings.Builder
	// Reserve spaces.
	sb.Grow(40)
	for _, s := range in {
		fmt.Fprint(&sb, s)
	}

	got := sb.String()

	if got != want {
		t.Errorf("want %s, but got = %s\n", want, got)
	}
}
