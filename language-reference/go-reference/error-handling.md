https://go.dev/blog/error-handling-and-go

# Defining Custom Error Types

By convention, these variables are exported and their names start with `Err`

```go
var (
    ErrNotNumber    = errs.New("Data is not numeric.")
    ErrNoFiles      = errs.New("No input files.")
)
```

This works for simple programs and/or simple errors, but for more complex uses, you should actually define new types. In go, the built-in interface `Error` defines a single method with the signature `Error() string`, so your custom errors should implement this type:

```go
type stepErr struct {
    step    string
    msg     string
    cause   error
}

func (s *stepErr) Error() string {
    return fmt.Sprintf("Step: %q: %s: Cause: %v", s.step, s.msg, s.cause)
}
```

You can also implement the methods `Is()` and `Unwrap()`:

```go
func (s *stepErr) Is(target error) bool {
     // This is not the preferred way
    // see https://pkg.go.dev/errors
    // and https://go.dev/blog/go1.13-errors
    t, ok := target.(*stepErr)
    if !ok {
        return false
    }

    return t.step == s.step
}

func (s *stepErr) Unwrap() error {
    return s.cause
}
```



# Returning an instance of type `error` including the inner error

```go
allData, err := cr.ReadAll()
if err != nil {
    return fmt.Errorf("Cannot read data from file, %w", err)
}
```