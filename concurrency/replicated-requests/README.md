# Replicated requests

```console
$ go build
$ ./replicated-requests -h
Usage of ./replicated-requests:
  -n int
        number of replicated requests (default 1)
$ for n in 1 10 50 100; do ./replicated-requests -n $n; done
elapsed: 249.914073ms
goroutine counts: 1
elapsed: 116.811412ms
goroutine counts: 1
elapsed: 15.810965ms
goroutine counts: 1
elapsed: 5.915573ms
goroutine counts: 1
```
