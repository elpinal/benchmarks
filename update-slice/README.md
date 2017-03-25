    ∆ go version
    go version devel +ecc6a81617 Sat Mar 25 00:35:35 2017 +0000 darwin/amd64
    ∆ go test -bench . -benchmem -count 5
    goos: darwin
    goarch: amd64
    pkg: github.com/elpinal/benchmarks/update-slice
    BenchmarkMake-4                  5000000               244 ns/op             896 B/op          1 allocs/op
    BenchmarkMake-4                  5000000               264 ns/op             896 B/op          1 allocs/op
    BenchmarkMake-4                  5000000               238 ns/op             896 B/op          1 allocs/op
    BenchmarkMake-4                  5000000               240 ns/op             896 B/op          1 allocs/op
    BenchmarkMake-4                  5000000               239 ns/op             896 B/op          1 allocs/op
    BenchmarkAppend-4               50000000                26.3 ns/op             0 B/op          0 allocs/op
    BenchmarkAppend-4               50000000                26.1 ns/op             0 B/op          0 allocs/op
    BenchmarkAppend-4               50000000                26.7 ns/op             0 B/op          0 allocs/op
    BenchmarkAppend-4               50000000                25.9 ns/op             0 B/op          0 allocs/op
    BenchmarkAppend-4               50000000                27.3 ns/op             0 B/op          0 allocs/op
    BenchmarkAppendFullSlice-4       5000000               249 ns/op             896 B/op          1 allocs/op
    BenchmarkAppendFullSlice-4       5000000               255 ns/op             896 B/op          1 allocs/op
    BenchmarkAppendFullSlice-4       5000000               249 ns/op             896 B/op          1 allocs/op
    BenchmarkAppendFullSlice-4       5000000               248 ns/op             896 B/op          1 allocs/op
    BenchmarkAppendFullSlice-4       5000000               246 ns/op             896 B/op          1 allocs/op
    PASS
    ok      github.com/elpinal/benchmarks/update-slice      21.875s
