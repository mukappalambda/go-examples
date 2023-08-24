# Benchmark web servers

First of all, I have the [wrk](https://github.com/wg/wrk) tool installed on my machine.

For those who don't have `wrk` yet, here are the steps to get this tool:

```bash
git clone git@github.com:wg/wrk.git
cd wrk
make
cd ..
sudo mv wrk/ /opt/wrk
echo 'export PATH=/opt/wrk:$PATH >> ~/.zshrc'
source ~/.zshrc
```

Make these web servers up and running:

```bash
# in the first terminal
go run server_http.go
# in the second terminal
go run server_fasthttp.go
# and so on
```

Then we can use `wrk` to obtain the rps metric for each server.

```bash
wrk -c100 -t12 -d 15s http://localhost:8080/
```

```
Running 15s test @ http://localhost:8080/
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   648.15us    1.46ms  64.21ms   91.63%
    Req/Sec    27.36k     3.90k   48.94k    70.08%
  4927160 requests in 15.08s, 676.64MB read
Requests/sec: 326782.70
Transfer/sec:     44.88MB
```

```bash
wrk -c100 -t12 -d 15s http://localhost:8081/
```

```
Running 15s test @ http://localhost:8081/
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   611.24us    2.18ms  61.65ms   94.68%
    Req/Sec    51.03k     7.86k   83.78k    70.53%
  9182552 requests in 15.09s, 1.42GB read
Requests/sec: 608525.87
Transfer/sec:     96.34MB
```

Future Direction:
- [ ] Benchmark other web frameworks as well.
