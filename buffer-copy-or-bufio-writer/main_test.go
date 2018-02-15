package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var s string

func TestMain(m *testing.M) {
	for _, t := range []string{
		"a",
		"abcdef",
		strings.Repeat("x", 10000),
	} {
		s = t
		code := m.Run()
		if code != 0 {
			os.Exit(code)
		}
	}
}

func BenchmarkBufferCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		buf.WriteString(s)
		io.Copy(ioutil.Discard, &buf)
	}
}

func BenchmarkBufioWriter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bufio.NewWriter(ioutil.Discard)
		buf.WriteString(s)
		buf.Flush()
	}
}

func BenchmarkBufferCopy32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(make([]byte, 0, 32))
		buf.WriteString(s)
		io.CopyBuffer(ioutil.Discard, buf, make([]byte, 32))
	}
}

func BenchmarkBufioWriter16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bufio.NewWriterSize(ioutil.Discard, 16)
		buf.WriteString(s)
		buf.Flush()
	}
}

func BenchmarkBufferFprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		buf.WriteString(s)
		fmt.Fprintf(ioutil.Discard, buf.String())
	}
}
