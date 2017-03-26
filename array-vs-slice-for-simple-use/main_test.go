package main

import "testing"

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = [...]int{100: 0}
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []int{100: 0}
	}
}
