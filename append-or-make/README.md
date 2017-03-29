`go version`:

    go version devel +9f232c1786 Wed Mar 29 00:26:20 2017 +0000 darwin/amd64

Summary (`benchstat`, Append -> Make):

    name                old time/op    new time/op    delta
    Append/1/1-4          81.3ns ± 4%    72.0ns ± 2%  -11.40%  (p=0.029 n=4+4)
    Append/1/2-4          88.6ns ± 4%    76.6ns ± 0%  -13.59%  (p=0.029 n=4+4)
    Append/1/4-4          87.7ns ± 1%    82.5ns ± 0%   -5.93%  (p=0.029 n=4+4)
    Append/1/8-4          95.4ns ± 0%    93.1ns ± 1%   -2.38%  (p=0.029 n=4+4)
    Append/1/16-4          115ns ± 7%     118ns ± 0%     ~     (p=0.343 n=4+4)
    Append/1/32-4          144ns ± 1%     175ns ± 3%  +21.32%  (p=0.029 n=4+4)
    Append/1/64-4          220ns ±14%     278ns ± 0%  +25.99%  (p=0.029 n=4+4)
    Append/1/128-4         353ns ± 2%     474ns ± 0%  +34.11%  (p=0.029 n=4+4)
    Append/1/256-4         529ns ± 1%     814ns ± 1%  +53.85%  (p=0.029 n=4+4)
    Append/1/512-4        1.04µs ± 1%    1.59µs ± 1%  +53.33%  (p=0.029 n=4+4)
    Append/1/1024-4       2.01µs ± 0%    3.01µs ± 0%  +49.84%  (p=0.029 n=4+4)
    Append/2/1-4          94.8ns ± 1%    85.1ns ± 1%  -10.26%  (p=0.029 n=4+4)
    Append/2/2-4          94.9ns ± 1%    86.8ns ± 1%   -8.46%  (p=0.029 n=4+4)
    Append/2/4-4          97.3ns ± 1%    91.9ns ± 1%   -5.47%  (p=0.029 n=4+4)
    Append/2/8-4           105ns ± 1%     102ns ± 0%   -2.61%  (p=0.029 n=4+4)
    Append/2/16-4          122ns ± 2%     129ns ± 1%   +5.94%  (p=0.029 n=4+4)
    Append/2/32-4          157ns ± 1%     184ns ± 1%  +17.22%  (p=0.029 n=4+4)
    Append/2/64-4          220ns ± 0%     290ns ± 1%  +31.89%  (p=0.029 n=4+4)
    Append/2/128-4         362ns ± 1%     487ns ± 0%  +34.46%  (p=0.029 n=4+4)
    Append/2/256-4         546ns ± 1%     820ns ± 1%  +50.05%  (p=0.029 n=4+4)
    Append/2/512-4        1.05µs ± 1%    1.60µs ± 0%  +52.12%  (p=0.029 n=4+4)
    Append/2/1024-4       2.03µs ± 1%    3.04µs ± 3%  +49.78%  (p=0.029 n=4+4)
    Append/4/1-4           104ns ± 0%      93ns ± 1%  -10.19%  (p=0.029 n=4+4)
    Append/4/2-4           105ns ± 2%      95ns ± 0%   -9.98%  (p=0.029 n=4+4)
    Append/4/4-4           104ns ± 0%     100ns ± 1%   -4.09%  (p=0.029 n=4+4)
    Append/4/8-4           112ns ± 0%     112ns ± 1%     ~     (p=1.000 n=4+4)
    Append/4/16-4          128ns ± 0%     138ns ± 1%   +8.04%  (p=0.029 n=4+4)
    Append/4/32-4          158ns ± 1%     190ns ± 0%  +19.91%  (p=0.029 n=4+4)
    Append/4/64-4          223ns ± 1%     295ns ± 1%  +32.25%  (p=0.029 n=4+4)
    Append/4/128-4         365ns ± 3%     489ns ± 0%  +33.97%  (p=0.029 n=4+4)
    Append/4/256-4         554ns ± 2%     830ns ± 1%  +49.95%  (p=0.029 n=4+4)
    Append/4/512-4        1.05µs ± 1%    1.63µs ± 4%  +54.81%  (p=0.029 n=4+4)
    Append/4/1024-4       2.04µs ± 1%    3.00µs ± 1%  +46.91%  (p=0.029 n=4+4)
    Append/8/1-4           128ns ± 0%     109ns ± 2%  -14.51%  (p=0.029 n=4+4)
    Append/8/2-4           127ns ± 1%     111ns ± 1%  -12.77%  (p=0.029 n=4+4)
    Append/8/4-4           128ns ± 1%     116ns ± 1%   -9.00%  (p=0.029 n=4+4)
    Append/8/8-4           128ns ± 1%     128ns ± 0%     ~     (p=1.000 n=4+4)
    Append/8/16-4          143ns ± 1%     151ns ± 0%   +5.58%  (p=0.029 n=4+4)
    Append/8/32-4          177ns ± 0%     207ns ± 1%  +16.81%  (p=0.029 n=4+4)
    Append/8/64-4          231ns ± 0%     307ns ± 1%  +32.65%  (p=0.029 n=4+4)
    Append/8/128-4         371ns ± 0%     502ns ± 0%  +35.18%  (p=0.029 n=4+4)
    Append/8/256-4         565ns ± 1%     844ns ± 0%  +49.45%  (p=0.029 n=4+4)
    Append/8/512-4        1.07µs ± 1%    1.66µs ± 6%  +55.66%  (p=0.029 n=4+4)
    Append/8/1024-4       2.07µs ± 3%    3.00µs ± 1%  +45.21%  (p=0.029 n=4+4)
    Append/16/1-4          174ns ± 0%     144ns ± 1%  -17.77%  (p=0.029 n=4+4)
    Append/16/2-4          175ns ± 0%     144ns ± 1%  -17.45%  (p=0.029 n=4+4)
    Append/16/4-4          175ns ± 2%     151ns ± 1%  -13.84%  (p=0.029 n=4+4)
    Append/16/8-4          177ns ± 2%     163ns ± 1%   -8.05%  (p=0.029 n=4+4)
    Append/16/16-4         182ns ± 4%     186ns ± 1%     ~     (p=0.686 n=4+4)
    Append/16/32-4         208ns ± 0%     236ns ± 0%  +13.43%  (p=0.029 n=4+4)
    Append/16/64-4         277ns ± 1%     349ns ± 2%  +26.11%  (p=0.029 n=4+4)
    Append/16/128-4        396ns ± 1%     530ns ± 2%  +33.84%  (p=0.029 n=4+4)
    Append/16/256-4        598ns ± 7%    1181ns ±64%  +97.45%  (p=0.029 n=4+4)
    Append/16/512-4       1.12µs ± 4%    1.78µs ± 4%  +59.55%  (p=0.029 n=4+4)
    Append/16/1024-4      2.07µs ± 1%    3.11µs ± 2%  +50.46%  (p=0.029 n=4+4)
    Append/32/1-4          266ns ± 0%     214ns ± 0%  -19.61%  (p=0.029 n=4+4)
    Append/32/2-4          271ns ± 1%     215ns ± 0%  -20.68%  (p=0.029 n=4+4)
    Append/32/4-4          270ns ± 1%     216ns ± 1%  -19.67%  (p=0.029 n=4+4)
    Append/32/8-4          273ns ± 0%     230ns ± 1%  -15.92%  (p=0.029 n=4+4)
    Append/32/16-4         274ns ± 1%     253ns ± 0%   -7.66%  (p=0.029 n=4+4)
    Append/32/32-4         272ns ± 0%     305ns ± 2%  +12.14%  (p=0.029 n=4+4)
    Append/32/64-4         352ns ± 2%     423ns ± 1%  +20.18%  (p=0.029 n=4+4)
    Append/32/128-4        475ns ± 1%     638ns ± 1%  +34.32%  (p=0.029 n=4+4)
    Append/32/256-4        631ns ± 2%     968ns ± 1%  +53.45%  (p=0.029 n=4+4)
    Append/32/512-4       1.18µs ± 1%    1.81µs ± 1%  +53.27%  (p=0.029 n=4+4)
    Append/32/1024-4      2.17µs ± 3%    3.27µs ± 1%  +50.55%  (p=0.029 n=4+4)
    Append/64/1-4          476ns ± 1%     362ns ± 0%  -23.97%  (p=0.029 n=4+4)
    Append/64/2-4          478ns ± 1%     371ns ± 7%  -22.40%  (p=0.029 n=4+4)
    Append/64/4-4          474ns ± 1%     367ns ± 0%  -22.59%  (p=0.029 n=4+4)
    Append/64/8-4          477ns ± 1%     369ns ± 2%  -22.55%  (p=0.029 n=4+4)
    Append/64/16-4         490ns ± 3%     400ns ± 1%  -18.34%  (p=0.029 n=4+4)
    Append/64/32-4         484ns ± 1%     461ns ± 0%   -4.65%  (p=0.029 n=4+4)
    Append/64/64-4         480ns ± 2%     541ns ± 1%  +12.66%  (p=0.029 n=4+4)
    Append/64/128-4        622ns ± 2%     752ns ± 1%  +20.82%  (p=0.029 n=4+4)
    Append/64/256-4        883ns ± 1%    1164ns ± 0%  +31.81%  (p=0.029 n=4+4)
    Append/64/512-4       1.28µs ± 4%    1.81µs ± 4%  +41.70%  (p=0.029 n=4+4)
    Append/64/1024-4      2.20µs ± 1%    3.17µs ± 1%  +44.03%  (p=0.029 n=4+4)
    Append/128/1-4         879ns ± 0%     650ns ± 1%  -26.02%  (p=0.029 n=4+4)
    Append/128/2-4         894ns ± 3%     648ns ± 0%  -27.50%  (p=0.029 n=4+4)
    Append/128/4-4         856ns ± 1%     652ns ± 1%  -23.88%  (p=0.029 n=4+4)
    Append/128/8-4         866ns ± 2%     658ns ± 1%  -23.99%  (p=0.029 n=4+4)
    Append/128/16-4        880ns ± 1%     671ns ± 2%  -23.80%  (p=0.029 n=4+4)
    Append/128/32-4        904ns ± 1%     739ns ± 0%  -18.27%  (p=0.029 n=4+4)
    Append/128/64-4        868ns ± 1%     858ns ± 3%     ~     (p=0.286 n=4+4)
    Append/128/128-4       868ns ± 1%    1047ns ± 5%  +20.66%  (p=0.029 n=4+4)
    Append/128/256-4      1.03µs ± 1%    1.31µs ± 1%  +26.50%  (p=0.029 n=4+4)
    Append/128/512-4      1.52µs ± 1%    2.12µs ± 3%  +39.28%  (p=0.029 n=4+4)
    Append/128/1024-4     2.44µs ± 3%    3.41µs ± 3%  +39.88%  (p=0.029 n=4+4)
    Append/256/1-4        1.69µs ± 2%    1.19µs ± 0%  -29.68%  (p=0.029 n=4+4)
    Append/256/2-4        1.68µs ± 1%    1.22µs ± 2%  -27.58%  (p=0.029 n=4+4)
    Append/256/4-4        1.65µs ± 1%    1.20µs ± 1%  -27.25%  (p=0.029 n=4+4)
    Append/256/8-4        1.69µs ± 2%    1.21µs ± 2%  -28.36%  (p=0.029 n=4+4)
    Append/256/16-4       1.69µs ± 3%    1.21µs ± 1%  -28.59%  (p=0.029 n=4+4)
    Append/256/32-4       1.69µs ± 1%    1.23µs ± 1%  -27.08%  (p=0.029 n=4+4)
    Append/256/64-4       1.69µs ± 4%    1.44µs ± 2%  -15.10%  (p=0.029 n=4+4)
    Append/256/128-4      1.68µs ± 2%    1.45µs ± 1%  -13.88%  (p=0.029 n=4+4)
    Append/256/256-4      1.71µs ± 2%    1.96µs ± 1%  +15.12%  (p=0.029 n=4+4)
    Append/256/512-4      1.99µs ± 2%    2.55µs ± 1%  +28.25%  (p=0.029 n=4+4)
    Append/256/1024-4     2.88µs ± 1%    3.85µs ± 2%  +33.64%  (p=0.029 n=4+4)
    Append/512/1-4        3.24µs ± 2%    2.29µs ± 1%  -29.26%  (p=0.029 n=4+4)
    Append/512/2-4        3.25µs ± 3%    2.28µs ± 0%  -29.73%  (p=0.029 n=4+4)
    Append/512/4-4        3.14µs ± 1%    2.32µs ± 4%  -26.09%  (p=0.029 n=4+4)
    Append/512/8-4        3.17µs ± 1%    2.32µs ± 1%  -26.85%  (p=0.029 n=4+4)
    Append/512/16-4       3.16µs ± 4%    2.33µs ± 2%  -26.12%  (p=0.029 n=4+4)
    Append/512/32-4       3.21µs ± 3%    2.42µs ± 8%  -24.67%  (p=0.029 n=4+4)
    Append/512/64-4       3.14µs ± 1%    2.42µs ± 1%  -23.14%  (p=0.029 n=4+4)
    Append/512/128-4      3.17µs ± 1%    2.64µs ± 1%  -16.75%  (p=0.029 n=4+4)
    Append/512/256-4      3.18µs ± 1%    2.87µs ± 2%   -9.63%  (p=0.029 n=4+4)
    Append/512/512-4      3.26µs ± 2%    3.81µs ± 1%  +16.70%  (p=0.029 n=4+4)
    Append/512/1024-4     4.00µs ± 5%    5.14µs ± 4%  +28.44%  (p=0.029 n=4+4)
    Append/1024/1-4       4.53µs ± 1%    4.34µs ± 3%   -4.21%  (p=0.029 n=4+4)
    Append/1024/2-4       4.51µs ± 1%    4.31µs ± 0%   -4.48%  (p=0.029 n=4+4)
    Append/1024/4-4       4.56µs ± 1%    4.33µs ± 2%   -5.07%  (p=0.029 n=4+4)
    Append/1024/8-4       4.52µs ± 2%    4.31µs ± 2%   -4.64%  (p=0.029 n=4+4)
    Append/1024/16-4      4.55µs ± 0%    4.29µs ± 1%   -5.86%  (p=0.029 n=4+4)
    Append/1024/32-4      4.58µs ± 2%    4.36µs ± 1%   -4.91%  (p=0.029 n=4+4)
    Append/1024/64-4      4.54µs ± 2%    4.40µs ± 1%   -2.97%  (p=0.029 n=4+4)
    Append/1024/128-4     4.62µs ± 2%    4.52µs ± 1%     ~     (p=0.057 n=4+4)
    Append/1024/256-4     4.55µs ± 1%    4.88µs ± 1%   +7.21%  (p=0.029 n=4+4)
    Append/1024/512-4     5.33µs ± 1%    5.67µs ± 1%   +6.41%  (p=0.029 n=4+4)
    Append/1024/1024-4    6.89µs ± 5%    6.87µs ± 1%     ~     (p=0.343 n=4+4)
    
    name                old alloc/op   new alloc/op   delta
    Append/1/1-4           24.0B ± 0%     24.0B ± 0%     ~     (all equal)
    Append/1/2-4           40.0B ± 0%     40.0B ± 0%     ~     (all equal)
    Append/1/4-4           56.0B ± 0%     56.0B ± 0%     ~     (all equal)
    Append/1/8-4           88.0B ± 0%     88.0B ± 0%     ~     (all equal)
    Append/1/16-4           152B ± 0%      152B ± 0%     ~     (all equal)
    Append/1/32-4           296B ± 0%      296B ± 0%     ~     (all equal)
    Append/1/64-4           584B ± 0%      584B ± 0%     ~     (all equal)
    Append/1/128-4        1.16kB ± 0%    1.16kB ± 0%     ~     (all equal)
    Append/1/256-4        2.31kB ± 0%    2.31kB ± 0%     ~     (all equal)
    Append/1/512-4        4.87kB ± 0%    4.87kB ± 0%     ~     (all equal)
    Append/1/1024-4       9.48kB ± 0%    9.48kB ± 0%     ~     (all equal)
    Append/2/1-4           48.0B ± 0%     48.0B ± 0%     ~     (all equal)
    Append/2/2-4           48.0B ± 0%     48.0B ± 0%     ~     (all equal)
    Append/2/4-4           64.0B ± 0%     64.0B ± 0%     ~     (all equal)
    Append/2/8-4           96.0B ± 0%     96.0B ± 0%     ~     (all equal)
    Append/2/16-4           160B ± 0%      160B ± 0%     ~     (all equal)
    Append/2/32-4           304B ± 0%      304B ± 0%     ~     (all equal)
    Append/2/64-4           592B ± 0%      592B ± 0%     ~     (all equal)
    Append/2/128-4        1.17kB ± 0%    1.17kB ± 0%     ~     (all equal)
    Append/2/256-4        2.32kB ± 0%    2.32kB ± 0%     ~     (all equal)
    Append/2/512-4        4.88kB ± 0%    4.88kB ± 0%     ~     (all equal)
    Append/2/1024-4       9.49kB ± 0%    9.49kB ± 0%     ~     (all equal)
    Append/4/1-4           96.0B ± 0%     80.0B ± 0%  -16.67%  (p=0.029 n=4+4)
    Append/4/2-4           96.0B ± 0%     80.0B ± 0%  -16.67%  (p=0.029 n=4+4)
    Append/4/4-4           96.0B ± 0%     96.0B ± 0%     ~     (all equal)
    Append/4/8-4            128B ± 0%      128B ± 0%     ~     (all equal)
    Append/4/16-4           192B ± 0%      192B ± 0%     ~     (all equal)
    Append/4/32-4           320B ± 0%      320B ± 0%     ~     (all equal)
    Append/4/64-4           608B ± 0%      608B ± 0%     ~     (all equal)
    Append/4/128-4        1.18kB ± 0%    1.18kB ± 0%     ~     (all equal)
    Append/4/256-4        2.34kB ± 0%    2.34kB ± 0%     ~     (all equal)
    Append/4/512-4        4.90kB ± 0%    4.90kB ± 0%     ~     (all equal)
    Append/4/1024-4       9.50kB ± 0%    9.50kB ± 0%     ~     (all equal)
    Append/8/1-4            192B ± 0%      144B ± 0%  -25.00%  (p=0.029 n=4+4)
    Append/8/2-4            192B ± 0%      144B ± 0%  -25.00%  (p=0.029 n=4+4)
    Append/8/4-4            192B ± 0%      160B ± 0%  -16.67%  (p=0.029 n=4+4)
    Append/8/8-4            192B ± 0%      192B ± 0%     ~     (all equal)
    Append/8/16-4           256B ± 0%      256B ± 0%     ~     (all equal)
    Append/8/32-4           384B ± 0%      384B ± 0%     ~     (all equal)
    Append/8/64-4           640B ± 0%      640B ± 0%     ~     (all equal)
    Append/8/128-4        1.22kB ± 0%    1.22kB ± 0%     ~     (all equal)
    Append/8/256-4        2.37kB ± 0%    2.37kB ± 0%     ~     (all equal)
    Append/8/512-4        4.93kB ± 0%    4.93kB ± 0%     ~     (all equal)
    Append/8/1024-4       9.54kB ± 0%    9.54kB ± 0%     ~     (all equal)
    Append/16/1-4           384B ± 0%      272B ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/16/2-4           384B ± 0%      272B ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/16/4-4           384B ± 0%      288B ± 0%  -25.00%  (p=0.029 n=4+4)
    Append/16/8-4           384B ± 0%      320B ± 0%  -16.67%  (p=0.029 n=4+4)
    Append/16/16-4          384B ± 0%      384B ± 0%     ~     (all equal)
    Append/16/32-4          512B ± 0%      512B ± 0%     ~     (all equal)
    Append/16/64-4          768B ± 0%      768B ± 0%     ~     (all equal)
    Append/16/128-4       1.28kB ± 0%    1.28kB ± 0%     ~     (all equal)
    Append/16/256-4       2.43kB ± 0%    2.43kB ± 0%     ~     (all equal)
    Append/16/512-4       4.99kB ± 0%    4.99kB ± 0%     ~     (all equal)
    Append/16/1024-4      9.60kB ± 0%    9.60kB ± 0%     ~     (all equal)
    Append/32/1-4           768B ± 0%      544B ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/32/2-4           768B ± 0%      544B ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/32/4-4           768B ± 0%      544B ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/32/8-4           768B ± 0%      576B ± 0%  -25.00%  (p=0.029 n=4+4)
    Append/32/16-4          768B ± 0%      640B ± 0%  -16.67%  (p=0.029 n=4+4)
    Append/32/32-4          768B ± 0%      768B ± 0%     ~     (all equal)
    Append/32/64-4        1.02kB ± 0%    1.02kB ± 0%     ~     (all equal)
    Append/32/128-4       1.54kB ± 0%    1.54kB ± 0%     ~     (all equal)
    Append/32/256-4       2.56kB ± 0%    2.56kB ± 0%     ~     (all equal)
    Append/32/512-4       5.12kB ± 0%    5.12kB ± 0%     ~     (all equal)
    Append/32/1024-4      9.73kB ± 0%    9.73kB ± 0%     ~     (all equal)
    Append/64/1-4         1.54kB ± 0%    1.09kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/64/2-4         1.54kB ± 0%    1.09kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/64/4-4         1.54kB ± 0%    1.09kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/64/8-4         1.54kB ± 0%    1.09kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/64/16-4        1.54kB ± 0%    1.15kB ± 0%  -25.00%  (p=0.029 n=4+4)
    Append/64/32-4        1.54kB ± 0%    1.28kB ± 0%  -16.67%  (p=0.029 n=4+4)
    Append/64/64-4        1.54kB ± 0%    1.54kB ± 0%     ~     (all equal)
    Append/64/128-4       2.05kB ± 0%    2.05kB ± 0%     ~     (all equal)
    Append/64/256-4       3.20kB ± 0%    3.20kB ± 0%     ~     (all equal)
    Append/64/512-4       5.38kB ± 0%    5.38kB ± 0%     ~     (all equal)
    Append/64/1024-4      10.0kB ± 0%    10.0kB ± 0%     ~     (all equal)
    Append/128/1-4        3.07kB ± 0%    2.18kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/128/2-4        3.07kB ± 0%    2.18kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/128/4-4        3.07kB ± 0%    2.18kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/128/8-4        3.07kB ± 0%    2.18kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/128/16-4       3.07kB ± 0%    2.18kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/128/32-4       3.07kB ± 0%    2.30kB ± 0%  -25.00%  (p=0.029 n=4+4)
    Append/128/64-4       3.07kB ± 0%    2.56kB ± 0%  -16.67%  (p=0.029 n=4+4)
    Append/128/128-4      3.07kB ± 0%    3.07kB ± 0%     ~     (all equal)
    Append/128/256-4      4.10kB ± 0%    4.10kB ± 0%     ~     (all equal)
    Append/128/512-4      6.40kB ± 0%    6.40kB ± 0%     ~     (all equal)
    Append/128/1024-4     10.5kB ± 0%    10.5kB ± 0%     ~     (all equal)
    Append/256/1-4        6.14kB ± 0%    4.35kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/256/2-4        6.14kB ± 0%    4.35kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/256/4-4        6.14kB ± 0%    4.35kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/256/8-4        6.14kB ± 0%    4.35kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/256/16-4       6.14kB ± 0%    4.35kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/256/32-4       6.14kB ± 0%    4.35kB ± 0%  -29.17%  (p=0.029 n=4+4)
    Append/256/64-4       6.14kB ± 0%    4.74kB ± 0%  -22.92%  (p=0.029 n=4+4)
    Append/256/128-4      6.14kB ± 0%    5.12kB ± 0%  -16.67%  (p=0.029 n=4+4)
    Append/256/256-4      6.14kB ± 0%    6.14kB ± 0%     ~     (all equal)
    Append/256/512-4      8.19kB ± 0%    8.19kB ± 0%     ~     (all equal)
    Append/256/1024-4     12.3kB ± 0%    12.3kB ± 0%     ~     (all equal)
    Append/512/1-4        12.3kB ± 0%     9.0kB ± 0%  -27.08%  (p=0.029 n=4+4)
    Append/512/2-4        12.3kB ± 0%     9.0kB ± 0%  -27.08%  (p=0.029 n=4+4)
    Append/512/4-4        12.3kB ± 0%     9.0kB ± 0%  -27.08%  (p=0.029 n=4+4)
    Append/512/8-4        12.3kB ± 0%     9.0kB ± 0%  -27.08%  (p=0.029 n=4+4)
    Append/512/16-4       12.3kB ± 0%     9.0kB ± 0%  -27.08%  (p=0.029 n=4+4)
    Append/512/32-4       12.3kB ± 0%     9.0kB ± 0%  -27.08%  (p=0.029 n=4+4)
    Append/512/64-4       12.3kB ± 0%     9.0kB ± 0%  -27.08%  (p=0.029 n=4+4)
    Append/512/128-4      12.3kB ± 0%     9.5kB ± 0%  -22.92%  (p=0.029 n=4+4)
    Append/512/256-4      12.3kB ± 0%    10.2kB ± 0%  -16.67%  (p=0.029 n=4+4)
    Append/512/512-4      12.3kB ± 0%    12.3kB ± 0%     ~     (all equal)
    Append/512/1024-4     16.4kB ± 0%    16.4kB ± 0%     ~     (all equal)
    Append/1024/1-4       18.4kB ± 0%    17.7kB ± 0%   -4.17%  (p=0.029 n=4+4)
    Append/1024/2-4       18.4kB ± 0%    17.7kB ± 0%   -4.17%  (p=0.029 n=4+4)
    Append/1024/4-4       18.4kB ± 0%    17.7kB ± 0%   -4.17%  (p=0.029 n=4+4)
    Append/1024/8-4       18.4kB ± 0%    17.7kB ± 0%   -4.17%  (p=0.029 n=4+4)
    Append/1024/16-4      18.4kB ± 0%    17.7kB ± 0%   -4.17%  (p=0.029 n=4+4)
    Append/1024/32-4      18.4kB ± 0%    17.7kB ± 0%   -4.17%  (p=0.029 n=4+4)
    Append/1024/64-4      18.4kB ± 0%    17.7kB ± 0%   -4.17%  (p=0.029 n=4+4)
    Append/1024/128-4     18.4kB ± 0%    17.7kB ± 0%   -4.17%  (p=0.029 n=4+4)
    Append/1024/256-4     18.4kB ± 0%    18.4kB ± 0%     ~     (all equal)
    Append/1024/512-4     21.8kB ± 0%    20.5kB ± 0%   -5.88%  (p=0.029 n=4+4)
    Append/1024/1024-4    28.7kB ± 0%    24.6kB ± 0%  -14.29%  (p=0.029 n=4+4)
    
    name                old allocs/op  new allocs/op  delta
    Append/1/1-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1/2-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1/4-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1/8-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1/16-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1/32-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1/64-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1/128-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1/256-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1/512-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1/1024-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/2/1-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/2/2-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/2/4-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/2/8-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/2/16-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/2/32-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/2/64-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/2/128-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/2/256-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/2/512-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/2/1024-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/4/1-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/4/2-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/4/4-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/4/8-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/4/16-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/4/32-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/4/64-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/4/128-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/4/256-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/4/512-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/4/1024-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/8/1-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/8/2-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/8/4-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/8/8-4            2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/8/16-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/8/32-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/8/64-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/8/128-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/8/256-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/8/512-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/8/1024-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/16/1-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/16/2-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/16/4-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/16/8-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/16/16-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/16/32-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/16/64-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/16/128-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/16/256-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/16/512-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/16/1024-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/32/1-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/32/2-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/32/4-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/32/8-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/32/16-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/32/32-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/32/64-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/32/128-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/32/256-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/32/512-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/32/1024-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/64/1-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/64/2-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/64/4-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/64/8-4           2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/64/16-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/64/32-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/64/64-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/64/128-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/64/256-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/64/512-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/64/1024-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/128/1-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/128/2-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/128/4-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/128/8-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/128/16-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/128/32-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/128/64-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/128/128-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/128/256-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/128/512-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/128/1024-4       2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/256/1-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/256/2-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/256/4-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/256/8-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/256/16-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/256/32-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/256/64-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/256/128-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/256/256-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/256/512-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/256/1024-4       2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/512/1-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/512/2-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/512/4-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/512/8-4          2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/512/16-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/512/32-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/512/64-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/512/128-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/512/256-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/512/512-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/512/1024-4       2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1024/1-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1024/2-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1024/4-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1024/8-4         2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1024/16-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1024/32-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1024/64-4        2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1024/128-4       2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1024/256-4       2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1024/512-4       2.00 ± 0%      2.00 ± 0%     ~     (all equal)
    Append/1024/1024-4      2.00 ± 0%      2.00 ± 0%     ~     (all equal)

