# TCP Client Set Deadline

```bash
$ go run main.go -help
Usage of /tmp/go-build203211016/b001/exe/main:
  -client.timeout duration
        delay after reading from conn and before writing to conn (default 500ms)
  -server.delay duration
        delay after reading from conn and before writing to conn (default 1s)
$ go run main.go
tcp server running at 127.0.0.1:46397
client set up the connection from 127.0.0.1:37912
client set timeout to 500ms
2024/12/27 23:10:20 client wrote to server successfully
2024/12/27 23:10:20 error reading from server: read tcp 127.0.0.1:37912->127.0.0.1:46397: i/o timeout
exit status 1
$ go run main.go -server.delay 100ms
tcp server running at 127.0.0.1:36629
client set up the connection from 127.0.0.1:60578
client set timeout to 500ms
2024/12/27 23:10:42 client wrote to server successfully
2024/12/27 23:10:42 accept tcp 127.0.0.1:36629: use of closed network connection
exit status 1
```
