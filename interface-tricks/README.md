# Interface Tricks

I try to use the following example to demonstrate this interface trick that is commonly seen in the Golang's source code, For examples, see the `http.Handler` interface and the `http.HandlerFunc` type.

This trick is intended to convert a function whose signature is the same as the method signature of an interface into a type that implements that interface.

As shown in `main.go`, an interface called `Ingester` is defined, having a method called `Ingest` that accepts a slice of float64 as input and returns a slice of float64.

```go
type Ingester interface {
	Ingest([]float64) []float64
}
```

Then a type `IngesterFunc` is immediately defined. The type definition is just the `Ingest`'s method signature of the `Ingester` interface.

```go
type IngesterFunc func([]float64) []float64
```

In order to let the `IngesterFunc` type implement the `Ingester` interface, the `Ingest` method needs to be provided. And as you can see its implementation is straightforward: the argument `data` that the `Ingest` method accepts is directly passed to the `IngesterFunc` type and returned.

```go
func (i IngesterFunc) Ingest(data []float64) []float64 {
	return i(data)
}
```

What is the benefit of doing this?

Next time you have a function, say `addOne`, whose signature is the same as the `Ingest` method of `Ingester`:

```go
func addOne(data []float64) []float64
```

You can pass this function to `IngesterFunc` to derive the following:

```go
IngesterFunc(addOne)
```

which can then be passed to a function (e.g., `ScaledTransform`) that accepts the `Ingester` interface as part of its arguments:

```go
func ScaledTransform(i Ingester, scale float64) Ingester
```

The argument `scale` and return type `Ingester` do not matter a lot in this example. The point here is that now the `ScaledTransform` function can accepts an interface as its input argument, rather than other primitive or struct types.

Think about what if we don't define the `IngesterFunc` type and let it implement the `Ingester` interface, we would have to write the `ScaledTransform` function as:

```go
func ScaledTransform(f func([]float64) []float64, scale float64) func([]float64) []float64
```

which is less readable compared to the version that adopts this interface trick.

It completes this demonstration.
