# Web performance tests using wrk

## Results of concurrency level of 100, on Thinkpad X200, 8GB RAM, Ubuntu 14.04:

Language / Framework          |  Requests per Second  |  Latency (ms, Avg / Stdev / Max)  |  Transfer Rate  |
------------------------------|-----------------------|-----------------------------------|-----------------|
Golang / stdlib net/http      |  22136                |  82 / 133 / 1090                  |    2.70 MB/s    |
Elixir / Plug + Cowboy        |  14793                |  11 /  14 /  164                  |    2.78 MB/s    |
Elixir / Phoenix (Cowboy)     |  2961                 |  32 /  12 /  104                  |    0.71 MB/s    |
Ruby / Sinatra + Puma         |  1717                 |   9 /   8 /   77                  |    0.31 MB/s    |
Java / Spring Boot + Tomcat   |  11453                |  13 /  18 /  232                  |    1.38 MB/s    |
Groovy / Spring Boot + Tomcat |  10938                |  14 /  19 /  227                  |    1.32 MB/s    |

[Details](wrk_c100.md)
