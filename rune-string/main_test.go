package main

import (
	"strings"
	"testing"

	"github.com/mattn/go-runewidth"
)

var s = []rune(strings.Repeat("A", 10))

func runesWidth(s []rune) (width int) {
	for _, r := range s {
		width += runewidth.RuneWidth(r)
	}
	return width
}

func BenchmarkRune(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n += runesWidth(s)
	}
}

func BenchmarkString(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n += runewidth.StringWidth(string(s))
	}
}
