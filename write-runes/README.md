    BenchmarkIOWriteString-4        20000000                92.6 ns/op             7 B/op          1 allocs/op
    BenchmarkCopyThenWrite-4        20000000                65.6 ns/op             4 B/op          1 allocs/op
    BenchmarkBufioWriteString-4     20000000                61.8 ns/op             3 B/op          0 allocs/op
    BenchmarkBufioWriteRune-4       30000000                34.9 ns/op             2 B/op          0 allocs/op
    BenchmarkFmtFprint-4            10000000               221 ns/op              23 B/op          2 allocs/op
    PASS
    BenchmarkIOWriteString-4        10000000               132 ns/op              22 B/op          1 allocs/op
    BenchmarkCopyThenWrite-4        20000000               113 ns/op              18 B/op          1 allocs/op
    BenchmarkBufioWriteString-4     20000000                97.4 ns/op            14 B/op          0 allocs/op
    BenchmarkBufioWriteRune-4       20000000                71.6 ns/op            14 B/op          0 allocs/op
    BenchmarkFmtFprint-4             5000000               270 ns/op              38 B/op          2 allocs/op
    PASS
    BenchmarkIOWriteString-4         5000000               278 ns/op              65 B/op          1 allocs/op
    BenchmarkCopyThenWrite-4         5000000               251 ns/op              49 B/op          1 allocs/op
    BenchmarkBufioWriteString-4      5000000               239 ns/op              33 B/op          0 allocs/op
    BenchmarkBufioWriteRune-4       10000000               199 ns/op              33 B/op          0 allocs/op
    BenchmarkFmtFprint-4             3000000               409 ns/op             103 B/op          2 allocs/op
    PASS
    BenchmarkIOWriteString-4         2000000               883 ns/op             214 B/op          1 allocs/op
    BenchmarkCopyThenWrite-4         2000000               899 ns/op             278 B/op          2 allocs/op
    BenchmarkBufioWriteString-4      2000000               928 ns/op             281 B/op          1 allocs/op
    BenchmarkBufioWriteRune-4        2000000               748 ns/op             201 B/op          0 allocs/op
    BenchmarkFmtFprint-4             1000000              1013 ns/op             230 B/op          2 allocs/op
    PASS
    BenchmarkIOWriteString-4          500000              3290 ns/op             824 B/op          1 allocs/op
    BenchmarkCopyThenWrite-4          500000              3403 ns/op            1080 B/op          2 allocs/op
    BenchmarkBufioWriteString-4       500000              3409 ns/op            1093 B/op          1 allocs/op
    BenchmarkBufioWriteRune-4         500000              2949 ns/op             805 B/op          0 allocs/op
    BenchmarkFmtFprint-4              500000              3425 ns/op             840 B/op          2 allocs/op
    PASS
