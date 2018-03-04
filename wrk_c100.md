# Wrk test results of concurrency level of 100, on Thinkpad X200, 8GB RAM, Ubuntu 14.04:

## Go 1.10: stdlib net/http

```
ygui@thinkpad:~$ wrk -t12 -c100 -d10s http://127.0.0.1:8080/
Running 10s test @ http://127.0.0.1:8080/
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    81.85ms  132.95ms   1.09s    84.83%
    Req/Sec     2.04k     2.60k   11.55k    86.00%
  221859 requests in 10.02s, 27.08MB read
Requests/sec:  22135.83
Transfer/sec:      2.70MB
```
## Elixir 1.6.1, Erlang/OTP 20: Plug + Cowboy

```
ygui@thinkpad:~$ wrk -t12 -c100 -d10s http://127.0.0.1:8080/
Running 10s test @ http://127.0.0.1:8080/
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    10.70ms   14.41ms 164.31ms   87.87%
    Req/Sec     1.25k   653.57    11.03k    73.19%
  149403 requests in 10.10s, 28.10MB read
Requests/sec:  14793.00
Transfer/sec:      2.78MB
```

## Elixir 1.6.1, Erlang/OTP 20: Phoenix (Cowboy) 

```
ygui@thinkpad:~$ wrk -t12 -c100 -d10s http://127.0.0.1:8080/api/
Running 10s test @ http://127.0.0.1:8080/api/
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    32.38ms   12.37ms 103.77ms   68.73%
    Req/Sec   248.07     34.66   350.00     69.50%
  29678 requests in 10.02s, 6.91MB read
Requests/sec:   2960.95
Transfer/sec:    706.06KB
```

## Ruby 2.4.0: Sinatra + Puma

```
ygui@thinkpad:~$ wrk -t12 -c100 -d10s http://127.0.0.1:4567/
Running 10s test @ http://127.0.0.1:4567/
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.28ms    8.42ms  77.27ms   74.74%
    Req/Sec   576.27    155.32     0.98k    63.33%
  17221 requests in 10.03s, 3.02MB read
Requests/sec:   1717.00
Transfer/sec:    308.52KB
```

## Java 1.8: Spring Boot + Tomcat, after warming up

```
ygui@thinkpad:~$ wrk -t12 -c100 -d10s http://127.0.0.1:8080/
Running 10s test @ http://127.0.0.1:8080/
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    12.97ms   17.78ms 232.21ms   89.15%
    Req/Sec     0.97k   426.06     4.63k    74.59%
  115656 requests in 10.10s, 13.92MB read
Requests/sec:  11453.32
Transfer/sec:      1.38MB
```

## Groovy: Spring Boot + Tomcat, after warming up

```
ygui@thinkpad:~$ wrk -t12 -c100 -d10s http://127.0.0.1:8080/
Running 10s test @ http://127.0.0.1:8080/
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    13.95ms   18.85ms 227.13ms   88.32%
    Req/Sec     0.93k   417.30     5.65k    77.46%
  110403 requests in 10.09s, 13.29MB read
Requests/sec:  10938.19
Transfer/sec:      1.32MB
```
