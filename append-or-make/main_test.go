package main

import (
	"fmt"
	"testing"
)

var tests = struct {
	lp []int
	la []int
}{
	lp: make([]int, 11),
	la: make([]int, 11),
}

func init() {
	for i := range [...]int{10: 0} {
		tests.lp[i] = 1 << uint(i)
		tests.la[i] = 1 << uint(i)
	}
}

func BenchmarkAppend(b *testing.B) {
	for _, lp := range tests.lp {
		for _, la := range tests.la {
			a := make([]int, la)
			name := fmt.Sprintf("%v/%v", lp, la)
			b.Run(name, func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					p := make([]int, lp)
					for i := range p {
						p[i] = i * i
					}
					p = append(p, a...)
				}
			})
		}
	}
}

func BenchmarkMake(b *testing.B) {
	for _, lp := range tests.lp {
		for _, la := range tests.la {
			a := make([]int, la)
			name := fmt.Sprintf("%v/%v", lp, la)
			b.Run(name, func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					p := make([]int, lp)
					for i := range p {
						p[i] = i * i
					}
					v := make([]int, len(p)+len(a))
					copy(v, p)
					for i := len(p); i < len(p)+len(a); i++ {
						v[i] = a[i-len(p)]
					}
					p = v
				}
			})
		}
	}
}
