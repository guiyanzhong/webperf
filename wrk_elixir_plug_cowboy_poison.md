# Wrk tests for Elixir Plug+Cowboy (Returning a JSON string using Poison.encode!).

ygui@thinkpad:~$ wrk -c2 -d2s http://127.0.0.1:8080/
```
Running 2s test @ http://127.0.0.1:8080/
  2 threads and 2 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   220.43us  115.32us   2.89ms   83.21%
    Req/Sec     4.56k   594.23     8.20k    97.56%
  18592 requests in 2.10s, 3.74MB read
Requests/sec:   8853.49
Transfer/sec:      1.78MB
```

ygui@thinkpad:~$ wrk -c10 -d2s http://127.0.0.1:8080/
```
Running 2s test @ http://127.0.0.1:8080/
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.07ms    1.31ms  12.24ms   85.95%
    Req/Sec     7.41k   411.32     8.55k    82.50%
  29475 requests in 2.01s, 5.94MB read
Requests/sec:  14684.69
Transfer/sec:      2.96MB
```

ygui@thinkpad:~$ wrk -c20 -d2s http://127.0.0.1:8080/
```
Running 2s test @ http://127.0.0.1:8080/
  2 threads and 20 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.34ms    2.99ms  27.61ms   85.91%
    Req/Sec     7.39k   399.27     8.33k    72.50%
  29434 requests in 2.01s, 5.93MB read
Requests/sec:  14657.64
Transfer/sec:      2.95MB
```

ygui@thinkpad:~$ wrk -c100 -d2s http://127.0.0.1:8080/
```
Running 2s test @ http://127.0.0.1:8080/
  2 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    11.94ms   16.59ms 189.67ms   89.56%
    Req/Sec     7.04k     1.39k   11.25k    77.50%
  28006 requests in 2.01s, 5.64MB read
Requests/sec:  13945.77
Transfer/sec:      2.81MB
```

ygui@thinkpad:~$ wrk -c1000 -d2s http://127.0.0.1:8080/
```
Running 2s test @ http://127.0.0.1:8080/
  2 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    47.94ms   94.60ms   1.00s    89.76%
    Req/Sec     6.77k     5.21k   13.94k    47.50%
  26977 requests in 2.05s, 5.43MB read
Requests/sec:  13136.29
Transfer/sec:      2.64MB
ygui@thinkpad:~$ 
```
