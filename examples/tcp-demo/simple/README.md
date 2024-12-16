# Simple TCP server

To start the server:

```bash
go run main.go
```

To start the client:

```bash
nc localhost 8080
```

Server side:

```bash
server is listening at 8080
client from "127.0.0.1:41648" is connected.
[client@127.0.0.1:41648] 2024-12-16 10:55:51 > hi
[client@127.0.0.1:41648] 2024-12-16 10:56:03 > ...
2024/12/16 10:56:10 EOF
exit status 1
```

Client side:

```bash
hi
[server@127.0.0.1:8080] 2024-12-16 10:55:51 > haha you said: "hi"?
...
[server@127.0.0.1:8080] 2024-12-16 10:56:03 > haha you said: "..."?
^C
```
