# TCP Server

Build the binary:

```console
$ go build -race
```

NB: the `-race`flag is added here to make sure there exists no race conditions.

Start the server on the first terminal:

```console
$ go build -race
$ ./tcp-server
2025/08/23 01:04:46 Server is listening on "127.0.0.1:8080"
```

Start the clients on other terminals:

Terminal 2:

```console
$ nc localhost 8080
$ Type some words here
```

Terminal 3:

```console
$ nc localhost 8080
$ Type some words here
```

Press `Ctrl-C` to shut down the server:

```console
$ ./tcp-server
2025/08/23 01:03:39 Server is listening on "127.0.0.1:8080"
^C2025/08/23 01:03:41 Server is shutting down...
2025/08/23 01:03:41 Listener has gracefully shut down.
2025/08/23 01:03:41 Server shut down successfully.
```
