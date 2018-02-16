```
$ go version
go version devel +6631f22776 Thu Jan 11 02:39:43 2018 +0000 darwin/amd64

$ benchstat old.txt new.txt
name           old time/op    new time/op    delta
PrintfError/0-4      101ns ± 2%     196ns ± 1%   +95.00%  (p=0.002 n=6+6)
PrintfError/1-4      135ns ± 2%     198ns ± 3%   +46.91%  (p=0.002 n=6+6)
PrintfError/2-4      138ns ± 1%     202ns ± 1%   +46.67%  (p=0.002 n=6+6)
PrintfError/3-4      139ns ± 1%     201ns ± 4%   +44.90%  (p=0.002 n=6+6)
PrintfError/4-4      141ns ± 3%     207ns ± 3%   +46.57%  (p=0.002 n=6+6)
PrintfError/5-4      142ns ± 1%     208ns ± 2%   +47.06%  (p=0.002 n=6+6)
PrintfError/6-4      142ns ± 1%     209ns ± 1%   +47.07%  (p=0.002 n=6+6)
PrintfError/7-4      151ns ± 3%     212ns ± 3%   +40.20%  (p=0.002 n=6+6)
PrintfError/8-4      158ns ± 1%     221ns ± 1%   +40.27%  (p=0.002 n=6+6)
PrintfError/9-4      171ns ± 1%     246ns ± 3%   +43.48%  (p=0.002 n=6+6)
PrintfError/10-4     200ns ± 2%     272ns ± 2%   +36.03%  (p=0.002 n=6+6)
PrintfError/11-4     274ns ± 1%     340ns ± 2%   +24.27%  (p=0.002 n=6+6)

name           old alloc/op   new alloc/op   delta
PrintfError/0-4      0.00B          0.00B           ~     (all equal)
PrintfError/1-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/2-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/3-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/4-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/5-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/6-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/7-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/8-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/9-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/10-4     16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/11-4     16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)

name           old allocs/op  new allocs/op  delta
PrintfError/0-4       0.00           0.00           ~     (all equal)
PrintfError/1-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/2-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/3-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/4-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/5-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/6-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/7-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/8-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/9-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/10-4      1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/11-4      1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
```

```
$ go test -count=6 -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/elpinal/benchmarks/printf-error
BenchmarkExplicit/0-4   20000000               101 ns/op               0 B/op          0 allocs/op
BenchmarkExplicit/0-4   20000000               102 ns/op               0 B/op          0 allocs/op
BenchmarkExplicit/0-4   20000000               100 ns/op               0 B/op          0 allocs/op
BenchmarkExplicit/0-4   20000000               100 ns/op               0 B/op          0 allocs/op
BenchmarkExplicit/0-4   20000000                99.1 ns/op             0 B/op          0 allocs/op
BenchmarkExplicit/0-4   20000000               102 ns/op               0 B/op          0 allocs/op
BenchmarkExplicit/1-4   10000000               134 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/1-4   10000000               133 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/1-4   10000000               138 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/1-4   10000000               135 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/1-4   10000000               133 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/1-4   10000000               135 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/2-4   10000000               138 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/2-4   10000000               136 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/2-4   10000000               137 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/2-4   10000000               139 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/2-4   10000000               136 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/2-4   10000000               139 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/3-4   10000000               137 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/3-4   10000000               139 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/3-4   10000000               140 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/3-4   10000000               138 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/3-4   10000000               140 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/3-4   10000000               139 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/4-4   10000000               139 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/4-4   10000000               145 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/4-4   10000000               138 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/4-4   10000000               137 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/4-4   10000000               144 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/4-4   10000000               143 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/5-4   10000000               141 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/5-4   10000000               141 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/5-4   10000000               143 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/5-4   10000000               142 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/5-4   10000000               143 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/5-4   10000000               140 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/6-4   10000000               141 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/6-4   10000000               142 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/6-4   10000000               143 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/6-4   10000000               141 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/6-4   10000000               141 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/6-4   10000000               144 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/7-4   10000000               149 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/7-4   10000000               151 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/7-4   10000000               156 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/7-4   10000000               149 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/7-4   10000000               149 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/7-4   10000000               154 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/8-4   10000000               157 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/8-4   10000000               159 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/8-4   10000000               159 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/8-4   10000000               158 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/8-4   10000000               156 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/8-4   10000000               157 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/9-4   10000000               172 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/9-4   10000000               173 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/9-4   10000000               172 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/9-4   10000000               171 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/9-4   10000000               171 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/9-4   10000000               169 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/10-4          10000000               204 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/10-4          10000000               198 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/10-4          10000000               197 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/10-4          10000000               203 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/10-4          10000000               200 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/10-4          10000000               197 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/11-4           5000000               276 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/11-4           5000000               272 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/11-4           5000000               275 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/11-4           5000000               275 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/11-4           5000000               271 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/11-4           5000000               275 ns/op              16 B/op          1 allocs/op
BenchmarkImplicit/0-4           10000000               195 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/0-4           10000000               198 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/0-4           10000000               197 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/0-4           10000000               199 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/0-4           10000000               194 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/0-4           10000000               195 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/1-4           10000000               194 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/1-4           10000000               203 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/1-4           10000000               198 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/1-4           10000000               195 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/1-4           10000000               199 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/1-4           10000000               198 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/2-4           10000000               203 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/2-4           10000000               199 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/2-4           10000000               204 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/2-4           10000000               200 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/2-4           10000000               201 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/2-4           10000000               203 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/3-4           10000000               198 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/3-4           10000000               198 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/3-4           10000000               198 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/3-4           10000000               205 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/3-4           10000000               199 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/3-4           10000000               209 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/4-4           10000000               201 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/4-4           10000000               211 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/4-4           10000000               200 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/4-4           10000000               210 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/4-4           10000000               207 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/4-4           10000000               211 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/5-4           10000000               210 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/5-4           10000000               211 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/5-4           10000000               210 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/5-4           10000000               205 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/5-4           10000000               207 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/5-4            5000000               207 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/6-4           10000000               211 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/6-4           10000000               207 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/6-4           10000000               209 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/6-4           10000000               207 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/6-4           10000000               209 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/6-4           10000000               210 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/7-4           10000000               210 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/7-4           10000000               213 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/7-4            5000000               218 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/7-4           10000000               213 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/7-4           10000000               210 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/7-4           10000000               209 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/8-4           10000000               220 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/8-4           10000000               224 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/8-4            5000000               220 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/8-4           10000000               223 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/8-4           10000000               221 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/8-4           10000000               219 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/9-4            5000000               250 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/9-4            5000000               238 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/9-4            5000000               239 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/9-4            5000000               245 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/9-4            5000000               250 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/9-4            5000000               253 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/10-4           5000000               275 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/10-4           5000000               271 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/10-4           5000000               276 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/10-4           5000000               269 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/10-4           5000000               274 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/10-4           5000000               266 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/11-4           5000000               348 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/11-4           5000000               337 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/11-4           5000000               338 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/11-4           5000000               338 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/11-4           5000000               346 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/11-4           5000000               336 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/elpinal/benchmarks/printf-error      276.084s
```


