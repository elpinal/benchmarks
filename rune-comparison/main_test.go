package main

import (
	"reflect"
	"strings"
	"testing"
)

var (
	r1 = []rune(strings.Repeat("a", 0))
	r2 = []rune(strings.Repeat("a", 0))

	r3 = []rune(strings.Repeat("a", 10))
	r4 = []rune(strings.Repeat("a", 10))

	r5 = []rune(strings.Repeat("a", 100))
	r6 = []rune(strings.Repeat("a", 100))

	r7 = []rune(strings.Repeat("a", 101))

	r8 = []rune("aa" + strings.Repeat("a", 98))
	r9 = []rune("ab" + strings.Repeat("a", 98))
)

func forRune(r1, r2 []rune) bool {
	if len(r1) != len(r2) {
		return false
	}
	for n, r := range r1 {
		if r2[n] == r {
			return false
		}
	}
	return true
}

func coerceToString(r1, r2 []rune) bool {
	return string(r1) == string(r2)
}

func deepEqual(r1, r2 []rune) bool {
	return reflect.DeepEqual(r1, r2)
}

func BenchmarkFor(b *testing.B) {
	b.Run("0", func(b *testing.B) { benchmarkFor(b, r1, r2) })
	b.Run("10", func(b *testing.B) { benchmarkFor(b, r3, r4) })
	b.Run("100", func(b *testing.B) { benchmarkFor(b, r5, r6) })
	b.Run("100:101", func(b *testing.B) { benchmarkFor(b, r6, r7) })
	b.Run("100(self)", func(b *testing.B) { benchmarkFor(b, r5, r5) })
	b.Run("100(diff)", func(b *testing.B) { benchmarkFor(b, r8, r9) })
}

func benchmarkFor(b *testing.B, r1, r2 []rune) {
	for i := 0; i < b.N; i++ {
		_ = forRune(r1, r2)
	}
}

func BenchmarkString(b *testing.B) {
	b.Run("0", func(b *testing.B) { benchmarkString(b, r1, r2) })
	b.Run("10", func(b *testing.B) { benchmarkString(b, r3, r4) })
	b.Run("100", func(b *testing.B) { benchmarkString(b, r5, r6) })
	b.Run("100:101", func(b *testing.B) { benchmarkString(b, r6, r7) })
	b.Run("100(self)", func(b *testing.B) { benchmarkString(b, r5, r5) })
	b.Run("100(diff)", func(b *testing.B) { benchmarkString(b, r8, r9) })
}

func benchmarkString(b *testing.B, r1, r2 []rune) {
	for i := 0; i < b.N; i++ {
		_ = coerceToString(r1, r2)
	}
}

func BenchmarkDeepEqual(b *testing.B) {
	b.Run("0", func(b *testing.B) { benchmarkDeepEqual(b, r1, r2) })
	b.Run("10", func(b *testing.B) { benchmarkDeepEqual(b, r3, r4) })
	b.Run("100", func(b *testing.B) { benchmarkDeepEqual(b, r5, r6) })
	b.Run("100:101", func(b *testing.B) { benchmarkDeepEqual(b, r6, r7) })
	b.Run("100(self)", func(b *testing.B) { benchmarkDeepEqual(b, r5, r5) })
	b.Run("100(diff)", func(b *testing.B) { benchmarkDeepEqual(b, r8, r9) })
}

func benchmarkDeepEqual(b *testing.B, r1, r2 []rune) {
	for i := 0; i < b.N; i++ {
		_ = deepEqual(r1, r2)
	}
}
