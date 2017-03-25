package main

import "testing"

func BenchmarkMake(b *testing.B) {
	slice := []int{99: 1}
	newSlice := []int{100: 1}
	for i := 0; i < b.N; i++ {
		slice := make([]int, len(newSlice))
		copy(slice, newSlice)
	}
	_ = slice
}

func BenchmarkAppend(b *testing.B) {
	slice := []int{99: 1}
	newSlice := []int{100: 1}
	for i := 0; i < b.N; i++ {
		slice = append(slice[:0], newSlice...)
	}
	_ = slice
}

func BenchmarkAppendFullSlice(b *testing.B) {
	slice := []int{99: 1}
	newSlice := []int{100: 1}
	for i := 0; i < b.N; i++ {
		slice = append(slice[:0:0], newSlice...)
	}
	_ = slice
}
