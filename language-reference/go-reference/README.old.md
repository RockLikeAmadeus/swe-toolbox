# Contents

- [Contents](#contents)
- [Haven't written Go code lately? Start here.](#havent-written-go-code-lately-start-here)
	- [Organzing Go Code](#organzing-go-code)
	- [Basic Syntax](#basic-syntax)
	- [Basic Commands](#basic-commands)
	- [Best Practices](#best-practices)
	- [Surprising Features](#surprising-features)
	- [Key Style Notes](#key-style-notes)
	- [Writing tests](#writing-tests)
		- [Testable Examples](#testable-examples)
		- [Table Driven Tests](#table-driven-tests)
		- [Executing only short-running tests](#executing-only-short-running-tests)
	- [Tools:](#tools)
		- [Utilities](#utilities)

TODO: Write a quick tool that I can run against a directory containing markdown files, that will recursively search for all markdown files and upate them with a useful table of contents at the top.

# Haven't written Go code lately? Start here.

## Organzing Go Code

Go **Modules** are composed of **Packages** (run `go help modules`). Go **Packages** are just directories of `.go` files that start with the `package` statement (run `go help packages`).    

To create a new Go module (project), navigate to the directory that you want to be the project root, and run `$ go mod init github.com/RockLikeAmadeus/my-project`. For more info, run `go help mod init`. For advice on structuring your new project, check out [this article](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project).

The `main()` function in Go is the entry point of executable applications.

Try `$ go help mod`

More info [here](./code-organization.md)

## Basic Syntax

For basic syntax, go [here](https://learnxinyminutes.com/docs/go/).

## Basic Commands

```bash
$ go mod tidy # Clean up your Go mod file
```

## Best Practices

[Effective Go](https://go.dev/doc/effective_go)

## Surprising Features



## Key Style Notes





## Writing tests

Go unit tests can be organized in two ways. The most basic Go unit tests live in files alongside (at the same level of) the files that they test, and must have the same name as the file suffixed with `_test`.

The other, probably better way to organize tests is to have your tests in a `_test` package that imports the package under test, so that it can truly exercises only the public interface of the package.

In their most basic form, tests look like this:

```go
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

You can also define subtests. Additionally, it's a good idea to define helper methods to make your code simpler when possible, which requires passing an instance of `testing.TB` and calling `t.Helper()`.

```go
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

```

Run all tests with 
```bash
$ go test
$ go test ./... # recursive
$ go test ./... # verbose
$ go test -v ./... # recursive and verbose
```

### Testable Examples

For now, see [here](https://go.dev/blog/examples).

### Table Driven Tests

If you want to test identical behavior with a large number of different inputs, you can make your test logic more compact and easier to extend with table-driven tests. In this example, we start by initializing a slice of struct instances, where the type is defined inline using _anonymous struct_ syntax.

```go
func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v received %g expected %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
```

With the above syntax, the output of the test will be quite friendly, like this:

```
--- FAIL: TestArea (0.00s)
    --- FAIL: TestArea/Rectangle (0.00s)
        shapes_test.go:33: main.Rectangle{Width:12, Height:6} received 72.00 expected 72.10
```

### Executing only short-running tests

```bash
go test -short ./...
```

We can add the following to our _acceptance_ tests to see if the user wants to run our acceptance tests by inspecting the value of the flag

```go
if testing.Short() {
	t.Skip()
}
```

## Tools: 

[Awesome Go](https://awesome-go.com/) - like blessed.rs, but for Go.

[Wails](https://wails.io/) - like Tauri, but for Go.

[Bubble Tea](https://github.com/charmbracelet/bubbletea) - For TUIS; comes from [Charm](https://charm.sh/)

[Cobra](https://github.com/spf13/cobra) - for CLIs, widely used for tools written in Go.

[errcheck](https://github.com/kisielk/errcheck) - automatically exercise your error checking coverage.

[Charm Log](https://github.com/charmbracelet/log) - logging, also from Charm

### Utilities
[Funk](https://github.com/thoas/go-funk) - map, filter, contains, etc.