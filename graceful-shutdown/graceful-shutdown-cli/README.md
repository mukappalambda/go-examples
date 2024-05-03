# Graceful Shutdown

Example output:

```bash
$ go run main.go
Task completed 2024-05-03 12:16:17.947 +0000 UTC
Task completed 2024-05-03 12:16:19.948 +0000 UTC
Task completed 2024-05-03 12:16:21.949 +0000 UTC
^CShutting shutdown...
Task completed 2024-05-03 12:16:23.95 +0000 UTC
Cleaning resources... 2024-05-03 12:16:23.95 +0000 UTC
Resources have been cleaned. Bye. 2024-05-03 12:16:24.95 +0000 UTC
```
