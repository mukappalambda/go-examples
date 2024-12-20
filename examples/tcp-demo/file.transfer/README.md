# File Transfer via TCP

Server side:

```console
➜  file.transfer ✗ cd server
➜  server ✗ go run main.go
server is listening at [::]:8080
2024/12/20 09:34:32 client@127.0.0.1:35516 is connected
2024/12/20 09:34:32 127.0.0.1:35516 sent 29 bytes successfully
2024/12/20 09:34:50 client@127.0.0.1:54708 is connected
2024/12/20 09:34:50 127.0.0.1:54708 sent 29 bytes successfully
```

Client side:

```console
➜  file.transfer ✗ cd client
➜  client ✗ go run main.go
2024/12/20 09:34:32 connected to the network successfully
sent file content successfully: 29 bytes sent
➜  client ✗ go run main.go
2024/12/20 09:34:50 connected to the network successfully
sent file content successfully: 29 bytes sent
```
