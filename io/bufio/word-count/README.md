# Word Count Example

This code accepts texts from stdin and returns the word count result to stdout.

```console
$ go run main.go
foo bar
--- word count result ---
map[bar:1 foo:1]
--- word count result ---
foo bar baz
--- word count result ---
map[bar:2 baz:1 foo:2]
--- word count result ---
^Csignal: interrupt
```
