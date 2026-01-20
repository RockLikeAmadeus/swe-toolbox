https://go.dev/blog/error-handling-and-go

# Defining Custom Error Types

By convention, these variables are exported and their names start with `Err`

```go
var (
    ErrNotNumber    = errs.New("Data is not numeric.")
    ErrNoFiles      = errs.New("No input files.")
)
```

# Returning an instance of type `error` including the inner error

```go
allData, err := cr.ReadAll()
if err != nil {
    return fmt.Errorf("Cannot read data from file, %w", err)
}
```

