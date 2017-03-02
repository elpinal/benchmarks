package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

var s []rune

func TestMain(m *testing.M) {
	benchmarks := []int{1, 4, 16, 64, 256}
	for _, bm := range benchmarks {
		s = []rune(strings.Repeat("a", bm))
		m.Run()
	}
}

func BenchmarkIOWriteString(b *testing.B) {
	var w bytes.Buffer
	for i := 0; i < b.N; i++ {
		io.WriteString(&w, string(s))
	}
}

func BenchmarkCopyThenWrite(b *testing.B) {
	var w bytes.Buffer
	for i := 0; i < b.N; i++ {
		b := make([]byte, len(s))
		copy(b, string(s))
		w.Write(b)
	}
}

func BenchmarkBufioWriteString(b *testing.B) {
	var w bytes.Buffer
	bw := bufio.NewWriterSize(&w, 32)
	for i := 0; i < b.N; i++ {
		bw.WriteString(string(s))
		bw.Flush()
	}
}

func BenchmarkBufioWriteRune(b *testing.B) {
	var w bytes.Buffer
	bw := bufio.NewWriterSize(&w, 32)
	for i := 0; i < b.N; i++ {
		for _, r := range s {
			bw.WriteRune(r)
		}
		bw.Flush()
	}
}
