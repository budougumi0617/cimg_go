package cimg_go

import "strings"

type MyBuilder struct {
	sb strings.Builder
}

func (b *MyBuilder) Write(p []byte) (n int, err error) {
	return b.sb.Write(p)
}

func (b *MyBuilder) Grow(n int) {
	b.sb.Grow(n)
}

func (b *MyBuilder) String() string{
	return b.sb.String()
}