```
$ go version
go version devel +41d3d153eb Fri Feb 16 03:31:46 2018 +0000 darwin/amd64

name           old time/op    new time/op    delta
PrintfError/0-4     98.9ns ± 1%   198.5ns ± 3%  +100.78%  (p=0.002 n=6+6)
PrintfError/1-4      136ns ± 4%     198ns ± 2%   +45.96%  (p=0.002 n=6+6)
PrintfError/2-4      136ns ± 2%     198ns ± 1%   +45.01%  (p=0.004 n=5+6)
PrintfError/3-4      136ns ± 2%     199ns ± 2%   +46.14%  (p=0.002 n=6+6)
PrintfError/4-4      136ns ± 1%     202ns ± 1%   +47.80%  (p=0.002 n=6+6)
PrintfError/5-4      140ns ± 1%     205ns ± 3%   +45.77%  (p=0.004 n=5+6)
PrintfError/6-4      144ns ± 2%     210ns ± 2%   +45.89%  (p=0.002 n=6+6)
PrintfError/7-4      150ns ± 4%     213ns ± 1%   +41.82%  (p=0.004 n=6+5)
PrintfError/8-4      156ns ± 2%     222ns ± 2%   +41.64%  (p=0.002 n=6+6)
PrintfError/9-4      171ns ± 1%     238ns ± 2%   +39.38%  (p=0.002 n=6+6)
PrintfError/10-4     203ns ± 1%     272ns ± 2%   +34.46%  (p=0.002 n=6+6)
PrintfError/11-4     278ns ± 3%     339ns ± 2%   +21.74%  (p=0.002 n=6+6)

name           old alloc/op   new alloc/op   delta
PrintfError/0-4      0.00B          0.00B           ~     (all equal)
PrintfError/1-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/2-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/3-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/4-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/5-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/6-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/7-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/8-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/9-4      16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/10-4     16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)
PrintfError/11-4     16.0B ± 0%      0.0B       -100.00%  (p=0.002 n=6+6)

name           old allocs/op  new allocs/op  delta
PrintfError/0-4       0.00           0.00           ~     (all equal)
PrintfError/1-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/2-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/3-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/4-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/5-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/6-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/7-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/8-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/9-4       1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/10-4      1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
PrintfError/11-4      1.00 ± 0%      0.00       -100.00%  (p=0.002 n=6+6)
```

