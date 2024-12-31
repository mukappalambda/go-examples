# HTTP response compression

The server has two paths, `/uncompressed` and `/compressed`.

The response of `/uncompressed` is `data`, and that of `/compressed` is `<wrapped by middleware>data<wrapped by middleware>%` where the `<wrapped by middleware>` pattern is added intentionally to verify that the `compressionMiddleware` middleware does do the work.

Run:

```bash
go run main.go
```

Expected output:

```console
data
<wrapped by middleware>data<wrapped by middleware>%
```
