# Contents

- [Contents](#contents)
- [Haven't written Go code lately? Start here.](#havent-written-go-code-lately-start-here)
	- [Organzing Go Code](#organzing-go-code)
	- [Basic Syntax](#basic-syntax)
	- [Basic Commands](#basic-commands)
	- [Best Practices](#best-practices)
	- [Surprising Features](#surprising-features)
	- [Key Style Notes](#key-style-notes)
	- [Collections](#collections)
		- [Arrays](#arrays)
		- [Slices](#slices)
	- [Structs](#structs)
		- [Constructing Structs](#constructing-structs)
	- [Interfaces](#interfaces)
	- [Error Handling](#error-handling)
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

Go is like most programming languages, but these are the main things to be aware of that make it a little different:

- Variables can be declared at the package level, but the `:=` short initialization syntax is not valid at package level, so `var` must be used.

- Go does not allow unused variables. If a value is intentionally not used, give it the name `_`.

- Functions and methods can return multiple values.

- There are no private and public keywords in Go. To make a field (or a `struct` itself) public, you just need to capitalize the first letter of its name. For example, Laptop instead of laptop and Cpu instead of cpu. Data visibility in Go is scoped to the package, meaning private fields, methods, or types are still accessible to other package members.

- You can define _named return values_, which create a new variable in your function initialized with its zero value. This will provide some automatic documentation functionality, but it should generally be used only when the meaning of the output isn't clear from the context. It looks like this:
```go
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
```

- When you have more than one argument of the same type in a function rather than having (x int, y int) you can shorten it to (x, y int). This works for declaring variables (and I think structs) as well.

- Go permits _variadic functions_ that can take a variable number of arguments.

- Functions can be declared within other functions to reduce the surface area of your API while still breaking your logic up into sub-functions to follow the Clean Code style for readability. It would look like this:
``` go
func outerFunc() {
	innerFunc := func() {
		...
	}
	...
}
```

- Interfaces are implicitly implemented by structs. There is no `implements` keyword, or anything similar.

- When you call a function or method, the arguments are always copied. Passing by reference, or attempting to avoid copies of large data structures, requires explicit handling of pointers. Struct pointers are automatically de-referenced so if the type is `*Account` (a pointer to an account), you can access it like `myAccount.Balance()`, which is really a shortcut for `(*myAccount).Balance()`.

- When a function returns a pointer, you need to make sure you manualy check for `nil`.

- Go lets you create simple new types as wrappers of other types, as in `type Money int`. You can create instances of this wrapper type with the syntax `Money(225)`. The type is still represented as an integer, but you can declare methods on these types, which makes this feature surprisingly useful. For instance, you can implement the `Stringer` interface from `fmt` to define how your type is printed when used with the `%s` format string.

- Casting syntax is backwards; the value, not the type, goes in parenthesis: `int64(*myIntWrapper)`

## Key Style Notes

- names should use `camelCase` or `PascalCase`, rather than `snake_case`

## Collections

The foundational collections are arrays, slices, and maps. Iterate over arrays and slices with `range`

```go
for _, number := range numbers {
	sum += number
}
```

### Arrays

Arrays have a fixed capacity at declaration, and can be initialized in the following ways

```go
var myArray [5]int
numbers := [5]int{1, 2, 3, 4, 5}
moreNums := [...]int{1, 2, 3, 4, 5}
```

The type of an array, i.e. in function arguments, includes the arryay's size, and looks like this

```go
func Sum(numbers [5]int) int {
	...
}
```


### Slices

Slices look just like arrays, but don't encode the size in their types and are as such more flexible. They are implemented as references to an underlying array, and can be declared in a few different ways

```go
 // slice points to nil
var slice []int
 // allocates the memory with a starting capacity and zero-value elements
mySlice := make([]int, 3)
 // initialize the slize with real values
numbers := []int{1, 2, 3}
```

Slices can be sliced like `slice[low:high]`, where either `low` or `high` can be omitted.

## Structs

```go
type Rectangle struct {
	Width  float64
	Height float64
}

// A method is a function with a receiver. A receiver is _not_ an argument.
func (r Rectangle) Area() float64 {
	// By convention, name the receiver variable the first letter of the type
	return r.Width * r.Height
}
```

Often (most of the time?) you'll want to specify pointer receivers (`func (r *Rectangle)`), particularly if your methods will change the underlying value of the object. If you use a pointer receiver for any of your methods, be sure and update all of your methods to take a pointer receiver, for consistency.

### Constructing Structs

```go
cash := new(Money) // zero value; the type of cash is *Money
cashVal := *(new(Money)) // the type of cashVal here is Money. Is this bad style, though?
```

## Interfaces

```go
type Shape interface {
	Area() float64
}
```

## Error Handling

If it's possible for your function to fail, it is idiomatic to return an `err` for the caller to check and act on. The return type (or one of the return types) on your function should be the interface `error`, which you import from the `errors` package.

```go
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("oh no")
	}

	w.balance -= amount
	return nil
}
```

Some unit tests will _expect_ an error. Errors can be converted to a string with `.Error()`. If we don't receive an error, we don't want to test properties of the error, so we call `t.Fatal()`.

```go
assertError := func(t testing.TB, got error, want string) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
```

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