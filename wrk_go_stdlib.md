# Wrk tests for Go stdlib net/http.

ygui@thinkpad:~/github/guiyanzhong/webperf$ wrk -c2 -d2s http://127.0.0.1:8080/
```
Running 2s test @ http://127.0.0.1:8080/
  2 threads and 2 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   155.42us  485.12us   9.85ms   97.20%
    Req/Sec    10.57k     1.18k   11.69k    88.10%
  44136 requests in 2.10s, 5.39MB read
Requests/sec:  21020.69
Transfer/sec:      2.57MB
```

ygui@thinkpad:~/github/guiyanzhong/webperf$ wrk -c10 -d2s http://127.0.0.1:8080/
```
Running 2s test @ http://127.0.0.1:8080/
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.97ms   18.88ms 174.69ms   90.51%
    Req/Sec    11.11k     1.43k   18.00k    90.24%
  45299 requests in 2.10s, 5.53MB read
Requests/sec:  21564.79
Transfer/sec:      2.63MB
```

ygui@thinkpad:~/github/guiyanzhong/webperf$ wrk -c20 -d2s http://127.0.0.1:8080/
```
Running 2s test @ http://127.0.0.1:8080/
  2 threads and 20 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.91ms   28.52ms 293.98ms   91.59%
    Req/Sec    10.86k     1.02k   12.99k    72.50%
  43188 requests in 2.00s, 5.27MB read
Requests/sec:  21577.09
Transfer/sec:      2.63MB
```

ygui@thinkpad:~/github/guiyanzhong/webperf$ wrk -c100 -d2s http://127.0.0.1:8080/
```
Running 2s test @ http://127.0.0.1:8080/
  2 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    18.66ms   38.86ms 257.93ms   89.04%
    Req/Sec    10.68k     0.96k   12.26k    60.00%
  42490 requests in 2.02s, 5.19MB read
Requests/sec:  21031.68
Transfer/sec:      2.57MB
```

ygui@thinkpad:~/github/guiyanzhong/webperf$ wrk -c1000 -d2s http://127.0.0.1:8080/
```
Running 2s test @ http://127.0.0.1:8080/
  2 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   153.92ms  351.31ms   1.81s    87.10%
    Req/Sec    10.58k     0.95k   11.67k    82.50%
  42086 requests in 2.05s, 5.14MB read
Requests/sec:  20539.68
Transfer/sec:      2.51MB
```
