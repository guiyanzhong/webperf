## Web performance tests for different programming languages and frameworks

Here are the results of 10000 requests with concurrency level of 100, on Thinkpad X200, Ubuntu 14.04:

Language / Framework   |  Requests per Second  |  Time per Request  |  Longest Request  |  Transfer Rate  |
-----------------------|-----------------------|--------------------|-------------------|-----------------|
Golang / stdlib        |  8645                 |  11.6ms            |   38ms            |   1080.64 KB/s  |
Elixir / Phoenix       |  2002                 |  49.9ms            |  107ms            |    477.12 KB/s  |
