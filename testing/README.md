# Go Testing

## Materials

The followings are a collection of testing and benchmark examples/best practices in Go:

- [Go by Example: Testing and Benchmarking](https://gobyexample.com/testing-and-benchmarking)
- [Go Doc: Add a test](https://go.dev/doc/tutorial/add-a-test)
- [The Go Blog: Using Subtests and Sub-benchmarks](https://go.dev/blog/subtests)
- [The Go Blog: Testable Examples in Go](https://go.dev/blog/examples)

---

## Cheatsheet

```bash
go test

# verbose mode
go test -v

# filter test case
go test -run Driven -v
go test -run Example

go test -v | grep PASS
go test -v | grep FAIL

go test -bench .
```
