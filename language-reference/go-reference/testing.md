# Write tests for the output of a tool that prints to STDOUT or to some other location

When a application, such as a command line tool, prints output to the user, or writes output to a file, or some other location, we need a way to write automated tests for the functions that do the writing. The appropriate way is for those functions (perhaps a `run()` function in a command-line tool) to have an output parameter that accepts an instance of the interface `io.Writer`:

```go
func run(filename string, out io.Writer) error {
...
```

The `main` function which calls run will pass `os.Stdout` to this function:

```go
if err := run(*filename, os.Stdout); err != nil {
...
```

Then, in the test case for this behavior, use a `bytes.Buffer` to capture the output:

```go
var mockStdOut bytes.Buffer

if err := run(inputFile, &mockStdOut); err != nil {
    t.Fatal(err)
}

output := strings.TrimSpace(mockStdOut.String())
```

# Test Doubles

For examples of how this is done straight from the Go standard library, see [the source of the `exec_test.go` file](https://go.dev/src/os/exec/exec_test.go) from the `os` package.