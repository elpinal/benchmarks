    $ go version
    go version devel +6b6b9f69fd Wed Aug 9 05:33:09 2017 +0000 darwin/amd64

    $ go test -bench . -benchmem -count 4 
    goos: darwin
    goarch: amd64
    pkg: github.com/elpinal/benchmarks/printf-println
    BenchmarkPrintf-4               20000000                62.9 ns/op             0 B/op          0 allocs/op
    BenchmarkPrintf-4               20000000                66.6 ns/op             0 B/op          0 allocs/op
    BenchmarkPrintf-4               20000000                64.0 ns/op             0 B/op          0 allocs/op
    BenchmarkPrintf-4               20000000                64.6 ns/op             0 B/op          0 allocs/op
    BenchmarkPrintln-4              20000000                86.6 ns/op             0 B/op          0 allocs/op
    BenchmarkPrintln-4              20000000                81.3 ns/op             0 B/op          0 allocs/op
    BenchmarkPrintln-4              20000000                85.3 ns/op             0 B/op          0 allocs/op
    BenchmarkPrintln-4              20000000                85.0 ns/op             0 B/op          0 allocs/op
    BenchmarkWrite-4                30000000                41.5 ns/op             8 B/op          1 allocs/op
    BenchmarkWrite-4                30000000                41.6 ns/op             8 B/op          1 allocs/op
    BenchmarkWrite-4                30000000                41.1 ns/op             8 B/op          1 allocs/op
    BenchmarkWrite-4                30000000                40.8 ns/op             8 B/op          1 allocs/op
    BenchmarkPrintfMany-4           10000000               150 ns/op               0 B/op          0 allocs/op
    BenchmarkPrintfMany-4           10000000               150 ns/op               0 B/op          0 allocs/op
    BenchmarkPrintfMany-4           10000000               151 ns/op               0 B/op          0 allocs/op
    BenchmarkPrintfMany-4           10000000               152 ns/op               0 B/op          0 allocs/op
    BenchmarkPrintlnMany-4          10000000               127 ns/op              16 B/op          1 allocs/op
    BenchmarkPrintlnMany-4          10000000               127 ns/op              16 B/op          1 allocs/op
    BenchmarkPrintlnMany-4          10000000               127 ns/op              16 B/op          1 allocs/op
    BenchmarkPrintlnMany-4          10000000               127 ns/op              16 B/op          1 allocs/op
    BenchmarkWriteMany-4            20000000                75.3 ns/op           112 B/op          1 allocs/op
    BenchmarkWriteMany-4            20000000                77.2 ns/op           112 B/op          1 allocs/op
    BenchmarkWriteMany-4            20000000                74.3 ns/op           112 B/op          1 allocs/op
    BenchmarkWriteMany-4            20000000                75.5 ns/op           112 B/op          1 allocs/op
