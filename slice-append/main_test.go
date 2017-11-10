package main

import "testing"

var t []int
var s []int

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s = make([]int, 8, 8)
		t = append(s, 0)
	}
}

func BenchmarkAppendFirst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s = make([]int, 8, 8)
		t = append([]int{0}, s...)
	}
}
