    ∆ go version
    go version devel +ecc6a81617 Sat Mar 25 00:35:35 2017 +0000 darwin/amd64
    ∆ go test -bench . -benchmem -count 5
    goos: darwin
    goarch: amd64
    pkg: github.com/elpinal/benchmarks/array-vs-slice-for-simple-use
    BenchmarkArray-4        2000000000               0.46 ns/op            0 B/op          0 allocs/op
    BenchmarkArray-4        2000000000               0.47 ns/op            0 B/op          0 allocs/op
    BenchmarkArray-4        2000000000               0.46 ns/op            0 B/op          0 allocs/op
    BenchmarkArray-4        2000000000               0.46 ns/op            0 B/op          0 allocs/op
    BenchmarkArray-4        2000000000               0.46 ns/op            0 B/op          0 allocs/op
    BenchmarkSlice-4        50000000                35.4 ns/op             0 B/op          0 allocs/op
    BenchmarkSlice-4        50000000                35.2 ns/op             0 B/op          0 allocs/op
    BenchmarkSlice-4        50000000                35.7 ns/op             0 B/op          0 allocs/op
    BenchmarkSlice-4        50000000                36.9 ns/op             0 B/op          0 allocs/op
    BenchmarkSlice-4        50000000                35.5 ns/op             0 B/op          0 allocs/op
    PASS
    ok      github.com/elpinal/benchmarks/array-vs-slice-for-simple-use     13.995s