old.txt (Append):

    goos: darwin
    goarch: amd64
    pkg: github.com/elpinal/benchmarks/append-or-make
    BenchmarkAppend/1/1-4         	20000000	        80.4 ns/op	      24 B/op	       2 allocs/op
    BenchmarkAppend/1/1-4         	20000000	        79.6 ns/op	      24 B/op	       2 allocs/op
    BenchmarkAppend/1/1-4         	20000000	        80.7 ns/op	      24 B/op	       2 allocs/op
    BenchmarkAppend/1/1-4         	20000000	        84.6 ns/op	      24 B/op	       2 allocs/op
    BenchmarkAppend/1/2-4         	20000000	        91.4 ns/op	      40 B/op	       2 allocs/op
    BenchmarkAppend/1/2-4         	20000000	        90.8 ns/op	      40 B/op	       2 allocs/op
    BenchmarkAppend/1/2-4         	20000000	        87.4 ns/op	      40 B/op	       2 allocs/op
    BenchmarkAppend/1/2-4         	20000000	        85.0 ns/op	      40 B/op	       2 allocs/op
    BenchmarkAppend/1/4-4         	20000000	        87.4 ns/op	      56 B/op	       2 allocs/op
    BenchmarkAppend/1/4-4         	20000000	        87.5 ns/op	      56 B/op	       2 allocs/op
    BenchmarkAppend/1/4-4         	20000000	        88.3 ns/op	      56 B/op	       2 allocs/op
    BenchmarkAppend/1/4-4         	20000000	        87.7 ns/op	      56 B/op	       2 allocs/op
    BenchmarkAppend/1/8-4         	20000000	        95.4 ns/op	      88 B/op	       2 allocs/op
    BenchmarkAppend/1/8-4         	20000000	        95.3 ns/op	      88 B/op	       2 allocs/op
    BenchmarkAppend/1/8-4         	20000000	        95.2 ns/op	      88 B/op	       2 allocs/op
    BenchmarkAppend/1/8-4         	20000000	        95.7 ns/op	      88 B/op	       2 allocs/op
    BenchmarkAppend/1/16-4        	10000000	       123 ns/op	     152 B/op	       2 allocs/op
    BenchmarkAppend/1/16-4        	10000000	       115 ns/op	     152 B/op	       2 allocs/op
    BenchmarkAppend/1/16-4        	10000000	       113 ns/op	     152 B/op	       2 allocs/op
    BenchmarkAppend/1/16-4        	10000000	       110 ns/op	     152 B/op	       2 allocs/op
    BenchmarkAppend/1/32-4        	10000000	       144 ns/op	     296 B/op	       2 allocs/op
    BenchmarkAppend/1/32-4        	10000000	       144 ns/op	     296 B/op	       2 allocs/op
    BenchmarkAppend/1/32-4        	10000000	       145 ns/op	     296 B/op	       2 allocs/op
    BenchmarkAppend/1/32-4        	10000000	       144 ns/op	     296 B/op	       2 allocs/op
    BenchmarkAppend/1/64-4        	 5000000	       250 ns/op	     584 B/op	       2 allocs/op
    BenchmarkAppend/1/64-4        	10000000	       210 ns/op	     584 B/op	       2 allocs/op
    BenchmarkAppend/1/64-4        	10000000	       210 ns/op	     584 B/op	       2 allocs/op
    BenchmarkAppend/1/64-4        	10000000	       211 ns/op	     584 B/op	       2 allocs/op
    BenchmarkAppend/1/128-4       	 5000000	       353 ns/op	    1160 B/op	       2 allocs/op
    BenchmarkAppend/1/128-4       	 5000000	       362 ns/op	    1160 B/op	       2 allocs/op
    BenchmarkAppend/1/128-4       	 5000000	       350 ns/op	    1160 B/op	       2 allocs/op
    BenchmarkAppend/1/128-4       	 5000000	       348 ns/op	    1160 B/op	       2 allocs/op
    BenchmarkAppend/1/256-4       	 3000000	       529 ns/op	    2312 B/op	       2 allocs/op
    BenchmarkAppend/1/256-4       	 3000000	       527 ns/op	    2312 B/op	       2 allocs/op
    BenchmarkAppend/1/256-4       	 3000000	       533 ns/op	    2312 B/op	       2 allocs/op
    BenchmarkAppend/1/256-4       	 3000000	       528 ns/op	    2312 B/op	       2 allocs/op
    BenchmarkAppend/1/512-4       	 1000000	      1050 ns/op	    4872 B/op	       2 allocs/op
    BenchmarkAppend/1/512-4       	 1000000	      1042 ns/op	    4872 B/op	       2 allocs/op
    BenchmarkAppend/1/512-4       	 1000000	      1027 ns/op	    4872 B/op	       2 allocs/op
    BenchmarkAppend/1/512-4       	 1000000	      1031 ns/op	    4872 B/op	       2 allocs/op
    BenchmarkAppend/1/1024-4      	 1000000	      2006 ns/op	    9480 B/op	       2 allocs/op
    BenchmarkAppend/1/1024-4      	 1000000	      2005 ns/op	    9480 B/op	       2 allocs/op
    BenchmarkAppend/1/1024-4      	 1000000	      2009 ns/op	    9480 B/op	       2 allocs/op
    BenchmarkAppend/1/1024-4      	  500000	      2007 ns/op	    9480 B/op	       2 allocs/op
    BenchmarkAppend/2/1-4         	20000000	        94.5 ns/op	      48 B/op	       2 allocs/op
    BenchmarkAppend/2/1-4         	20000000	        95.0 ns/op	      48 B/op	       2 allocs/op
    BenchmarkAppend/2/1-4         	20000000	        95.4 ns/op	      48 B/op	       2 allocs/op
    BenchmarkAppend/2/1-4         	20000000	        94.3 ns/op	      48 B/op	       2 allocs/op
    BenchmarkAppend/2/2-4         	20000000	        93.9 ns/op	      48 B/op	       2 allocs/op
    BenchmarkAppend/2/2-4         	20000000	        96.1 ns/op	      48 B/op	       2 allocs/op
    BenchmarkAppend/2/2-4         	20000000	        94.9 ns/op	      48 B/op	       2 allocs/op
    BenchmarkAppend/2/2-4         	20000000	        94.6 ns/op	      48 B/op	       2 allocs/op
    BenchmarkAppend/2/4-4         	20000000	        96.9 ns/op	      64 B/op	       2 allocs/op
    BenchmarkAppend/2/4-4         	20000000	        98.0 ns/op	      64 B/op	       2 allocs/op
    BenchmarkAppend/2/4-4         	20000000	        97.0 ns/op	      64 B/op	       2 allocs/op
    BenchmarkAppend/2/4-4         	20000000	        97.2 ns/op	      64 B/op	       2 allocs/op
    BenchmarkAppend/2/8-4         	10000000	       106 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/2/8-4         	10000000	       105 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/2/8-4         	10000000	       105 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/2/8-4         	10000000	       105 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/2/16-4        	10000000	       122 ns/op	     160 B/op	       2 allocs/op
    BenchmarkAppend/2/16-4        	10000000	       120 ns/op	     160 B/op	       2 allocs/op
    BenchmarkAppend/2/16-4        	10000000	       123 ns/op	     160 B/op	       2 allocs/op
    BenchmarkAppend/2/16-4        	10000000	       123 ns/op	     160 B/op	       2 allocs/op
    BenchmarkAppend/2/32-4        	10000000	       158 ns/op	     304 B/op	       2 allocs/op
    BenchmarkAppend/2/32-4        	10000000	       157 ns/op	     304 B/op	       2 allocs/op
    BenchmarkAppend/2/32-4        	10000000	       155 ns/op	     304 B/op	       2 allocs/op
    BenchmarkAppend/2/32-4        	10000000	       157 ns/op	     304 B/op	       2 allocs/op
    BenchmarkAppend/2/64-4        	10000000	       219 ns/op	     592 B/op	       2 allocs/op
    BenchmarkAppend/2/64-4        	10000000	       220 ns/op	     592 B/op	       2 allocs/op
    BenchmarkAppend/2/64-4        	10000000	       219 ns/op	     592 B/op	       2 allocs/op
    BenchmarkAppend/2/64-4        	10000000	       220 ns/op	     592 B/op	       2 allocs/op
    BenchmarkAppend/2/128-4       	 5000000	       360 ns/op	    1168 B/op	       2 allocs/op
    BenchmarkAppend/2/128-4       	 5000000	       364 ns/op	    1168 B/op	       2 allocs/op
    BenchmarkAppend/2/128-4       	 5000000	       363 ns/op	    1168 B/op	       2 allocs/op
    BenchmarkAppend/2/128-4       	 5000000	       361 ns/op	    1168 B/op	       2 allocs/op
    BenchmarkAppend/2/256-4       	 3000000	       543 ns/op	    2320 B/op	       2 allocs/op
    BenchmarkAppend/2/256-4       	 3000000	       549 ns/op	    2320 B/op	       2 allocs/op
    BenchmarkAppend/2/256-4       	 2000000	       550 ns/op	    2320 B/op	       2 allocs/op
    BenchmarkAppend/2/256-4       	 3000000	       544 ns/op	    2320 B/op	       2 allocs/op
    BenchmarkAppend/2/512-4       	 1000000	      1051 ns/op	    4880 B/op	       2 allocs/op
    BenchmarkAppend/2/512-4       	 1000000	      1061 ns/op	    4880 B/op	       2 allocs/op
    BenchmarkAppend/2/512-4       	 1000000	      1050 ns/op	    4880 B/op	       2 allocs/op
    BenchmarkAppend/2/512-4       	 1000000	      1057 ns/op	    4880 B/op	       2 allocs/op
    BenchmarkAppend/2/1024-4      	 1000000	      2012 ns/op	    9488 B/op	       2 allocs/op
    BenchmarkAppend/2/1024-4      	 1000000	      2040 ns/op	    9488 B/op	       2 allocs/op
    BenchmarkAppend/2/1024-4      	 1000000	      2021 ns/op	    9488 B/op	       2 allocs/op
    BenchmarkAppend/2/1024-4      	 1000000	      2035 ns/op	    9488 B/op	       2 allocs/op
    BenchmarkAppend/4/1-4         	10000000	       104 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/4/1-4         	20000000	       104 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/4/1-4         	20000000	       104 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/4/1-4         	10000000	       104 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/4/2-4         	10000000	       105 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/4/2-4         	10000000	       104 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/4/2-4         	20000000	       105 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/4/2-4         	10000000	       107 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/4/4-4         	10000000	       105 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/4/4-4         	20000000	       105 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/4/4-4         	10000000	       104 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/4/4-4         	20000000	       104 ns/op	      96 B/op	       2 allocs/op
    BenchmarkAppend/4/8-4         	10000000	       112 ns/op	     128 B/op	       2 allocs/op
    BenchmarkAppend/4/8-4         	10000000	       112 ns/op	     128 B/op	       2 allocs/op
    BenchmarkAppend/4/8-4         	10000000	       112 ns/op	     128 B/op	       2 allocs/op
    BenchmarkAppend/4/8-4         	10000000	       112 ns/op	     128 B/op	       2 allocs/op
    BenchmarkAppend/4/16-4        	10000000	       127 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/4/16-4        	10000000	       127 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/4/16-4        	10000000	       128 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/4/16-4        	10000000	       128 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/4/32-4        	10000000	       159 ns/op	     320 B/op	       2 allocs/op
    BenchmarkAppend/4/32-4        	10000000	       158 ns/op	     320 B/op	       2 allocs/op
    BenchmarkAppend/4/32-4        	10000000	       159 ns/op	     320 B/op	       2 allocs/op
    BenchmarkAppend/4/32-4        	10000000	       157 ns/op	     320 B/op	       2 allocs/op
    BenchmarkAppend/4/64-4        	 5000000	       222 ns/op	     608 B/op	       2 allocs/op
    BenchmarkAppend/4/64-4        	 5000000	       223 ns/op	     608 B/op	       2 allocs/op
    BenchmarkAppend/4/64-4        	 5000000	       222 ns/op	     608 B/op	       2 allocs/op
    BenchmarkAppend/4/64-4        	 5000000	       226 ns/op	     608 B/op	       2 allocs/op
    BenchmarkAppend/4/128-4       	 5000000	       376 ns/op	    1184 B/op	       2 allocs/op
    BenchmarkAppend/4/128-4       	 5000000	       360 ns/op	    1184 B/op	       2 allocs/op
    BenchmarkAppend/4/128-4       	 5000000	       362 ns/op	    1184 B/op	       2 allocs/op
    BenchmarkAppend/4/128-4       	 5000000	       362 ns/op	    1184 B/op	       2 allocs/op
    BenchmarkAppend/4/256-4       	 3000000	       548 ns/op	    2336 B/op	       2 allocs/op
    BenchmarkAppend/4/256-4       	 2000000	       565 ns/op	    2336 B/op	       2 allocs/op
    BenchmarkAppend/4/256-4       	 3000000	       548 ns/op	    2336 B/op	       2 allocs/op
    BenchmarkAppend/4/256-4       	 3000000	       553 ns/op	    2336 B/op	       2 allocs/op
    BenchmarkAppend/4/512-4       	 1000000	      1051 ns/op	    4896 B/op	       2 allocs/op
    BenchmarkAppend/4/512-4       	 1000000	      1060 ns/op	    4896 B/op	       2 allocs/op
    BenchmarkAppend/4/512-4       	 1000000	      1056 ns/op	    4896 B/op	       2 allocs/op
    BenchmarkAppend/4/512-4       	 1000000	      1046 ns/op	    4896 B/op	       2 allocs/op
    BenchmarkAppend/4/1024-4      	 1000000	      2037 ns/op	    9504 B/op	       2 allocs/op
    BenchmarkAppend/4/1024-4      	 1000000	      2026 ns/op	    9504 B/op	       2 allocs/op
    BenchmarkAppend/4/1024-4      	 1000000	      2051 ns/op	    9504 B/op	       2 allocs/op
    BenchmarkAppend/4/1024-4      	 1000000	      2064 ns/op	    9504 B/op	       2 allocs/op
    BenchmarkAppend/8/1-4         	10000000	       128 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/1-4         	10000000	       128 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/1-4         	10000000	       127 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/1-4         	10000000	       127 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/2-4         	10000000	       128 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/2-4         	10000000	       127 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/2-4         	10000000	       127 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/2-4         	10000000	       127 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/4-4         	10000000	       128 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/4-4         	10000000	       127 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/4-4         	10000000	       127 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/4-4         	10000000	       129 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/8-4         	10000000	       128 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/8-4         	10000000	       128 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/8-4         	10000000	       128 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/8-4         	10000000	       129 ns/op	     192 B/op	       2 allocs/op
    BenchmarkAppend/8/16-4        	10000000	       143 ns/op	     256 B/op	       2 allocs/op
    BenchmarkAppend/8/16-4        	10000000	       143 ns/op	     256 B/op	       2 allocs/op
    BenchmarkAppend/8/16-4        	10000000	       144 ns/op	     256 B/op	       2 allocs/op
    BenchmarkAppend/8/16-4        	10000000	       143 ns/op	     256 B/op	       2 allocs/op
    BenchmarkAppend/8/32-4        	10000000	       177 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/8/32-4        	10000000	       177 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/8/32-4        	10000000	       177 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/8/32-4        	10000000	       177 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/8/64-4        	 5000000	       231 ns/op	     640 B/op	       2 allocs/op
    BenchmarkAppend/8/64-4        	 5000000	       231 ns/op	     640 B/op	       2 allocs/op
    BenchmarkAppend/8/64-4        	 5000000	       231 ns/op	     640 B/op	       2 allocs/op
    BenchmarkAppend/8/64-4        	 5000000	       232 ns/op	     640 B/op	       2 allocs/op
    BenchmarkAppend/8/128-4       	 5000000	       371 ns/op	    1216 B/op	       2 allocs/op
    BenchmarkAppend/8/128-4       	 5000000	       371 ns/op	    1216 B/op	       2 allocs/op
    BenchmarkAppend/8/128-4       	 5000000	       372 ns/op	    1216 B/op	       2 allocs/op
    BenchmarkAppend/8/128-4       	 5000000	       370 ns/op	    1216 B/op	       2 allocs/op
    BenchmarkAppend/8/256-4       	 3000000	       564 ns/op	    2368 B/op	       2 allocs/op
    BenchmarkAppend/8/256-4       	 3000000	       564 ns/op	    2368 B/op	       2 allocs/op
    BenchmarkAppend/8/256-4       	 3000000	       562 ns/op	    2368 B/op	       2 allocs/op
    BenchmarkAppend/8/256-4       	 3000000	       569 ns/op	    2368 B/op	       2 allocs/op
    BenchmarkAppend/8/512-4       	 1000000	      1074 ns/op	    4928 B/op	       2 allocs/op
    BenchmarkAppend/8/512-4       	 1000000	      1061 ns/op	    4928 B/op	       2 allocs/op
    BenchmarkAppend/8/512-4       	 1000000	      1073 ns/op	    4928 B/op	       2 allocs/op
    BenchmarkAppend/8/512-4       	 1000000	      1070 ns/op	    4928 B/op	       2 allocs/op
    BenchmarkAppend/8/1024-4      	 1000000	      2128 ns/op	    9536 B/op	       2 allocs/op
    BenchmarkAppend/8/1024-4      	 1000000	      2043 ns/op	    9536 B/op	       2 allocs/op
    BenchmarkAppend/8/1024-4      	  500000	      2054 ns/op	    9536 B/op	       2 allocs/op
    BenchmarkAppend/8/1024-4      	 1000000	      2043 ns/op	    9536 B/op	       2 allocs/op
    BenchmarkAppend/16/1-4        	10000000	       174 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/1-4        	10000000	       175 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/1-4        	10000000	       175 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/1-4        	10000000	       174 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/2-4        	10000000	       175 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/2-4        	10000000	       175 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/2-4        	10000000	       174 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/2-4        	10000000	       175 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/4-4        	10000000	       178 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/4-4        	10000000	       175 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/4-4        	10000000	       174 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/4-4        	10000000	       174 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/8-4        	10000000	       176 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/8-4        	10000000	       175 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/8-4        	10000000	       180 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/8-4        	10000000	       177 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/16-4       	10000000	       176 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/16-4       	10000000	       190 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/16-4       	10000000	       185 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/16-4       	10000000	       177 ns/op	     384 B/op	       2 allocs/op
    BenchmarkAppend/16/32-4       	10000000	       208 ns/op	     512 B/op	       2 allocs/op
    BenchmarkAppend/16/32-4       	10000000	       208 ns/op	     512 B/op	       2 allocs/op
    BenchmarkAppend/16/32-4       	10000000	       209 ns/op	     512 B/op	       2 allocs/op
    BenchmarkAppend/16/32-4       	10000000	       209 ns/op	     512 B/op	       2 allocs/op
    BenchmarkAppend/16/64-4       	 5000000	       278 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/16/64-4       	 5000000	       277 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/16/64-4       	 5000000	       273 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/16/64-4       	 5000000	       279 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/16/128-4      	 3000000	       392 ns/op	    1280 B/op	       2 allocs/op
    BenchmarkAppend/16/128-4      	 3000000	       400 ns/op	    1280 B/op	       2 allocs/op
    BenchmarkAppend/16/128-4      	 3000000	       396 ns/op	    1280 B/op	       2 allocs/op
    BenchmarkAppend/16/128-4      	 3000000	       396 ns/op	    1280 B/op	       2 allocs/op
    BenchmarkAppend/16/256-4      	 2000000	       588 ns/op	    2432 B/op	       2 allocs/op
    BenchmarkAppend/16/256-4      	 2000000	       638 ns/op	    2432 B/op	       2 allocs/op
    BenchmarkAppend/16/256-4      	 2000000	       587 ns/op	    2432 B/op	       2 allocs/op
    BenchmarkAppend/16/256-4      	 2000000	       580 ns/op	    2432 B/op	       2 allocs/op
    BenchmarkAppend/16/512-4      	 1000000	      1160 ns/op	    4992 B/op	       2 allocs/op
    BenchmarkAppend/16/512-4      	 1000000	      1097 ns/op	    4992 B/op	       2 allocs/op
    BenchmarkAppend/16/512-4      	 1000000	      1104 ns/op	    4992 B/op	       2 allocs/op
    BenchmarkAppend/16/512-4      	 1000000	      1101 ns/op	    4992 B/op	       2 allocs/op
    BenchmarkAppend/16/1024-4     	 1000000	      2081 ns/op	    9600 B/op	       2 allocs/op
    BenchmarkAppend/16/1024-4     	  500000	      2052 ns/op	    9600 B/op	       2 allocs/op
    BenchmarkAppend/16/1024-4     	 1000000	      2069 ns/op	    9600 B/op	       2 allocs/op
    BenchmarkAppend/16/1024-4     	 1000000	      2068 ns/op	    9600 B/op	       2 allocs/op
    BenchmarkAppend/32/1-4        	 5000000	       266 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/1-4        	 5000000	       267 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/1-4        	 5000000	       267 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/1-4        	 5000000	       266 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/2-4        	 5000000	       267 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/2-4        	 5000000	       273 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/2-4        	 5000000	       271 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/2-4        	 5000000	       272 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/4-4        	 5000000	       270 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/4-4        	 5000000	       268 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/4-4        	 5000000	       270 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/4-4        	 5000000	       270 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/8-4        	 5000000	       272 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/8-4        	 5000000	       274 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/8-4        	 5000000	       273 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/8-4        	 5000000	       274 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/16-4       	 5000000	       272 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/16-4       	 5000000	       276 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/16-4       	 5000000	       274 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/16-4       	 5000000	       274 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/32-4       	 5000000	       273 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/32-4       	 5000000	       272 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/32-4       	 5000000	       271 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/32-4       	 5000000	       271 ns/op	     768 B/op	       2 allocs/op
    BenchmarkAppend/32/64-4       	 5000000	       349 ns/op	    1024 B/op	       2 allocs/op
    BenchmarkAppend/32/64-4       	 5000000	       359 ns/op	    1024 B/op	       2 allocs/op
    BenchmarkAppend/32/64-4       	 5000000	       349 ns/op	    1024 B/op	       2 allocs/op
    BenchmarkAppend/32/64-4       	 5000000	       350 ns/op	    1024 B/op	       2 allocs/op
    BenchmarkAppend/32/128-4      	 3000000	       471 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/32/128-4      	 3000000	       476 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/32/128-4      	 3000000	       476 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/32/128-4      	 3000000	       477 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/32/256-4      	 2000000	       626 ns/op	    2560 B/op	       2 allocs/op
    BenchmarkAppend/32/256-4      	 2000000	       631 ns/op	    2560 B/op	       2 allocs/op
    BenchmarkAppend/32/256-4      	 2000000	       624 ns/op	    2560 B/op	       2 allocs/op
    BenchmarkAppend/32/256-4      	 2000000	       643 ns/op	    2560 B/op	       2 allocs/op
    BenchmarkAppend/32/512-4      	 1000000	      1196 ns/op	    5120 B/op	       2 allocs/op
    BenchmarkAppend/32/512-4      	 1000000	      1182 ns/op	    5120 B/op	       2 allocs/op
    BenchmarkAppend/32/512-4      	 1000000	      1175 ns/op	    5120 B/op	       2 allocs/op
    BenchmarkAppend/32/512-4      	 1000000	      1174 ns/op	    5120 B/op	       2 allocs/op
    BenchmarkAppend/32/1024-4     	 1000000	      2129 ns/op	    9728 B/op	       2 allocs/op
    BenchmarkAppend/32/1024-4     	 1000000	      2234 ns/op	    9728 B/op	       2 allocs/op
    BenchmarkAppend/32/1024-4     	 1000000	      2176 ns/op	    9728 B/op	       2 allocs/op
    BenchmarkAppend/32/1024-4     	 1000000	      2140 ns/op	    9728 B/op	       2 allocs/op
    BenchmarkAppend/64/1-4        	 3000000	       479 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/1-4        	 3000000	       480 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/1-4        	 3000000	       472 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/1-4        	 3000000	       471 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/2-4        	 3000000	       481 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/2-4        	 3000000	       483 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/2-4        	 3000000	       471 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/2-4        	 3000000	       476 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/4-4        	 3000000	       471 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/4-4        	 3000000	       475 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/4-4        	 3000000	       476 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/4-4        	 3000000	       473 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/8-4        	 3000000	       478 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/8-4        	 3000000	       481 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/8-4        	 3000000	       472 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/8-4        	 3000000	       476 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/16-4       	 3000000	       483 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/16-4       	 3000000	       487 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/16-4       	 3000000	       504 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/16-4       	 3000000	       484 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/32-4       	 3000000	       484 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/32-4       	 3000000	       480 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/32-4       	 3000000	       481 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/32-4       	 3000000	       490 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/64-4       	 3000000	       475 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/64-4       	 3000000	       479 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/64-4       	 3000000	       477 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/64-4       	 3000000	       489 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkAppend/64/128-4      	 2000000	       615 ns/op	    2048 B/op	       2 allocs/op
    BenchmarkAppend/64/128-4      	 2000000	       616 ns/op	    2048 B/op	       2 allocs/op
    BenchmarkAppend/64/128-4      	 2000000	       633 ns/op	    2048 B/op	       2 allocs/op
    BenchmarkAppend/64/128-4      	 2000000	       624 ns/op	    2048 B/op	       2 allocs/op
    BenchmarkAppend/64/256-4      	 2000000	       883 ns/op	    3200 B/op	       2 allocs/op
    BenchmarkAppend/64/256-4      	 2000000	       886 ns/op	    3200 B/op	       2 allocs/op
    BenchmarkAppend/64/256-4      	 2000000	       889 ns/op	    3200 B/op	       2 allocs/op
    BenchmarkAppend/64/256-4      	 2000000	       875 ns/op	    3200 B/op	       2 allocs/op
    BenchmarkAppend/64/512-4      	 1000000	      1332 ns/op	    5376 B/op	       2 allocs/op
    BenchmarkAppend/64/512-4      	 1000000	      1255 ns/op	    5376 B/op	       2 allocs/op
    BenchmarkAppend/64/512-4      	 1000000	      1270 ns/op	    5376 B/op	       2 allocs/op
    BenchmarkAppend/64/512-4      	 1000000	      1253 ns/op	    5376 B/op	       2 allocs/op
    BenchmarkAppend/64/1024-4     	 1000000	      2181 ns/op	    9984 B/op	       2 allocs/op
    BenchmarkAppend/64/1024-4     	 1000000	      2189 ns/op	    9984 B/op	       2 allocs/op
    BenchmarkAppend/64/1024-4     	  500000	      2204 ns/op	    9984 B/op	       2 allocs/op
    BenchmarkAppend/64/1024-4     	 1000000	      2234 ns/op	    9984 B/op	       2 allocs/op
    BenchmarkAppend/128/1-4       	 2000000	       875 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/1-4       	 2000000	       882 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/1-4       	 2000000	       879 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/1-4       	 2000000	       880 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/2-4       	 2000000	       884 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/2-4       	 2000000	       924 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/2-4       	 2000000	       866 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/2-4       	 2000000	       900 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/4-4       	 2000000	       849 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/4-4       	 2000000	       862 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/4-4       	 2000000	       858 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/4-4       	 2000000	       857 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/8-4       	 2000000	       863 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/8-4       	 2000000	       859 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/8-4       	 2000000	       858 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/8-4       	 2000000	       884 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/16-4      	 2000000	       872 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/16-4      	 2000000	       886 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/16-4      	 2000000	       885 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/16-4      	 2000000	       878 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/32-4      	 2000000	       912 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/32-4      	 2000000	       899 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/32-4      	 2000000	       901 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/32-4      	 2000000	       905 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/64-4      	 2000000	       866 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/64-4      	 2000000	       861 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/64-4      	 2000000	       872 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/64-4      	 2000000	       872 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/128-4     	 2000000	       874 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/128-4     	 2000000	       869 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/128-4     	 2000000	       869 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/128-4     	 2000000	       859 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkAppend/128/256-4     	 1000000	      1045 ns/op	    4096 B/op	       2 allocs/op
    BenchmarkAppend/128/256-4     	 1000000	      1031 ns/op	    4096 B/op	       2 allocs/op
    BenchmarkAppend/128/256-4     	 1000000	      1032 ns/op	    4096 B/op	       2 allocs/op
    BenchmarkAppend/128/256-4     	 1000000	      1028 ns/op	    4096 B/op	       2 allocs/op
    BenchmarkAppend/128/512-4     	 1000000	      1517 ns/op	    6400 B/op	       2 allocs/op
    BenchmarkAppend/128/512-4     	 1000000	      1513 ns/op	    6400 B/op	       2 allocs/op
    BenchmarkAppend/128/512-4     	 1000000	      1518 ns/op	    6400 B/op	       2 allocs/op
    BenchmarkAppend/128/512-4     	 1000000	      1537 ns/op	    6400 B/op	       2 allocs/op
    BenchmarkAppend/128/1024-4    	  500000	      2429 ns/op	   10496 B/op	       2 allocs/op
    BenchmarkAppend/128/1024-4    	  500000	      2386 ns/op	   10496 B/op	       2 allocs/op
    BenchmarkAppend/128/1024-4    	 1000000	      2508 ns/op	   10496 B/op	       2 allocs/op
    BenchmarkAppend/128/1024-4    	 1000000	      2441 ns/op	   10496 B/op	       2 allocs/op
    BenchmarkAppend/256/1-4       	 1000000	      1735 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/1-4       	 1000000	      1684 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/1-4       	 1000000	      1680 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/1-4       	 1000000	      1677 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/2-4       	 1000000	      1696 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/2-4       	 1000000	      1666 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/2-4       	 1000000	      1667 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/2-4       	 1000000	      1686 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/4-4       	 1000000	      1648 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/4-4       	 1000000	      1663 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/4-4       	 1000000	      1661 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/4-4       	 1000000	      1629 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/8-4       	 1000000	      1723 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/8-4       	 1000000	      1655 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/8-4       	 1000000	      1650 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/8-4       	 1000000	      1713 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/16-4      	 1000000	      1664 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/16-4      	 1000000	      1687 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/16-4      	 1000000	      1733 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/16-4      	 1000000	      1676 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/32-4      	 1000000	      1691 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/32-4      	 1000000	      1701 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/32-4      	 1000000	      1687 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/32-4      	 1000000	      1690 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/64-4      	 1000000	      1665 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/64-4      	 1000000	      1761 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/64-4      	 1000000	      1673 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/64-4      	 1000000	      1668 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/128-4     	 1000000	      1652 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/128-4     	 1000000	      1676 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/128-4     	 1000000	      1680 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/128-4     	 1000000	      1722 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/256-4     	 1000000	      1706 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/256-4     	 1000000	      1736 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/256-4     	 1000000	      1698 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/256-4     	 1000000	      1684 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkAppend/256/512-4     	 1000000	      2037 ns/op	    8192 B/op	       2 allocs/op
    BenchmarkAppend/256/512-4     	 1000000	      1988 ns/op	    8192 B/op	       2 allocs/op
    BenchmarkAppend/256/512-4     	 1000000	      1981 ns/op	    8192 B/op	       2 allocs/op
    BenchmarkAppend/256/512-4     	 1000000	      1960 ns/op	    8192 B/op	       2 allocs/op
    BenchmarkAppend/256/1024-4    	  500000	      2882 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/256/1024-4    	  500000	      2864 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/256/1024-4    	  500000	      2889 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/256/1024-4    	  500000	      2897 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/1-4       	  500000	      3293 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/1-4       	  500000	      3261 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/1-4       	  500000	      3174 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/1-4       	  500000	      3239 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/2-4       	  500000	      3225 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/2-4       	  500000	      3262 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/2-4       	  500000	      3171 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/2-4       	  300000	      3339 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/4-4       	  500000	      3139 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/4-4       	  500000	      3150 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/4-4       	  500000	      3101 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/4-4       	  500000	      3167 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/8-4       	  500000	      3162 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/8-4       	  500000	      3181 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/8-4       	  500000	      3211 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/8-4       	  500000	      3143 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/16-4      	  500000	      3150 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/16-4      	  500000	      3285 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/16-4      	  500000	      3097 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/16-4      	  500000	      3109 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/32-4      	  500000	      3229 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/32-4      	  500000	      3296 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/32-4      	  500000	      3127 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/32-4      	  500000	      3173 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/64-4      	  500000	      3157 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/64-4      	  500000	      3146 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/64-4      	  500000	      3111 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/64-4      	  500000	      3164 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/128-4     	  500000	      3157 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/128-4     	  500000	      3161 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/128-4     	  500000	      3209 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/128-4     	  500000	      3148 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/256-4     	  500000	      3169 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/256-4     	  500000	      3193 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/256-4     	  500000	      3138 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/256-4     	  500000	      3203 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/512-4     	  500000	      3249 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/512-4     	  500000	      3266 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/512-4     	  500000	      3323 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/512-4     	  500000	      3212 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkAppend/512/1024-4    	  300000	      3942 ns/op	   16384 B/op	       2 allocs/op
    BenchmarkAppend/512/1024-4    	  300000	      3939 ns/op	   16384 B/op	       2 allocs/op
    BenchmarkAppend/512/1024-4    	  300000	      4188 ns/op	   16384 B/op	       2 allocs/op
    BenchmarkAppend/512/1024-4    	  300000	      3934 ns/op	   16384 B/op	       2 allocs/op
    BenchmarkAppend/1024/1-4      	  300000	      4495 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/1-4      	  300000	      4532 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/1-4      	  300000	      4549 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/1-4      	  300000	      4553 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/2-4      	  300000	      4539 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/2-4      	  300000	      4461 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/2-4      	  300000	      4475 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/2-4      	  300000	      4564 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/4-4      	  300000	      4534 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/4-4      	  300000	      4601 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/4-4      	  300000	      4541 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/4-4      	  300000	      4564 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/8-4      	  300000	      4544 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/8-4      	  300000	      4583 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/8-4      	  300000	      4435 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/8-4      	  300000	      4499 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/16-4     	  300000	      4554 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/16-4     	  300000	      4547 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/16-4     	  300000	      4547 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/16-4     	  300000	      4561 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/32-4     	  300000	      4578 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/32-4     	  300000	      4537 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/32-4     	  300000	      4552 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/32-4     	  300000	      4665 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/64-4     	  300000	      4592 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/64-4     	  300000	      4633 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/64-4     	  300000	      4467 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/64-4     	  300000	      4450 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/128-4    	  300000	      4724 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/128-4    	  300000	      4588 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/128-4    	  300000	      4571 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/128-4    	  300000	      4585 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/256-4    	  300000	      4536 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/256-4    	  300000	      4544 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/256-4    	  300000	      4585 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/256-4    	  300000	      4542 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkAppend/1024/512-4    	  200000	      5264 ns/op	   21760 B/op	       2 allocs/op
    BenchmarkAppend/1024/512-4    	  200000	      5394 ns/op	   21760 B/op	       2 allocs/op
    BenchmarkAppend/1024/512-4    	  300000	      5348 ns/op	   21760 B/op	       2 allocs/op
    BenchmarkAppend/1024/512-4    	  200000	      5309 ns/op	   21760 B/op	       2 allocs/op
    BenchmarkAppend/1024/1024-4   	  200000	      6785 ns/op	   28672 B/op	       2 allocs/op
    BenchmarkAppend/1024/1024-4   	  200000	      6805 ns/op	   28672 B/op	       2 allocs/op
    BenchmarkAppend/1024/1024-4   	  200000	      6728 ns/op	   28672 B/op	       2 allocs/op
    BenchmarkAppend/1024/1024-4   	  200000	      7254 ns/op	   28672 B/op	       2 allocs/op
    PASS
    ok  	github.com/elpinal/benchmarks/append-or-make	887.791s

