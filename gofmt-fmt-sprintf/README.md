    ∆ go test -bench . -benchmem
    goos: darwin
    goarch: amd64
    pkg: github.com/elpinal/benchmarks/gofmt-fmt-sprintf
    BenchmarkFmtSprintf-4            2000000               883 ns/op             480 B/op         11 allocs/op
    BenchmarkPlus-4                  5000000               313 ns/op             192 B/op          4 allocs/op
    BenchmarkBuffer-4               10000000               218 ns/op             112 B/op          1 allocs/op
    PASS
    ok      github.com/elpinal/benchmarks/gofmt-fmt-sprintf 6.998s
