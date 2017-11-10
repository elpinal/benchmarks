```
$ go version
go version devel +fb1fd6aee6 Wed Nov 1 00:30:39 2017 +0000 darwin/amd64
```

```
$ go test -bench . -benchmem -count 4
goos: darwin
goarch: amd64
pkg: github.com/elpinal/benchmarks/slice-append
BenchmarkAppend-4               10000000               118 ns/op             192 B/op          2 allocs/op
BenchmarkAppend-4               10000000               127 ns/op             192 B/op          2 allocs/op
BenchmarkAppend-4               20000000               125 ns/op             192 B/op          2 allocs/op
BenchmarkAppend-4               10000000               159 ns/op             192 B/op          2 allocs/op
BenchmarkAppendFirst-4          10000000               166 ns/op             152 B/op          3 allocs/op
BenchmarkAppendFirst-4          10000000               189 ns/op             152 B/op          3 allocs/op
BenchmarkAppendFirst-4          10000000               189 ns/op             152 B/op          3 allocs/op
BenchmarkAppendFirst-4          10000000               178 ns/op             152 B/op          3 allocs/op
