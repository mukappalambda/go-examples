# Validating ISS and AUD

Run the server:

```console
$ go run main.go
2024/12/26 13:54:36 127.0.0.1:57582 unauthorized
2024/12/26 13:55:00 127.0.0.1:44738 unauthorized
2024/12/26 13:55:18 127.0.0.1:55054 invalid token
2024/12/26 13:55:26 127.0.0.1:56882 token is malformed: token contains an invalid number of segments
2024/12/26 13:55:34 127.0.0.1:57046 token is malformed: could not base64 decode header: illegal base64 data at input byte 0
2024/12/26 13:56:24 iss: "example.auth.server" aud: ["example.api"]
```

Use `curl` to test the handlers:

```console
$ curl localhost:8080/data
$ curl localhost:8080/data -H "Authorization: 123"
$ curl localhost:8080/data -H "Authorization: Bearer 123"
$ curl localhost:8080/data -H "Authorization: Bearer a.b.c"
$ curl localhost:8080/token
$ curl localhost:8080/data -H "Authorization: Bearer <token-string>"
```
