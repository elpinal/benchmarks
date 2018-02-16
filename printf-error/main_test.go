package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

var errs = []error{
	errors.New(""),
	errors.New("a"),
	errors.New("this is an error"),
	errors.New(strings.Repeat("a", 32)),
	errors.New(strings.Repeat("a", 64)),
	errors.New(strings.Repeat("a", 128)),
	errors.New(strings.Repeat("a", 256)),
	errors.New(strings.Repeat("a", 512)),
	errors.New(strings.Repeat("a", 1024)),
	errors.New(strings.Repeat("a", 2048)),
	errors.New(strings.Repeat("a", 4096)),
	errors.New(strings.Repeat("a", 8192)),
}

func BenchmarkExplicit(b *testing.B) {
	for i, err := range errs {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) { benchmarkExplicit(b, err) })
	}
}

func benchmarkExplicit(b *testing.B, err error) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%s\n", err.Error())
	}
}

func BenchmarkImplicit(b *testing.B) {
	for i, err := range errs {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) { benchmarkImplicit(b, err) })
	}
}

func benchmarkImplicit(b *testing.B, err error) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(ioutil.Discard, "%s\n", err)
	}
}
