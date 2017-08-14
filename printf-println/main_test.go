package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func BenchmarkPrintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "aaa\n")
	}
}

func BenchmarkPrintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintln(ioutil.Discard, "aaa")
	}
}

func BenchmarkWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioutil.Discard.Write([]byte("aaa\n"))
	}
}

var s1 = strings.Repeat("a", 100)
var s2 = s1 + "\n"

func BenchmarkPrintfMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, s2)
	}
}

func BenchmarkPrintlnMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintln(ioutil.Discard, s1)
	}
}

func BenchmarkWriteMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioutil.Discard.Write([]byte(s2))
	}
}
