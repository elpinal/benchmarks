package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"testing"
)

// *myWriter implements only io.Writer, not io.stringWriter.
// So io.WriteString for *myWriter is slower than for *bytes.Buffer which
// has WriteString method.
var _ io.Writer = (*myWriter)(nil)

type myWriter struct {
	b bytes.Buffer
}

func (w *myWriter) Write(b []byte) (int, error) {
	return w.b.Write(b)
}

func BenchmarkFmtPrintf(b *testing.B) {
	var w myWriter
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(&w, "aaa %v bbb", 18)
	}
}

func BenchmarkStrconv(b *testing.B) {
	var w myWriter
	for i := 0; i < b.N; i++ {
		io.WriteString(&w, "aaa "+strconv.Itoa(18)+" bbb")
	}
}

func BenchmarkStrconvAsBytes(b *testing.B) {
	var w myWriter
	for i := 0; i < b.N; i++ {
		w.Write([]byte("aaa "))
		w.Write([]byte(strconv.Itoa(18)))
		w.Write([]byte(" bbb"))
	}
}

//////// If the writer has WriteString method... ////////

func BenchmarkFmtPrintfForBytesBuffer(b *testing.B) {
	var w bytes.Buffer
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(&w, "aaa %v bbb", 18)
	}
}

func BenchmarkStrconvForBytesBuffer(b *testing.B) {
	var w bytes.Buffer
	for i := 0; i < b.N; i++ {
		io.WriteString(&w, "aaa "+strconv.Itoa(18)+" bbb")
	}
}

func BenchmarkStrconvAsBytesForBytesBuffer(b *testing.B) {
	var w bytes.Buffer
	for i := 0; i < b.N; i++ {
		w.Write([]byte("aaa "))
		w.Write([]byte(strconv.Itoa(18)))
		w.Write([]byte(" bbb"))
	}
}
