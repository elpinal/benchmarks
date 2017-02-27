package main

import "testing"

var list = [][]int{{1, 1, 1, 1, 1, 1, 1}, {2, 2, 2, 2, 2, 2, 2}}

func BenchmarkReuse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var l int
		var n []int
		for _, i := range list {
			n = i
			l += len(n)
		}
	}
}

func BenchmarkNoReuse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var l int
		for _, i := range list {
			n := i
			l += len(n)
		}
	}
}
