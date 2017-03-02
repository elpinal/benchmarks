    BenchmarkIOWriteString-4        20000000                89.4 ns/op             7 B/op          1 allocs/op
    BenchmarkCopyThenWrite-4        20000000                63.2 ns/op             4 B/op          1 allocs/op
    BenchmarkBufioWriteString-4     20000000                62.3 ns/op             3 B/op          0 allocs/op
    BenchmarkBufioWriteRune-4       50000000                35.2 ns/op             2 B/op          0 allocs/op
    PASS
    BenchmarkIOWriteString-4        10000000               121 ns/op              22 B/op          1 allocs/op
    BenchmarkCopyThenWrite-4        20000000               107 ns/op              18 B/op          1 allocs/op
    BenchmarkBufioWriteString-4     20000000                95.4 ns/op            14 B/op          0 allocs/op
    BenchmarkBufioWriteRune-4       20000000                65.2 ns/op            14 B/op          0 allocs/op
    PASS
    BenchmarkIOWriteString-4         5000000               276 ns/op              65 B/op          1 allocs/op
    BenchmarkCopyThenWrite-4         5000000               248 ns/op              49 B/op          1 allocs/op
    BenchmarkBufioWriteString-4      5000000               234 ns/op              33 B/op          0 allocs/op
    BenchmarkBufioWriteRune-4       10000000               195 ns/op              33 B/op          0 allocs/op
    PASS
    BenchmarkIOWriteString-4         2000000               843 ns/op             214 B/op          1 allocs/op
    BenchmarkCopyThenWrite-4         2000000               867 ns/op             278 B/op          2 allocs/op
    BenchmarkBufioWriteString-4      2000000               899 ns/op             281 B/op          1 allocs/op
    BenchmarkBufioWriteRune-4        2000000               728 ns/op             201 B/op          0 allocs/op
    PASS
    BenchmarkIOWriteString-4          500000              3011 ns/op             824 B/op          1 allocs/op
    BenchmarkCopyThenWrite-4          500000              3194 ns/op            1080 B/op          2 allocs/op
    BenchmarkBufioWriteString-4       500000              3334 ns/op            1093 B/op          1 allocs/op
    BenchmarkBufioWriteRune-4         500000              2859 ns/op             805 B/op          0 allocs/op
    PASS
