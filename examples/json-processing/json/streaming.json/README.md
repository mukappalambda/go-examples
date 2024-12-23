# Streaming JSON

```console
$ go run main.go
Type something!
Example:
> alpha 10
------
alpha
2024/12/23 11:45:44 not enough fields; got 1 field only
alpha beta
2024/12/23 11:45:50 error parsing "beta" to integer
alpha 1.23
2024/12/23 11:45:59 error parsing "1.23" to integer
alpha 123
{"Name":"alpha","Score":123}
alpha 123 other
{"Name":"alpha","Score":123}
^Csignal: interrupt
```
