```
$ go version
go version devel +6631f22776 Thu Jan 11 02:39:43 2018 +0000 darwin/amd64

$ go test -count=3 -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/elpinal/benchmarks/buffer-copy-or-bufio-writer
BenchmarkBufferCopy-4           10000000               113 ns/op             112 B/op          1 allocs/op
BenchmarkBufferCopy-4           10000000               112 ns/op             112 B/op          1 allocs/op
BenchmarkBufferCopy-4           10000000               113 ns/op             112 B/op          1 allocs/op
BenchmarkBufioWriter-4           2000000               957 ns/op            4096 B/op          1 allocs/op
BenchmarkBufioWriter-4           2000000               960 ns/op            4096 B/op          1 allocs/op
BenchmarkBufioWriter-4           2000000               959 ns/op            4096 B/op          1 allocs/op
BenchmarkBufferCopy32-4         10000000               184 ns/op             176 B/op          3 allocs/op
BenchmarkBufferCopy32-4         10000000               185 ns/op             176 B/op          3 allocs/op
BenchmarkBufferCopy32-4         10000000               181 ns/op             176 B/op          3 allocs/op
BenchmarkBufioWriter16-4        20000000                55.4 ns/op            16 B/op          1 allocs/op
BenchmarkBufioWriter16-4        20000000                56.1 ns/op            16 B/op          1 allocs/op
BenchmarkBufioWriter16-4        20000000                53.9 ns/op            16 B/op          1 allocs/op
BenchmarkBufferFprintf-4        10000000               151 ns/op             112 B/op          1 allocs/op
BenchmarkBufferFprintf-4        10000000               151 ns/op             112 B/op          1 allocs/op
BenchmarkBufferFprintf-4        10000000               151 ns/op             112 B/op          1 allocs/op
PASS
BenchmarkBufferCopy-4           10000000               114 ns/op             112 B/op          1 allocs/op
BenchmarkBufferCopy-4           10000000               117 ns/op             112 B/op          1 allocs/op
BenchmarkBufferCopy-4           10000000               114 ns/op             112 B/op          1 allocs/op
BenchmarkBufioWriter-4           2000000               972 ns/op            4096 B/op          1 allocs/op
BenchmarkBufioWriter-4           1000000              1008 ns/op            4096 B/op          1 allocs/op
BenchmarkBufioWriter-4           1000000              1728 ns/op            4096 B/op          1 allocs/op
BenchmarkBufferCopy32-4          5000000               404 ns/op             176 B/op          3 allocs/op
BenchmarkBufferCopy32-4          3000000               436 ns/op             176 B/op          3 allocs/op
BenchmarkBufferCopy32-4          3000000               365 ns/op             176 B/op          3 allocs/op
BenchmarkBufioWriter16-4        20000000               120 ns/op              16 B/op          1 allocs/op
BenchmarkBufioWriter16-4        20000000                74.3 ns/op            16 B/op          1 allocs/op
BenchmarkBufioWriter16-4        20000000                59.3 ns/op            16 B/op          1 allocs/op
BenchmarkBufferFprintf-4        10000000               165 ns/op             112 B/op          1 allocs/op
BenchmarkBufferFprintf-4        10000000               189 ns/op             112 B/op          1 allocs/op
BenchmarkBufferFprintf-4        10000000               214 ns/op             112 B/op          1 allocs/op
PASS
BenchmarkBufferCopy-4             500000              2587 ns/op           10352 B/op          2 allocs/op
BenchmarkBufferCopy-4             500000              2575 ns/op           10352 B/op          2 allocs/op
BenchmarkBufferCopy-4             500000              2618 ns/op           10352 B/op          2 allocs/op
BenchmarkBufioWriter-4           1000000              1193 ns/op            4096 B/op          1 allocs/op
BenchmarkBufioWriter-4           1000000              1204 ns/op            4096 B/op          1 allocs/op
BenchmarkBufioWriter-4           1000000              1240 ns/op            4096 B/op          1 allocs/op
BenchmarkBufferCopy32-4           500000              2674 ns/op           10416 B/op          4 allocs/op
BenchmarkBufferCopy32-4           500000              2691 ns/op           10416 B/op          4 allocs/op
BenchmarkBufferCopy32-4           500000              2655 ns/op           10416 B/op          4 allocs/op
BenchmarkBufioWriter16-4          200000             10786 ns/op              16 B/op          1 allocs/op
BenchmarkBufioWriter16-4          100000             10090 ns/op              16 B/op          1 allocs/op
BenchmarkBufioWriter16-4          200000             10279 ns/op              16 B/op          1 allocs/op
BenchmarkBufferFprintf-4          100000             11786 ns/op           20666 B/op          3 allocs/op
BenchmarkBufferFprintf-4          100000             11721 ns/op           20665 B/op          3 allocs/op
BenchmarkBufferFprintf-4          100000             12626 ns/op           20666 B/op          3 allocs/op
PASS
ok      github.com/elpinal/benchmarks/buffer-copy-or-bufio-writer       75.735s
```
