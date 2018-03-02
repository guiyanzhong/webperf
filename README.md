## Web performance tests for different programming languages and frameworks

Here are the ApacheBench results of 10000 requests with concurrency level of 100, on Thinkpad X200, Ubuntu 14.04:

Language / Framework          |  Requests per Second  |  Time per Request  |  Longest Request  |  Transfer Rate  |
------------------------------|-----------------------|--------------------|-------------------|-----------------|
Golang / stdlib (net/http)    |  8645                 |  12 ms             |   38 ms           |   1081 KB/s     |
Elixir / Plug (Cowboy)        |  5557                 |  18 ms             |   43 ms           |   1069 KB/s     |
Elixir / Phoenix (Cowboy)     |  2002                 |  50 ms             |  107 ms           |    477 KB/s     |
Ruby / Sinatra (Puma)         |  1517                 |  66 ms             |  151 ms           |    273 KB/s     |
Groovy / Spring Boot (Tomcat) |  1617                 |  62 ms             |  260 ms           |    229 KB/s     |