```
$ go test -count=6 -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/elpinal/benchmarks/printf-error
BenchmarkExplicit/0-4   20000000                99.5 ns/op             0 B/op          0 allocs/op
BenchmarkExplicit/0-4   20000000                99.1 ns/op             0 B/op          0 allocs/op
BenchmarkExplicit/0-4   20000000                97.4 ns/op             0 B/op          0 allocs/op
BenchmarkExplicit/0-4   20000000                98.7 ns/op             0 B/op          0 allocs/op
BenchmarkExplicit/0-4   20000000                98.5 ns/op             0 B/op          0 allocs/op
BenchmarkExplicit/0-4   20000000               100 ns/op               0 B/op          0 allocs/op
BenchmarkExplicit/1-4   10000000               134 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/1-4   10000000               138 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/1-4   10000000               141 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/1-4   10000000               133 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/1-4   10000000               135 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/1-4   10000000               135 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/2-4   10000000               134 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/2-4   10000000               148 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/2-4   10000000               137 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/2-4   10000000               136 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/2-4   10000000               136 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/2-4   10000000               138 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/3-4   10000000               135 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/3-4   10000000               136 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/3-4   10000000               137 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/3-4   10000000               137 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/3-4   10000000               138 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/3-4   10000000               134 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/4-4   10000000               135 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/4-4   10000000               137 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/4-4   10000000               138 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/4-4   10000000               136 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/4-4   10000000               137 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/4-4   10000000               135 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/5-4   10000000               139 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/5-4   10000000               141 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/5-4   10000000               141 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/5-4   10000000               144 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/5-4   10000000               141 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/5-4   10000000               140 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/6-4   10000000               142 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/6-4   10000000               143 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/6-4   10000000               146 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/6-4   10000000               145 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/6-4   10000000               145 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/6-4   10000000               142 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/7-4   10000000               151 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/7-4   10000000               151 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/7-4   10000000               146 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/7-4   10000000               147 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/7-4   10000000               156 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/7-4   10000000               151 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/8-4   10000000               156 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/8-4   10000000               155 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/8-4   10000000               157 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/8-4   10000000               156 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/8-4   10000000               159 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/8-4   10000000               156 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/9-4   10000000               169 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/9-4   10000000               172 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/9-4   10000000               173 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/9-4   10000000               170 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/9-4   10000000               172 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/9-4   10000000               170 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/10-4          10000000               204 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/10-4          10000000               203 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/10-4          10000000               205 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/10-4          10000000               202 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/10-4          10000000               201 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/10-4          10000000               201 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/11-4           5000000               275 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/11-4           5000000               280 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/11-4           5000000               280 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/11-4           5000000               282 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/11-4           5000000               282 ns/op              16 B/op          1 allocs/op
BenchmarkExplicit/11-4           5000000               271 ns/op              16 B/op          1 allocs/op
BenchmarkImplicit/0-4           10000000               193 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/0-4           10000000               200 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/0-4           10000000               200 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/0-4           10000000               195 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/0-4           10000000               201 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/0-4           10000000               202 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/1-4           10000000               196 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/1-4           10000000               199 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/1-4           10000000               202 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/1-4           10000000               197 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/1-4           10000000               199 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/1-4           10000000               198 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/2-4           10000000               200 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/2-4           10000000               198 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/2-4           10000000               197 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/2-4           10000000               198 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/2-4           10000000               195 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/2-4           10000000               197 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/3-4           10000000               195 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/3-4           10000000               201 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/3-4           10000000               195 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/3-4           10000000               202 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/3-4           10000000               201 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/3-4           10000000               200 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/4-4           10000000               201 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/4-4           10000000               203 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/4-4           10000000               204 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/4-4           10000000               201 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/4-4           10000000               200 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/4-4           10000000               200 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/5-4           10000000               205 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/5-4           10000000               203 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/5-4           10000000               210 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/5-4           10000000               199 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/5-4           10000000               207 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/5-4           10000000               204 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/6-4           10000000               211 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/6-4           10000000               213 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/6-4           10000000               207 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/6-4           10000000               211 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/6-4           10000000               208 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/6-4           10000000               209 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/7-4           10000000               213 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/7-4           10000000               212 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/7-4           10000000               214 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/7-4           10000000               217 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/7-4           10000000               213 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/7-4           10000000               214 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/8-4           10000000               220 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/8-4           10000000               218 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/8-4           10000000               225 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/8-4           10000000               223 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/8-4            5000000               222 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/8-4           10000000               222 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/9-4            5000000               238 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/9-4            5000000               236 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/9-4            5000000               238 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/9-4            5000000               242 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/9-4            5000000               235 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/9-4            5000000               241 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/10-4           5000000               278 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/10-4           5000000               274 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/10-4           5000000               271 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/10-4           5000000               269 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/10-4           5000000               274 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/10-4           5000000               269 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/11-4           5000000               339 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/11-4           5000000               331 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/11-4           5000000               343 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/11-4           5000000               336 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/11-4           5000000               343 ns/op               0 B/op          0 allocs/op
BenchmarkImplicit/11-4           5000000               341 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/elpinal/benchmarks/printf-error      276.922s
```
