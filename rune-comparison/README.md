```
$ go version
go version devel +fb1fd6aee6 Wed Nov 1 00:30:39 2017 +0000 darwin/amd64
```

```
$ go test -bench . -benchmem -count 4
goos: darwin
goarch: amd64
pkg: github.com/elpinal/benchmarks/rune-comparison
BenchmarkFor/0-4                        300000000                4.34 ns/op            0 B/op          0 allocs/op
BenchmarkFor/0-4                        300000000                4.38 ns/op            0 B/op          0 allocs/op
BenchmarkFor/0-4                        300000000                4.44 ns/op            0 B/op          0 allocs/op
BenchmarkFor/0-4                        300000000                4.31 ns/op            0 B/op          0 allocs/op
BenchmarkFor/10-4                       300000000                4.78 ns/op            0 B/op          0 allocs/op
BenchmarkFor/10-4                       300000000                4.76 ns/op            0 B/op          0 allocs/op
BenchmarkFor/10-4                       300000000                4.69 ns/op            0 B/op          0 allocs/op
BenchmarkFor/10-4                       300000000                4.69 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100-4                      300000000                4.70 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100-4                      300000000                4.77 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100-4                      300000000                4.98 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100-4                      300000000                4.85 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100:101-4                  300000000                4.34 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100:101-4                  300000000                4.51 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100:101-4                  300000000                4.59 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100:101-4                  300000000                4.39 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100(self)-4                300000000                5.28 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100(self)-4                300000000                4.69 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100(self)-4                300000000                4.81 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100(self)-4                300000000                4.89 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100(diff)-4                300000000                4.76 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100(diff)-4                300000000                4.81 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100(diff)-4                300000000                4.84 ns/op            0 B/op          0 allocs/op
BenchmarkFor/100(diff)-4                300000000                4.76 ns/op            0 B/op          0 allocs/op
BenchmarkString/0-4                     50000000                26.1 ns/op             0 B/op          0 allocs/op
BenchmarkString/0-4                     50000000                24.8 ns/op             0 B/op          0 allocs/op
BenchmarkString/0-4                     100000000               25.8 ns/op             0 B/op          0 allocs/op
BenchmarkString/0-4                     50000000                25.7 ns/op             0 B/op          0 allocs/op
BenchmarkString/10-4                     5000000               246 ns/op               0 B/op          0 allocs/op
BenchmarkString/10-4                     5000000               244 ns/op               0 B/op          0 allocs/op
BenchmarkString/10-4                     5000000               239 ns/op               0 B/op          0 allocs/op
BenchmarkString/10-4                     5000000               234 ns/op               0 B/op          0 allocs/op
BenchmarkString/100-4                    1000000              2307 ns/op             224 B/op          2 allocs/op
BenchmarkString/100-4                    1000000              2371 ns/op             224 B/op          2 allocs/op
BenchmarkString/100-4                    1000000              2303 ns/op             224 B/op          2 allocs/op
BenchmarkString/100-4                    1000000              2300 ns/op             224 B/op          2 allocs/op
BenchmarkString/100:101-4                 500000              2297 ns/op             224 B/op          2 allocs/op
BenchmarkString/100:101-4                 500000              2290 ns/op             224 B/op          2 allocs/op
BenchmarkString/100:101-4                 500000              2263 ns/op             224 B/op          2 allocs/op
BenchmarkString/100:101-4                1000000              2308 ns/op             224 B/op          2 allocs/op
BenchmarkString/100(self)-4              1000000              2291 ns/op             224 B/op          2 allocs/op
BenchmarkString/100(self)-4               500000              2300 ns/op             224 B/op          2 allocs/op
BenchmarkString/100(self)-4              1000000              2277 ns/op             224 B/op          2 allocs/op
BenchmarkString/100(self)-4               500000              2372 ns/op             224 B/op          2 allocs/op
BenchmarkString/100(diff)-4              1000000              2335 ns/op             224 B/op          2 allocs/op
BenchmarkString/100(diff)-4              1000000              2372 ns/op             224 B/op          2 allocs/op
BenchmarkString/100(diff)-4              1000000              2316 ns/op             224 B/op          2 allocs/op
BenchmarkString/100(diff)-4              1000000              2291 ns/op             224 B/op          2 allocs/op
BenchmarkDeepEqual/0-4                  10000000               209 ns/op              64 B/op          2 allocs/op
BenchmarkDeepEqual/0-4                  10000000               205 ns/op              64 B/op          2 allocs/op
BenchmarkDeepEqual/0-4                  10000000               210 ns/op              64 B/op          2 allocs/op
BenchmarkDeepEqual/0-4                  10000000               206 ns/op              64 B/op          2 allocs/op
BenchmarkDeepEqual/10-4                  1000000              1524 ns/op             144 B/op         22 allocs/op
BenchmarkDeepEqual/10-4                  1000000              1546 ns/op             144 B/op         22 allocs/op
BenchmarkDeepEqual/10-4                  1000000              1550 ns/op             144 B/op         22 allocs/op
BenchmarkDeepEqual/10-4                  1000000              1513 ns/op             144 B/op         22 allocs/op
BenchmarkDeepEqual/100-4                  100000             12724 ns/op             864 B/op        202 allocs/op
BenchmarkDeepEqual/100-4                  100000             12798 ns/op             864 B/op        202 allocs/op
BenchmarkDeepEqual/100-4                  100000             12813 ns/op             864 B/op        202 allocs/op
BenchmarkDeepEqual/100-4                  100000             13048 ns/op             864 B/op        202 allocs/op
BenchmarkDeepEqual/100:101-4            10000000               196 ns/op              64 B/op          2 allocs/op
BenchmarkDeepEqual/100:101-4            10000000               207 ns/op              64 B/op          2 allocs/op
BenchmarkDeepEqual/100:101-4            10000000               205 ns/op              64 B/op          2 allocs/op
BenchmarkDeepEqual/100:101-4            10000000               200 ns/op              64 B/op          2 allocs/op
BenchmarkDeepEqual/100(self)-4          10000000               208 ns/op              64 B/op          2 allocs/op
BenchmarkDeepEqual/100(self)-4          10000000               207 ns/op              64 B/op          2 allocs/op
BenchmarkDeepEqual/100(self)-4           5000000               206 ns/op              64 B/op          2 allocs/op
BenchmarkDeepEqual/100(self)-4          10000000               207 ns/op              64 B/op          2 allocs/op
BenchmarkDeepEqual/100(diff)-4           3000000               488 ns/op              80 B/op          6 allocs/op
BenchmarkDeepEqual/100(diff)-4           3000000               484 ns/op              80 B/op          6 allocs/op
BenchmarkDeepEqual/100(diff)-4           3000000               475 ns/op              80 B/op          6 allocs/op
BenchmarkDeepEqual/100(diff)-4           3000000               478 ns/op              80 B/op          6 allocs/op
```
