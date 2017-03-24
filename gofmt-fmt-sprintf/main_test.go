package main

// from: func replaceTempFilename(diff []byte, filename string) ([]byte, error)
//   in cmd/gofmt/gofmt.go

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

var result int

func BenchmarkFmtSprintf(b *testing.B) {
	var bs []byte
	f := "path/to/file.go"
	ts := []byte(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC).Format("\t2006-01-02 15:04:05.000000000 -0700"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bs = []byte(fmt.Sprintf("--- %s%s", f+".orig", ts))
		bs = []byte(fmt.Sprintf("+++ %s%s", f, ts))
	}
	result = len(bs)
}

func BenchmarkPlus(b *testing.B) {
	var bs []byte
	f := "path/to/file.go"
	ts := []byte(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC).Format("\t2006-01-02 15:04:05.000000000 -0700"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bs = append([]byte("--- "+f+".orig"), ts...)
		bs = append([]byte("+++ "+f), ts...)
	}
	result = len(bs)
}

func BenchmarkBuffer(b *testing.B) {
	var bs []byte
	f := "path/to/file.go"
	ts := []byte(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC).Format("\t2006-01-02 15:04:05.000000000 -0700"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		b.Write([]byte("--- "))
		b.WriteString(f)
		b.Write([]byte(".org"))
		b.Write(ts)
		bs = b.Bytes()

		b.Reset()
		b.Write([]byte("+++ "))
		b.WriteString(f)
		b.Write(ts)
		bs = b.Bytes()
	}
	result = len(bs)
}
