`go version`:

    go version devel +ecc6a81617 Sat Mar 25 00:35:35 2017 +0000 darwin/amd64

Summary (`benchstat`):

    name                            time/op
    FmtPrintf-4                      142ns ± 1%
    Strconv-4                        176ns ± 1%
    StrconvAsBytes-4                 101ns ± 1%
    FmtPrintfForBytesBuffer-4        139ns ± 2%
    StrconvForBytesBuffer-4          119ns ± 2%
    StrconvAsBytesForBytesBuffer-4  84.8ns ± 2%

    name                            alloc/op
    FmtPrintf-4                      31.0B ± 0%
    Strconv-4                        63.0B ± 0%
    StrconvAsBytes-4                 27.0B ± 0%
    FmtPrintfForBytesBuffer-4        31.0B ± 0%
    StrconvForBytesBuffer-4          47.0B ± 0%
    StrconvAsBytesForBytesBuffer-4   27.0B ± 0%

    name                            allocs/op
    FmtPrintf-4                       0.00
    Strconv-4                         2.00 ± 0%
    StrconvAsBytes-4                  0.00
    FmtPrintfForBytesBuffer-4         0.00
    StrconvForBytesBuffer-4           1.00 ± 0%
    StrconvAsBytesForBytesBuffer-4    0.00

`go test -bench . -benchmem -count 10`:

    goos: darwin
    goarch: amd64
    pkg: github.com/elpinal/benchmarks/write-int
    BenchmarkFmtPrintf-4                      	10000000	       163 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintf-4                      	10000000	       141 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintf-4                      	10000000	       142 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintf-4                      	10000000	       141 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintf-4                      	10000000	       147 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintf-4                      	10000000	       141 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintf-4                      	10000000	       143 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintf-4                      	10000000	       143 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintf-4                      	10000000	       143 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintf-4                      	10000000	       141 ns/op	      31 B/op	       0 allocs/op
    BenchmarkStrconv-4                        	10000000	       176 ns/op	      63 B/op	       2 allocs/op
    BenchmarkStrconv-4                        	10000000	       177 ns/op	      63 B/op	       2 allocs/op
    BenchmarkStrconv-4                        	10000000	       173 ns/op	      63 B/op	       2 allocs/op
    BenchmarkStrconv-4                        	10000000	       177 ns/op	      63 B/op	       2 allocs/op
    BenchmarkStrconv-4                        	10000000	       176 ns/op	      63 B/op	       2 allocs/op
    BenchmarkStrconv-4                        	10000000	       174 ns/op	      63 B/op	       2 allocs/op
    BenchmarkStrconv-4                        	10000000	       176 ns/op	      63 B/op	       2 allocs/op
    BenchmarkStrconv-4                        	10000000	       181 ns/op	      63 B/op	       2 allocs/op
    BenchmarkStrconv-4                        	10000000	       176 ns/op	      63 B/op	       2 allocs/op
    BenchmarkStrconv-4                        	10000000	       191 ns/op	      63 B/op	       2 allocs/op
    BenchmarkStrconvAsBytes-4                 	10000000	       121 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytes-4                 	10000000	       102 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytes-4                 	20000000	       106 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytes-4                 	20000000	       101 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytes-4                 	20000000	       100 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytes-4                 	20000000	        99.3 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytes-4                 	20000000	       101 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytes-4                 	20000000	       101 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytes-4                 	20000000	       100 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytes-4                 	20000000	       100 ns/op	      27 B/op	       0 allocs/op
    BenchmarkFmtPrintfForBytesBuffer-4        	10000000	       139 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintfForBytesBuffer-4        	10000000	       139 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintfForBytesBuffer-4        	10000000	       139 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintfForBytesBuffer-4        	10000000	       138 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintfForBytesBuffer-4        	10000000	       137 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintfForBytesBuffer-4        	10000000	       141 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintfForBytesBuffer-4        	10000000	       140 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintfForBytesBuffer-4        	10000000	       138 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintfForBytesBuffer-4        	10000000	       147 ns/op	      31 B/op	       0 allocs/op
    BenchmarkFmtPrintfForBytesBuffer-4        	10000000	       139 ns/op	      31 B/op	       0 allocs/op
    BenchmarkStrconvForBytesBuffer-4          	10000000	       124 ns/op	      47 B/op	       1 allocs/op
    BenchmarkStrconvForBytesBuffer-4          	10000000	       120 ns/op	      47 B/op	       1 allocs/op
    BenchmarkStrconvForBytesBuffer-4          	10000000	       120 ns/op	      47 B/op	       1 allocs/op
    BenchmarkStrconvForBytesBuffer-4          	10000000	       119 ns/op	      47 B/op	       1 allocs/op
    BenchmarkStrconvForBytesBuffer-4          	10000000	       122 ns/op	      47 B/op	       1 allocs/op
    BenchmarkStrconvForBytesBuffer-4          	10000000	       119 ns/op	      47 B/op	       1 allocs/op
    BenchmarkStrconvForBytesBuffer-4          	10000000	       118 ns/op	      47 B/op	       1 allocs/op
    BenchmarkStrconvForBytesBuffer-4          	10000000	       118 ns/op	      47 B/op	       1 allocs/op
    BenchmarkStrconvForBytesBuffer-4          	10000000	       118 ns/op	      47 B/op	       1 allocs/op
    BenchmarkStrconvForBytesBuffer-4          	10000000	       118 ns/op	      47 B/op	       1 allocs/op
    BenchmarkStrconvAsBytesForBytesBuffer-4   	20000000	        84.1 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytesForBytesBuffer-4   	20000000	        85.3 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytesForBytesBuffer-4   	20000000	        84.9 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytesForBytesBuffer-4   	20000000	        84.9 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytesForBytesBuffer-4   	20000000	        85.0 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytesForBytesBuffer-4   	20000000	        84.3 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytesForBytesBuffer-4   	20000000	        83.7 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytesForBytesBuffer-4   	20000000	        84.5 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytesForBytesBuffer-4   	20000000	        86.9 ns/op	      27 B/op	       0 allocs/op
    BenchmarkStrconvAsBytesForBytesBuffer-4   	20000000	        87.3 ns/op	      27 B/op	       0 allocs/op
    PASS
    ok  	github.com/elpinal/benchmarks/write-int	103.017s