new.txt (Make):

    goos: darwin
    goarch: amd64
    pkg: github.com/elpinal/benchmarks/append-or-make
    BenchmarkMake/1/1-4         	20000000	        72.8 ns/op	      24 B/op	       2 allocs/op
    BenchmarkMake/1/1-4         	20000000	        70.8 ns/op	      24 B/op	       2 allocs/op
    BenchmarkMake/1/1-4         	20000000	        72.5 ns/op	      24 B/op	       2 allocs/op
    BenchmarkMake/1/1-4         	20000000	        72.1 ns/op	      24 B/op	       2 allocs/op
    BenchmarkMake/1/2-4         	20000000	        76.4 ns/op	      40 B/op	       2 allocs/op
    BenchmarkMake/1/2-4         	20000000	        76.6 ns/op	      40 B/op	       2 allocs/op
    BenchmarkMake/1/2-4         	20000000	        76.7 ns/op	      40 B/op	       2 allocs/op
    BenchmarkMake/1/2-4         	20000000	        76.7 ns/op	      40 B/op	       2 allocs/op
    BenchmarkMake/1/4-4         	20000000	        82.2 ns/op	      56 B/op	       2 allocs/op
    BenchmarkMake/1/4-4         	20000000	        82.9 ns/op	      56 B/op	       2 allocs/op
    BenchmarkMake/1/4-4         	20000000	        82.3 ns/op	      56 B/op	       2 allocs/op
    BenchmarkMake/1/4-4         	20000000	        82.7 ns/op	      56 B/op	       2 allocs/op
    BenchmarkMake/1/8-4         	20000000	        93.3 ns/op	      88 B/op	       2 allocs/op
    BenchmarkMake/1/8-4         	20000000	        92.8 ns/op	      88 B/op	       2 allocs/op
    BenchmarkMake/1/8-4         	20000000	        92.6 ns/op	      88 B/op	       2 allocs/op
    BenchmarkMake/1/8-4         	20000000	        93.8 ns/op	      88 B/op	       2 allocs/op
    BenchmarkMake/1/16-4        	10000000	       118 ns/op	     152 B/op	       2 allocs/op
    BenchmarkMake/1/16-4        	10000000	       118 ns/op	     152 B/op	       2 allocs/op
    BenchmarkMake/1/16-4        	10000000	       119 ns/op	     152 B/op	       2 allocs/op
    BenchmarkMake/1/16-4        	10000000	       119 ns/op	     152 B/op	       2 allocs/op
    BenchmarkMake/1/32-4        	10000000	       174 ns/op	     296 B/op	       2 allocs/op
    BenchmarkMake/1/32-4        	10000000	       173 ns/op	     296 B/op	       2 allocs/op
    BenchmarkMake/1/32-4        	10000000	       172 ns/op	     296 B/op	       2 allocs/op
    BenchmarkMake/1/32-4        	10000000	       181 ns/op	     296 B/op	       2 allocs/op
    BenchmarkMake/1/64-4        	 5000000	       277 ns/op	     584 B/op	       2 allocs/op
    BenchmarkMake/1/64-4        	 5000000	       277 ns/op	     584 B/op	       2 allocs/op
    BenchmarkMake/1/64-4        	 5000000	       278 ns/op	     584 B/op	       2 allocs/op
    BenchmarkMake/1/64-4        	 5000000	       278 ns/op	     584 B/op	       2 allocs/op
    BenchmarkMake/1/128-4       	 3000000	       472 ns/op	    1160 B/op	       2 allocs/op
    BenchmarkMake/1/128-4       	 3000000	       476 ns/op	    1160 B/op	       2 allocs/op
    BenchmarkMake/1/128-4       	 3000000	       474 ns/op	    1160 B/op	       2 allocs/op
    BenchmarkMake/1/128-4       	 3000000	       473 ns/op	    1160 B/op	       2 allocs/op
    BenchmarkMake/1/256-4       	 2000000	       813 ns/op	    2312 B/op	       2 allocs/op
    BenchmarkMake/1/256-4       	 2000000	       812 ns/op	    2312 B/op	       2 allocs/op
    BenchmarkMake/1/256-4       	 2000000	       813 ns/op	    2312 B/op	       2 allocs/op
    BenchmarkMake/1/256-4       	 2000000	       819 ns/op	    2312 B/op	       2 allocs/op
    BenchmarkMake/1/512-4       	 1000000	      1586 ns/op	    4872 B/op	       2 allocs/op
    BenchmarkMake/1/512-4       	 1000000	      1602 ns/op	    4872 B/op	       2 allocs/op
    BenchmarkMake/1/512-4       	 1000000	      1584 ns/op	    4872 B/op	       2 allocs/op
    BenchmarkMake/1/512-4       	 1000000	      1591 ns/op	    4872 B/op	       2 allocs/op
    BenchmarkMake/1/1024-4      	  500000	      3011 ns/op	    9480 B/op	       2 allocs/op
    BenchmarkMake/1/1024-4      	  500000	      3000 ns/op	    9480 B/op	       2 allocs/op
    BenchmarkMake/1/1024-4      	  500000	      3006 ns/op	    9480 B/op	       2 allocs/op
    BenchmarkMake/1/1024-4      	  500000	      3011 ns/op	    9480 B/op	       2 allocs/op
    BenchmarkMake/2/1-4         	20000000	        85.0 ns/op	      48 B/op	       2 allocs/op
    BenchmarkMake/2/1-4         	20000000	        86.2 ns/op	      48 B/op	       2 allocs/op
    BenchmarkMake/2/1-4         	20000000	        84.4 ns/op	      48 B/op	       2 allocs/op
    BenchmarkMake/2/1-4         	20000000	        84.7 ns/op	      48 B/op	       2 allocs/op
    BenchmarkMake/2/2-4         	20000000	        86.9 ns/op	      48 B/op	       2 allocs/op
    BenchmarkMake/2/2-4         	20000000	        86.0 ns/op	      48 B/op	       2 allocs/op
    BenchmarkMake/2/2-4         	20000000	        86.7 ns/op	      48 B/op	       2 allocs/op
    BenchmarkMake/2/2-4         	20000000	        87.8 ns/op	      48 B/op	       2 allocs/op
    BenchmarkMake/2/4-4         	20000000	        92.3 ns/op	      64 B/op	       2 allocs/op
    BenchmarkMake/2/4-4         	20000000	        92.4 ns/op	      64 B/op	       2 allocs/op
    BenchmarkMake/2/4-4         	20000000	        91.8 ns/op	      64 B/op	       2 allocs/op
    BenchmarkMake/2/4-4         	20000000	        91.3 ns/op	      64 B/op	       2 allocs/op
    BenchmarkMake/2/8-4         	20000000	       103 ns/op	      96 B/op	       2 allocs/op
    BenchmarkMake/2/8-4         	20000000	       102 ns/op	      96 B/op	       2 allocs/op
    BenchmarkMake/2/8-4         	20000000	       103 ns/op	      96 B/op	       2 allocs/op
    BenchmarkMake/2/8-4         	10000000	       102 ns/op	      96 B/op	       2 allocs/op
    BenchmarkMake/2/16-4        	10000000	       129 ns/op	     160 B/op	       2 allocs/op
    BenchmarkMake/2/16-4        	10000000	       129 ns/op	     160 B/op	       2 allocs/op
    BenchmarkMake/2/16-4        	10000000	       130 ns/op	     160 B/op	       2 allocs/op
    BenchmarkMake/2/16-4        	10000000	       129 ns/op	     160 B/op	       2 allocs/op
    BenchmarkMake/2/32-4        	10000000	       183 ns/op	     304 B/op	       2 allocs/op
    BenchmarkMake/2/32-4        	10000000	       184 ns/op	     304 B/op	       2 allocs/op
    BenchmarkMake/2/32-4        	10000000	       183 ns/op	     304 B/op	       2 allocs/op
    BenchmarkMake/2/32-4        	10000000	       185 ns/op	     304 B/op	       2 allocs/op
    BenchmarkMake/2/64-4        	 5000000	       289 ns/op	     592 B/op	       2 allocs/op
    BenchmarkMake/2/64-4        	 5000000	       292 ns/op	     592 B/op	       2 allocs/op
    BenchmarkMake/2/64-4        	 5000000	       289 ns/op	     592 B/op	       2 allocs/op
    BenchmarkMake/2/64-4        	 5000000	       288 ns/op	     592 B/op	       2 allocs/op
    BenchmarkMake/2/128-4       	 3000000	       485 ns/op	    1168 B/op	       2 allocs/op
    BenchmarkMake/2/128-4       	 3000000	       488 ns/op	    1168 B/op	       2 allocs/op
    BenchmarkMake/2/128-4       	 3000000	       486 ns/op	    1168 B/op	       2 allocs/op
    BenchmarkMake/2/128-4       	 3000000	       488 ns/op	    1168 B/op	       2 allocs/op
    BenchmarkMake/2/256-4       	 2000000	       815 ns/op	    2320 B/op	       2 allocs/op
    BenchmarkMake/2/256-4       	 2000000	       824 ns/op	    2320 B/op	       2 allocs/op
    BenchmarkMake/2/256-4       	 2000000	       819 ns/op	    2320 B/op	       2 allocs/op
    BenchmarkMake/2/256-4       	 2000000	       822 ns/op	    2320 B/op	       2 allocs/op
    BenchmarkMake/2/512-4       	 1000000	      1604 ns/op	    4880 B/op	       2 allocs/op
    BenchmarkMake/2/512-4       	 1000000	      1608 ns/op	    4880 B/op	       2 allocs/op
    BenchmarkMake/2/512-4       	 1000000	      1606 ns/op	    4880 B/op	       2 allocs/op
    BenchmarkMake/2/512-4       	 1000000	      1600 ns/op	    4880 B/op	       2 allocs/op
    BenchmarkMake/2/1024-4      	  500000	      3040 ns/op	    9488 B/op	       2 allocs/op
    BenchmarkMake/2/1024-4      	  500000	      3028 ns/op	    9488 B/op	       2 allocs/op
    BenchmarkMake/2/1024-4      	  500000	      2961 ns/op	    9488 B/op	       2 allocs/op
    BenchmarkMake/2/1024-4      	  500000	      3115 ns/op	    9488 B/op	       2 allocs/op
    BenchmarkMake/4/1-4         	20000000	        93.5 ns/op	      80 B/op	       2 allocs/op
    BenchmarkMake/4/1-4         	20000000	        92.9 ns/op	      80 B/op	       2 allocs/op
    BenchmarkMake/4/1-4         	20000000	        93.7 ns/op	      80 B/op	       2 allocs/op
    BenchmarkMake/4/1-4         	20000000	        93.5 ns/op	      80 B/op	       2 allocs/op
    BenchmarkMake/4/2-4         	20000000	        94.4 ns/op	      80 B/op	       2 allocs/op
    BenchmarkMake/4/2-4         	20000000	        94.9 ns/op	      80 B/op	       2 allocs/op
    BenchmarkMake/4/2-4         	20000000	        94.6 ns/op	      80 B/op	       2 allocs/op
    BenchmarkMake/4/2-4         	20000000	        95.1 ns/op	      80 B/op	       2 allocs/op
    BenchmarkMake/4/4-4         	20000000	        99.9 ns/op	      96 B/op	       2 allocs/op
    BenchmarkMake/4/4-4         	20000000	       100 ns/op	      96 B/op	       2 allocs/op
    BenchmarkMake/4/4-4         	20000000	       100 ns/op	      96 B/op	       2 allocs/op
    BenchmarkMake/4/4-4         	20000000	       101 ns/op	      96 B/op	       2 allocs/op
    BenchmarkMake/4/8-4         	10000000	       112 ns/op	     128 B/op	       2 allocs/op
    BenchmarkMake/4/8-4         	10000000	       112 ns/op	     128 B/op	       2 allocs/op
    BenchmarkMake/4/8-4         	10000000	       112 ns/op	     128 B/op	       2 allocs/op
    BenchmarkMake/4/8-4         	10000000	       111 ns/op	     128 B/op	       2 allocs/op
    BenchmarkMake/4/16-4        	10000000	       138 ns/op	     192 B/op	       2 allocs/op
    BenchmarkMake/4/16-4        	10000000	       138 ns/op	     192 B/op	       2 allocs/op
    BenchmarkMake/4/16-4        	10000000	       137 ns/op	     192 B/op	       2 allocs/op
    BenchmarkMake/4/16-4        	10000000	       138 ns/op	     192 B/op	       2 allocs/op
    BenchmarkMake/4/32-4        	10000000	       189 ns/op	     320 B/op	       2 allocs/op
    BenchmarkMake/4/32-4        	10000000	       190 ns/op	     320 B/op	       2 allocs/op
    BenchmarkMake/4/32-4        	10000000	       190 ns/op	     320 B/op	       2 allocs/op
    BenchmarkMake/4/32-4        	10000000	       190 ns/op	     320 B/op	       2 allocs/op
    BenchmarkMake/4/64-4        	 5000000	       296 ns/op	     608 B/op	       2 allocs/op
    BenchmarkMake/4/64-4        	 5000000	       293 ns/op	     608 B/op	       2 allocs/op
    BenchmarkMake/4/64-4        	 5000000	       294 ns/op	     608 B/op	       2 allocs/op
    BenchmarkMake/4/64-4        	 5000000	       298 ns/op	     608 B/op	       2 allocs/op
    BenchmarkMake/4/128-4       	 3000000	       490 ns/op	    1184 B/op	       2 allocs/op
    BenchmarkMake/4/128-4       	 3000000	       489 ns/op	    1184 B/op	       2 allocs/op
    BenchmarkMake/4/128-4       	 3000000	       488 ns/op	    1184 B/op	       2 allocs/op
    BenchmarkMake/4/128-4       	 3000000	       489 ns/op	    1184 B/op	       2 allocs/op
    BenchmarkMake/4/256-4       	 2000000	       829 ns/op	    2336 B/op	       2 allocs/op
    BenchmarkMake/4/256-4       	 2000000	       826 ns/op	    2336 B/op	       2 allocs/op
    BenchmarkMake/4/256-4       	 2000000	       829 ns/op	    2336 B/op	       2 allocs/op
    BenchmarkMake/4/256-4       	 2000000	       836 ns/op	    2336 B/op	       2 allocs/op
    BenchmarkMake/4/512-4       	 1000000	      1614 ns/op	    4896 B/op	       2 allocs/op
    BenchmarkMake/4/512-4       	 1000000	      1611 ns/op	    4896 B/op	       2 allocs/op
    BenchmarkMake/4/512-4       	 1000000	      1602 ns/op	    4896 B/op	       2 allocs/op
    BenchmarkMake/4/512-4       	 1000000	      1695 ns/op	    4896 B/op	       2 allocs/op
    BenchmarkMake/4/1024-4      	  500000	      2985 ns/op	    9504 B/op	       2 allocs/op
    BenchmarkMake/4/1024-4      	  500000	      2999 ns/op	    9504 B/op	       2 allocs/op
    BenchmarkMake/4/1024-4      	  500000	      2993 ns/op	    9504 B/op	       2 allocs/op
    BenchmarkMake/4/1024-4      	  500000	      3037 ns/op	    9504 B/op	       2 allocs/op
    BenchmarkMake/8/1-4         	10000000	       108 ns/op	     144 B/op	       2 allocs/op
    BenchmarkMake/8/1-4         	10000000	       111 ns/op	     144 B/op	       2 allocs/op
    BenchmarkMake/8/1-4         	10000000	       109 ns/op	     144 B/op	       2 allocs/op
    BenchmarkMake/8/1-4         	10000000	       108 ns/op	     144 B/op	       2 allocs/op
    BenchmarkMake/8/2-4         	10000000	       111 ns/op	     144 B/op	       2 allocs/op
    BenchmarkMake/8/2-4         	10000000	       111 ns/op	     144 B/op	       2 allocs/op
    BenchmarkMake/8/2-4         	10000000	       112 ns/op	     144 B/op	       2 allocs/op
    BenchmarkMake/8/2-4         	10000000	       110 ns/op	     144 B/op	       2 allocs/op
    BenchmarkMake/8/4-4         	10000000	       116 ns/op	     160 B/op	       2 allocs/op
    BenchmarkMake/8/4-4         	10000000	       116 ns/op	     160 B/op	       2 allocs/op
    BenchmarkMake/8/4-4         	10000000	       117 ns/op	     160 B/op	       2 allocs/op
    BenchmarkMake/8/4-4         	10000000	       116 ns/op	     160 B/op	       2 allocs/op
    BenchmarkMake/8/8-4         	10000000	       128 ns/op	     192 B/op	       2 allocs/op
    BenchmarkMake/8/8-4         	10000000	       128 ns/op	     192 B/op	       2 allocs/op
    BenchmarkMake/8/8-4         	10000000	       128 ns/op	     192 B/op	       2 allocs/op
    BenchmarkMake/8/8-4         	10000000	       128 ns/op	     192 B/op	       2 allocs/op
    BenchmarkMake/8/16-4        	10000000	       151 ns/op	     256 B/op	       2 allocs/op
    BenchmarkMake/8/16-4        	10000000	       152 ns/op	     256 B/op	       2 allocs/op
    BenchmarkMake/8/16-4        	10000000	       151 ns/op	     256 B/op	       2 allocs/op
    BenchmarkMake/8/16-4        	10000000	       151 ns/op	     256 B/op	       2 allocs/op
    BenchmarkMake/8/32-4        	10000000	       207 ns/op	     384 B/op	       2 allocs/op
    BenchmarkMake/8/32-4        	10000000	       208 ns/op	     384 B/op	       2 allocs/op
    BenchmarkMake/8/32-4        	10000000	       206 ns/op	     384 B/op	       2 allocs/op
    BenchmarkMake/8/32-4        	10000000	       206 ns/op	     384 B/op	       2 allocs/op
    BenchmarkMake/8/64-4        	 5000000	       308 ns/op	     640 B/op	       2 allocs/op
    BenchmarkMake/8/64-4        	 5000000	       305 ns/op	     640 B/op	       2 allocs/op
    BenchmarkMake/8/64-4        	 5000000	       307 ns/op	     640 B/op	       2 allocs/op
    BenchmarkMake/8/64-4        	 5000000	       307 ns/op	     640 B/op	       2 allocs/op
    BenchmarkMake/8/128-4       	 3000000	       502 ns/op	    1216 B/op	       2 allocs/op
    BenchmarkMake/8/128-4       	 3000000	       501 ns/op	    1216 B/op	       2 allocs/op
    BenchmarkMake/8/128-4       	 3000000	       501 ns/op	    1216 B/op	       2 allocs/op
    BenchmarkMake/8/128-4       	 3000000	       502 ns/op	    1216 B/op	       2 allocs/op
    BenchmarkMake/8/256-4       	 2000000	       841 ns/op	    2368 B/op	       2 allocs/op
    BenchmarkMake/8/256-4       	 2000000	       841 ns/op	    2368 B/op	       2 allocs/op
    BenchmarkMake/8/256-4       	 2000000	       847 ns/op	    2368 B/op	       2 allocs/op
    BenchmarkMake/8/256-4       	 2000000	       847 ns/op	    2368 B/op	       2 allocs/op
    BenchmarkMake/8/512-4       	 1000000	      1663 ns/op	    4928 B/op	       2 allocs/op
    BenchmarkMake/8/512-4       	 1000000	      1760 ns/op	    4928 B/op	       2 allocs/op
    BenchmarkMake/8/512-4       	 1000000	      1615 ns/op	    4928 B/op	       2 allocs/op
    BenchmarkMake/8/512-4       	 1000000	      1621 ns/op	    4928 B/op	       2 allocs/op
    BenchmarkMake/8/1024-4      	  500000	      2986 ns/op	    9536 B/op	       2 allocs/op
    BenchmarkMake/8/1024-4      	  500000	      2999 ns/op	    9536 B/op	       2 allocs/op
    BenchmarkMake/8/1024-4      	  500000	      3018 ns/op	    9536 B/op	       2 allocs/op
    BenchmarkMake/8/1024-4      	  500000	      3003 ns/op	    9536 B/op	       2 allocs/op
    BenchmarkMake/16/1-4        	10000000	       142 ns/op	     272 B/op	       2 allocs/op
    BenchmarkMake/16/1-4        	10000000	       145 ns/op	     272 B/op	       2 allocs/op
    BenchmarkMake/16/1-4        	10000000	       143 ns/op	     272 B/op	       2 allocs/op
    BenchmarkMake/16/1-4        	10000000	       144 ns/op	     272 B/op	       2 allocs/op
    BenchmarkMake/16/2-4        	10000000	       143 ns/op	     272 B/op	       2 allocs/op
    BenchmarkMake/16/2-4        	10000000	       145 ns/op	     272 B/op	       2 allocs/op
    BenchmarkMake/16/2-4        	10000000	       145 ns/op	     272 B/op	       2 allocs/op
    BenchmarkMake/16/2-4        	10000000	       144 ns/op	     272 B/op	       2 allocs/op
    BenchmarkMake/16/4-4        	10000000	       152 ns/op	     288 B/op	       2 allocs/op
    BenchmarkMake/16/4-4        	10000000	       152 ns/op	     288 B/op	       2 allocs/op
    BenchmarkMake/16/4-4        	10000000	       150 ns/op	     288 B/op	       2 allocs/op
    BenchmarkMake/16/4-4        	10000000	       150 ns/op	     288 B/op	       2 allocs/op
    BenchmarkMake/16/8-4        	10000000	       161 ns/op	     320 B/op	       2 allocs/op
    BenchmarkMake/16/8-4        	10000000	       165 ns/op	     320 B/op	       2 allocs/op
    BenchmarkMake/16/8-4        	10000000	       163 ns/op	     320 B/op	       2 allocs/op
    BenchmarkMake/16/8-4        	10000000	       162 ns/op	     320 B/op	       2 allocs/op
    BenchmarkMake/16/16-4       	10000000	       188 ns/op	     384 B/op	       2 allocs/op
    BenchmarkMake/16/16-4       	10000000	       185 ns/op	     384 B/op	       2 allocs/op
    BenchmarkMake/16/16-4       	10000000	       184 ns/op	     384 B/op	       2 allocs/op
    BenchmarkMake/16/16-4       	10000000	       185 ns/op	     384 B/op	       2 allocs/op
    BenchmarkMake/16/32-4       	 5000000	       236 ns/op	     512 B/op	       2 allocs/op
    BenchmarkMake/16/32-4       	 5000000	       237 ns/op	     512 B/op	       2 allocs/op
    BenchmarkMake/16/32-4       	10000000	       237 ns/op	     512 B/op	       2 allocs/op
    BenchmarkMake/16/32-4       	10000000	       236 ns/op	     512 B/op	       2 allocs/op
    BenchmarkMake/16/64-4       	 5000000	       345 ns/op	     768 B/op	       2 allocs/op
    BenchmarkMake/16/64-4       	 5000000	       349 ns/op	     768 B/op	       2 allocs/op
    BenchmarkMake/16/64-4       	 5000000	       357 ns/op	     768 B/op	       2 allocs/op
    BenchmarkMake/16/64-4       	 5000000	       345 ns/op	     768 B/op	       2 allocs/op
    BenchmarkMake/16/128-4      	 3000000	       543 ns/op	    1280 B/op	       2 allocs/op
    BenchmarkMake/16/128-4      	 3000000	       525 ns/op	    1280 B/op	       2 allocs/op
    BenchmarkMake/16/128-4      	 3000000	       526 ns/op	    1280 B/op	       2 allocs/op
    BenchmarkMake/16/128-4      	 3000000	       526 ns/op	    1280 B/op	       2 allocs/op
    BenchmarkMake/16/256-4      	 2000000	       887 ns/op	    2432 B/op	       2 allocs/op
    BenchmarkMake/16/256-4      	 2000000	       886 ns/op	    2432 B/op	       2 allocs/op
    BenchmarkMake/16/256-4      	 2000000	      1933 ns/op	    2432 B/op	       2 allocs/op
    BenchmarkMake/16/256-4      	 1000000	      1019 ns/op	    2432 B/op	       2 allocs/op
    BenchmarkMake/16/512-4      	 1000000	      1825 ns/op	    4992 B/op	       2 allocs/op
    BenchmarkMake/16/512-4      	 1000000	      1716 ns/op	    4992 B/op	       2 allocs/op
    BenchmarkMake/16/512-4      	 1000000	      1813 ns/op	    4992 B/op	       2 allocs/op
    BenchmarkMake/16/512-4      	 1000000	      1765 ns/op	    4992 B/op	       2 allocs/op
    BenchmarkMake/16/1024-4     	  500000	      3064 ns/op	    9600 B/op	       2 allocs/op
    BenchmarkMake/16/1024-4     	  500000	      3076 ns/op	    9600 B/op	       2 allocs/op
    BenchmarkMake/16/1024-4     	  500000	      3163 ns/op	    9600 B/op	       2 allocs/op
    BenchmarkMake/16/1024-4     	  500000	      3140 ns/op	    9600 B/op	       2 allocs/op
    BenchmarkMake/32/1-4        	10000000	       215 ns/op	     544 B/op	       2 allocs/op
    BenchmarkMake/32/1-4        	10000000	       214 ns/op	     544 B/op	       2 allocs/op
    BenchmarkMake/32/1-4        	10000000	       214 ns/op	     544 B/op	       2 allocs/op
    BenchmarkMake/32/1-4        	10000000	       214 ns/op	     544 B/op	       2 allocs/op
    BenchmarkMake/32/2-4        	10000000	       215 ns/op	     544 B/op	       2 allocs/op
    BenchmarkMake/32/2-4        	10000000	       215 ns/op	     544 B/op	       2 allocs/op
    BenchmarkMake/32/2-4        	10000000	       215 ns/op	     544 B/op	       2 allocs/op
    BenchmarkMake/32/2-4        	10000000	       214 ns/op	     544 B/op	       2 allocs/op
    BenchmarkMake/32/4-4        	10000000	       218 ns/op	     544 B/op	       2 allocs/op
    BenchmarkMake/32/4-4        	10000000	       216 ns/op	     544 B/op	       2 allocs/op
    BenchmarkMake/32/4-4        	10000000	       216 ns/op	     544 B/op	       2 allocs/op
    BenchmarkMake/32/4-4        	 5000000	       216 ns/op	     544 B/op	       2 allocs/op
    BenchmarkMake/32/8-4        	 5000000	       229 ns/op	     576 B/op	       2 allocs/op
    BenchmarkMake/32/8-4        	 5000000	       229 ns/op	     576 B/op	       2 allocs/op
    BenchmarkMake/32/8-4        	 5000000	       232 ns/op	     576 B/op	       2 allocs/op
    BenchmarkMake/32/8-4        	 5000000	       229 ns/op	     576 B/op	       2 allocs/op
    BenchmarkMake/32/16-4       	 5000000	       252 ns/op	     640 B/op	       2 allocs/op
    BenchmarkMake/32/16-4       	 5000000	       252 ns/op	     640 B/op	       2 allocs/op
    BenchmarkMake/32/16-4       	 5000000	       254 ns/op	     640 B/op	       2 allocs/op
    BenchmarkMake/32/16-4       	 5000000	       254 ns/op	     640 B/op	       2 allocs/op
    BenchmarkMake/32/32-4       	 5000000	       310 ns/op	     768 B/op	       2 allocs/op
    BenchmarkMake/32/32-4       	 5000000	       303 ns/op	     768 B/op	       2 allocs/op
    BenchmarkMake/32/32-4       	 5000000	       302 ns/op	     768 B/op	       2 allocs/op
    BenchmarkMake/32/32-4       	 5000000	       304 ns/op	     768 B/op	       2 allocs/op
    BenchmarkMake/32/64-4       	 3000000	       419 ns/op	    1024 B/op	       2 allocs/op
    BenchmarkMake/32/64-4       	 3000000	       423 ns/op	    1024 B/op	       2 allocs/op
    BenchmarkMake/32/64-4       	 3000000	       421 ns/op	    1024 B/op	       2 allocs/op
    BenchmarkMake/32/64-4       	 3000000	       428 ns/op	    1024 B/op	       2 allocs/op
    BenchmarkMake/32/128-4      	 2000000	       635 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkMake/32/128-4      	 2000000	       643 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkMake/32/128-4      	 2000000	       637 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkMake/32/128-4      	 2000000	       637 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkMake/32/256-4      	 2000000	       968 ns/op	    2560 B/op	       2 allocs/op
    BenchmarkMake/32/256-4      	 2000000	       971 ns/op	    2560 B/op	       2 allocs/op
    BenchmarkMake/32/256-4      	 2000000	       960 ns/op	    2560 B/op	       2 allocs/op
    BenchmarkMake/32/256-4      	 2000000	       974 ns/op	    2560 B/op	       2 allocs/op
    BenchmarkMake/32/512-4      	 1000000	      1805 ns/op	    5120 B/op	       2 allocs/op
    BenchmarkMake/32/512-4      	 1000000	      1801 ns/op	    5120 B/op	       2 allocs/op
    BenchmarkMake/32/512-4      	 1000000	      1822 ns/op	    5120 B/op	       2 allocs/op
    BenchmarkMake/32/512-4      	 1000000	      1817 ns/op	    5120 B/op	       2 allocs/op
    BenchmarkMake/32/1024-4     	  500000	      3284 ns/op	    9728 B/op	       2 allocs/op
    BenchmarkMake/32/1024-4     	  500000	      3272 ns/op	    9728 B/op	       2 allocs/op
    BenchmarkMake/32/1024-4     	  500000	      3228 ns/op	    9728 B/op	       2 allocs/op
    BenchmarkMake/32/1024-4     	  500000	      3282 ns/op	    9728 B/op	       2 allocs/op
    BenchmarkMake/64/1-4        	 5000000	       361 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/1-4        	 5000000	       363 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/1-4        	 5000000	       360 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/1-4        	 5000000	       362 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/2-4        	 5000000	       363 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/2-4        	 5000000	       363 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/2-4        	 3000000	       396 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/2-4        	 5000000	       361 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/4-4        	 5000000	       368 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/4-4        	 5000000	       368 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/4-4        	 3000000	       365 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/4-4        	 5000000	       366 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/8-4        	 5000000	       370 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/8-4        	 5000000	       376 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/8-4        	 5000000	       363 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/8-4        	 5000000	       368 ns/op	    1088 B/op	       2 allocs/op
    BenchmarkMake/64/16-4       	 3000000	       400 ns/op	    1152 B/op	       2 allocs/op
    BenchmarkMake/64/16-4       	 3000000	       397 ns/op	    1152 B/op	       2 allocs/op
    BenchmarkMake/64/16-4       	 3000000	       401 ns/op	    1152 B/op	       2 allocs/op
    BenchmarkMake/64/16-4       	 3000000	       401 ns/op	    1152 B/op	       2 allocs/op
    BenchmarkMake/64/32-4       	 3000000	       460 ns/op	    1280 B/op	       2 allocs/op
    BenchmarkMake/64/32-4       	 3000000	       463 ns/op	    1280 B/op	       2 allocs/op
    BenchmarkMake/64/32-4       	 3000000	       459 ns/op	    1280 B/op	       2 allocs/op
    BenchmarkMake/64/32-4       	 3000000	       463 ns/op	    1280 B/op	       2 allocs/op
    BenchmarkMake/64/64-4       	 3000000	       545 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkMake/64/64-4       	 3000000	       543 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkMake/64/64-4       	 3000000	       540 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkMake/64/64-4       	 3000000	       535 ns/op	    1536 B/op	       2 allocs/op
    BenchmarkMake/64/128-4      	 2000000	       751 ns/op	    2048 B/op	       2 allocs/op
    BenchmarkMake/64/128-4      	 2000000	       749 ns/op	    2048 B/op	       2 allocs/op
    BenchmarkMake/64/128-4      	 2000000	       756 ns/op	    2048 B/op	       2 allocs/op
    BenchmarkMake/64/128-4      	 2000000	       750 ns/op	    2048 B/op	       2 allocs/op
    BenchmarkMake/64/256-4      	 1000000	      1163 ns/op	    3200 B/op	       2 allocs/op
    BenchmarkMake/64/256-4      	 1000000	      1165 ns/op	    3200 B/op	       2 allocs/op
    BenchmarkMake/64/256-4      	 1000000	      1162 ns/op	    3200 B/op	       2 allocs/op
    BenchmarkMake/64/256-4      	 1000000	      1167 ns/op	    3200 B/op	       2 allocs/op
    BenchmarkMake/64/512-4      	 1000000	      1797 ns/op	    5376 B/op	       2 allocs/op
    BenchmarkMake/64/512-4      	 1000000	      1783 ns/op	    5376 B/op	       2 allocs/op
    BenchmarkMake/64/512-4      	 1000000	      1783 ns/op	    5376 B/op	       2 allocs/op
    BenchmarkMake/64/512-4      	 1000000	      1878 ns/op	    5376 B/op	       2 allocs/op
    BenchmarkMake/64/1024-4     	  500000	      3189 ns/op	    9984 B/op	       2 allocs/op
    BenchmarkMake/64/1024-4     	  500000	      3168 ns/op	    9984 B/op	       2 allocs/op
    BenchmarkMake/64/1024-4     	  500000	      3137 ns/op	    9984 B/op	       2 allocs/op
    BenchmarkMake/64/1024-4     	  500000	      3192 ns/op	    9984 B/op	       2 allocs/op
    BenchmarkMake/128/1-4       	 2000000	       659 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/1-4       	 2000000	       644 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/1-4       	 2000000	       646 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/1-4       	 2000000	       652 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/2-4       	 2000000	       647 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/2-4       	 2000000	       648 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/2-4       	 2000000	       647 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/2-4       	 2000000	       649 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/4-4       	 2000000	       657 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/4-4       	 2000000	       649 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/4-4       	 2000000	       649 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/4-4       	 2000000	       653 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/8-4       	 2000000	       653 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/8-4       	 2000000	       661 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/8-4       	 2000000	       659 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/8-4       	 2000000	       660 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/16-4      	 2000000	       667 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/16-4      	 2000000	       671 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/16-4      	 2000000	       683 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/16-4      	 2000000	       662 ns/op	    2176 B/op	       2 allocs/op
    BenchmarkMake/128/32-4      	 2000000	       738 ns/op	    2304 B/op	       2 allocs/op
    BenchmarkMake/128/32-4      	 2000000	       736 ns/op	    2304 B/op	       2 allocs/op
    BenchmarkMake/128/32-4      	 2000000	       741 ns/op	    2304 B/op	       2 allocs/op
    BenchmarkMake/128/32-4      	 2000000	       741 ns/op	    2304 B/op	       2 allocs/op
    BenchmarkMake/128/64-4      	 2000000	       848 ns/op	    2560 B/op	       2 allocs/op
    BenchmarkMake/128/64-4      	 2000000	       853 ns/op	    2560 B/op	       2 allocs/op
    BenchmarkMake/128/64-4      	 2000000	       849 ns/op	    2560 B/op	       2 allocs/op
    BenchmarkMake/128/64-4      	 2000000	       880 ns/op	    2560 B/op	       2 allocs/op
    BenchmarkMake/128/128-4     	 1000000	      1097 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkMake/128/128-4     	 1000000	      1038 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkMake/128/128-4     	 1000000	      1022 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkMake/128/128-4     	 1000000	      1031 ns/op	    3072 B/op	       2 allocs/op
    BenchmarkMake/128/256-4     	 1000000	      1302 ns/op	    4096 B/op	       2 allocs/op
    BenchmarkMake/128/256-4     	 1000000	      1318 ns/op	    4096 B/op	       2 allocs/op
    BenchmarkMake/128/256-4     	 1000000	      1304 ns/op	    4096 B/op	       2 allocs/op
    BenchmarkMake/128/256-4     	 1000000	      1308 ns/op	    4096 B/op	       2 allocs/op
    BenchmarkMake/128/512-4     	  500000	      2183 ns/op	    6400 B/op	       2 allocs/op
    BenchmarkMake/128/512-4     	 1000000	      2096 ns/op	    6400 B/op	       2 allocs/op
    BenchmarkMake/128/512-4     	 1000000	      2103 ns/op	    6400 B/op	       2 allocs/op
    BenchmarkMake/128/512-4     	 1000000	      2093 ns/op	    6400 B/op	       2 allocs/op
    BenchmarkMake/128/1024-4    	  500000	      3383 ns/op	   10496 B/op	       2 allocs/op
    BenchmarkMake/128/1024-4    	  500000	      3373 ns/op	   10496 B/op	       2 allocs/op
    BenchmarkMake/128/1024-4    	  500000	      3528 ns/op	   10496 B/op	       2 allocs/op
    BenchmarkMake/128/1024-4    	  500000	      3374 ns/op	   10496 B/op	       2 allocs/op
    BenchmarkMake/256/1-4       	 1000000	      1195 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/1-4       	 1000000	      1187 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/1-4       	 1000000	      1195 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/1-4       	 1000000	      1188 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/2-4       	 1000000	      1235 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/2-4       	 1000000	      1237 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/2-4       	 1000000	      1193 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/2-4       	 1000000	      1198 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/4-4       	 1000000	      1185 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/4-4       	 1000000	      1218 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/4-4       	 1000000	      1203 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/4-4       	 1000000	      1196 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/8-4       	 1000000	      1182 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/8-4       	 1000000	      1205 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/8-4       	 1000000	      1207 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/8-4       	 1000000	      1235 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/16-4      	 1000000	      1200 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/16-4      	 1000000	      1208 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/16-4      	 1000000	      1217 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/16-4      	 1000000	      1202 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/32-4      	 1000000	      1236 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/32-4      	 1000000	      1228 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/32-4      	 1000000	      1242 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/32-4      	 1000000	      1230 ns/op	    4352 B/op	       2 allocs/op
    BenchmarkMake/256/64-4      	 1000000	      1432 ns/op	    4736 B/op	       2 allocs/op
    BenchmarkMake/256/64-4      	 1000000	      1463 ns/op	    4736 B/op	       2 allocs/op
    BenchmarkMake/256/64-4      	 1000000	      1428 ns/op	    4736 B/op	       2 allocs/op
    BenchmarkMake/256/64-4      	 1000000	      1422 ns/op	    4736 B/op	       2 allocs/op
    BenchmarkMake/256/128-4     	 1000000	      1440 ns/op	    5120 B/op	       2 allocs/op
    BenchmarkMake/256/128-4     	 1000000	      1450 ns/op	    5120 B/op	       2 allocs/op
    BenchmarkMake/256/128-4     	 1000000	      1449 ns/op	    5120 B/op	       2 allocs/op
    BenchmarkMake/256/128-4     	 1000000	      1457 ns/op	    5120 B/op	       2 allocs/op
    BenchmarkMake/256/256-4     	 1000000	      1968 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkMake/256/256-4     	 1000000	      1960 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkMake/256/256-4     	 1000000	      1978 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkMake/256/256-4     	 1000000	      1950 ns/op	    6144 B/op	       2 allocs/op
    BenchmarkMake/256/512-4     	  500000	      2551 ns/op	    8192 B/op	       2 allocs/op
    BenchmarkMake/256/512-4     	  500000	      2528 ns/op	    8192 B/op	       2 allocs/op
    BenchmarkMake/256/512-4     	  500000	      2559 ns/op	    8192 B/op	       2 allocs/op
    BenchmarkMake/256/512-4     	  500000	      2578 ns/op	    8192 B/op	       2 allocs/op
    BenchmarkMake/256/1024-4    	  300000	      3840 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkMake/256/1024-4    	  300000	      3794 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkMake/256/1024-4    	  300000	      3907 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkMake/256/1024-4    	  500000	      3870 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkMake/512/1-4       	 1000000	      2292 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/1-4       	  500000	      2307 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/1-4       	  500000	      2282 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/1-4       	  500000	      2292 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/2-4       	 1000000	      2282 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/2-4       	  500000	      2284 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/2-4       	 1000000	      2275 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/2-4       	 1000000	      2292 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/4-4       	 1000000	      2288 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/4-4       	  500000	      2298 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/4-4       	 1000000	      2287 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/4-4       	 1000000	      2408 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/8-4       	  500000	      2319 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/8-4       	  500000	      2304 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/8-4       	 1000000	      2331 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/8-4       	  500000	      2334 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/16-4      	 1000000	      2317 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/16-4      	 1000000	      2342 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/16-4      	  500000	      2383 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/16-4      	 1000000	      2297 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/32-4      	  500000	      2357 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/32-4      	  500000	      2354 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/32-4      	  500000	      2332 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/32-4      	  500000	      2618 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/64-4      	  500000	      2421 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/64-4      	 1000000	      2409 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/64-4      	  500000	      2435 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/64-4      	 1000000	      2402 ns/op	    8960 B/op	       2 allocs/op
    BenchmarkMake/512/128-4     	  500000	      2640 ns/op	    9472 B/op	       2 allocs/op
    BenchmarkMake/512/128-4     	  500000	      2650 ns/op	    9472 B/op	       2 allocs/op
    BenchmarkMake/512/128-4     	  500000	      2642 ns/op	    9472 B/op	       2 allocs/op
    BenchmarkMake/512/128-4     	  500000	      2620 ns/op	    9472 B/op	       2 allocs/op
    BenchmarkMake/512/256-4     	  500000	      2813 ns/op	   10240 B/op	       2 allocs/op
    BenchmarkMake/512/256-4     	  500000	      2891 ns/op	   10240 B/op	       2 allocs/op
    BenchmarkMake/512/256-4     	  500000	      2873 ns/op	   10240 B/op	       2 allocs/op
    BenchmarkMake/512/256-4     	  500000	      2903 ns/op	   10240 B/op	       2 allocs/op
    BenchmarkMake/512/512-4     	  300000	      3828 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkMake/512/512-4     	  300000	      3804 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkMake/512/512-4     	  300000	      3805 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkMake/512/512-4     	  500000	      3792 ns/op	   12288 B/op	       2 allocs/op
    BenchmarkMake/512/1024-4    	  200000	      5015 ns/op	   16384 B/op	       2 allocs/op
    BenchmarkMake/512/1024-4    	  200000	      5171 ns/op	   16384 B/op	       2 allocs/op
    BenchmarkMake/512/1024-4    	  200000	      5044 ns/op	   16384 B/op	       2 allocs/op
    BenchmarkMake/512/1024-4    	  300000	      5325 ns/op	   16384 B/op	       2 allocs/op
    BenchmarkMake/1024/1-4      	  300000	      4489 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/1-4      	  300000	      4264 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/1-4      	  300000	      4252 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/1-4      	  300000	      4361 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/2-4      	  300000	      4319 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/2-4      	  300000	      4297 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/2-4      	  300000	      4289 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/2-4      	  300000	      4326 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/4-4      	  300000	      4313 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/4-4      	  300000	      4260 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/4-4      	  300000	      4431 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/4-4      	  300000	      4312 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/8-4      	  300000	      4261 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/8-4      	  300000	      4380 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/8-4      	  300000	      4308 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/8-4      	  300000	      4274 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/16-4     	  300000	      4276 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/16-4     	  300000	      4265 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/16-4     	  300000	      4330 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/16-4     	  300000	      4271 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/32-4     	  300000	      4333 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/32-4     	  300000	      4371 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/32-4     	  300000	      4335 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/32-4     	  300000	      4392 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/64-4     	  300000	      4368 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/64-4     	  300000	      4425 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/64-4     	  300000	      4426 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/64-4     	  300000	      4384 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/128-4    	  300000	      4493 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/128-4    	  300000	      4540 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/128-4    	  300000	      4580 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/128-4    	  300000	      4486 ns/op	   17664 B/op	       2 allocs/op
    BenchmarkMake/1024/256-4    	  300000	      4864 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkMake/1024/256-4    	  300000	      4943 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkMake/1024/256-4    	  300000	      4831 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkMake/1024/256-4    	  300000	      4882 ns/op	   18432 B/op	       2 allocs/op
    BenchmarkMake/1024/512-4    	  200000	      5714 ns/op	   20480 B/op	       2 allocs/op
    BenchmarkMake/1024/512-4    	  200000	      5646 ns/op	   20480 B/op	       2 allocs/op
    BenchmarkMake/1024/512-4    	  200000	      5674 ns/op	   20480 B/op	       2 allocs/op
    BenchmarkMake/1024/512-4    	  200000	      5647 ns/op	   20480 B/op	       2 allocs/op
    BenchmarkMake/1024/1024-4   	  200000	      6909 ns/op	   24576 B/op	       2 allocs/op
    BenchmarkMake/1024/1024-4   	  200000	      6847 ns/op	   24576 B/op	       2 allocs/op
    BenchmarkMake/1024/1024-4   	  200000	      6844 ns/op	   24576 B/op	       2 allocs/op
    BenchmarkMake/1024/1024-4   	  200000	      6863 ns/op	   24576 B/op	       2 allocs/op
    PASS
    ok  	github.com/elpinal/benchmarks/append-or-make	865.291s
