# Polymorphic Function Use

You can make a type which defines a function signature. Any function that matches the signature qualifies as this type:

```go
type statsFunc func(data []float64) float64
```

Then, elsewhere:

```go
var opFunc statsFunc
```

Functions are first class citizens in Go. The variable `opFunc` can hold a reference to any function that satisfies the `statsFunc` signature, and can be called polymorphically (without knowing the actual function its calling).

```go
switch op {
case "sum":
    opFunc = sum
case "avg"
    opFunc = avg
}

res := opFunc(someSlice)
```