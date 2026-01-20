# Contents

- [Contents](#contents)
- [What makes Go unique?](#what-makes-go-unique)
- [What do I need installed to write Go code](#what-do-i-need-installed-to-write-go-code)
- [I just want to write a quick script with Go](#i-just-want-to-write-a-quick-script-with-go)
- [I'm starting a new software project with Go](#im-starting-a-new-software-project-with-go)
- [Syntax and Style](#syntax-and-style)
	- [Essential Syntax and Style](#essential-syntax-and-style)
		- [Collections](#collections)
			- [Arrays](#arrays)
			- [Slices](#slices)
		- [Structs](#structs)
			- [Constructing Structs](#constructing-structs)
		- [Interfaces](#interfaces)
		- [Error Handling](#error-handling)
		- [Style](#style)
	- [Comprehensive Syntax and Style](#comprehensive-syntax-and-style)
	- [Documenting  Code](#documenting--code)
- [Basics of Testing](#basics-of-testing)
	- [Subtests and Test Helpers](#subtests-and-test-helpers)
	- [Testable Examples](#testable-examples)
	- [Table Driven Tests](#table-driven-tests)
	- [Testing Functions That Require Files](#testing-functions-that-require-files)
	- [Executing only short-running tests](#executing-only-short-running-tests)
- [Useful Library Imports](#useful-library-imports)
- [Essential Language Toolset Commands](#essential-language-toolset-commands)
- [Tools](#tools)
- [Utilities](#utilities)
- [How To](#how-to)
	- [Types of Applications](#types-of-applications)
		- [CLI Utilities](#cli-utilities)
			- [Create a CLI Utility](#create-a-cli-utility)
			- [Create acceptance tests for a CLI tool (REF01)](#create-acceptance-tests-for-a-cli-tool-ref01)
			- [Handle an error condition in a command-line tool (REF01)](#handle-an-error-condition-in-a-command-line-tool-ref01)
			- [Parse command-line flags (REF02)](#parse-command-line-flags-ref02)
			- [Make your CLI tool print usage information when -h or an invalid flag is passed (REF03)](#make-your-cli-tool-print-usage-information-when--h-or-an-invalid-flag-is-passed-ref03)
			- [Allow your CLI tool's users to pipe input from other tools (REF04)](#allow-your-cli-tools-users-to-pipe-input-from-other-tools-ref04)
			- [Retrieve the current working directory for a CLI tool (REF04)](#retrieve-the-current-working-directory-for-a-cli-tool-ref04)
			- [Write an acceptance test for piping text into the input of a CLI tool (REF06)](#write-an-acceptance-test-for-piping-text-into-the-input-of-a-cli-tool-ref06)
		- [Write tests for the output of a tool that prints to STDOUT or to some other location](#write-tests-for-the-output-of-a-tool-that-prints-to-stdout-or-to-some-other-location)
	- [Strings](#strings)
		- [Define a type's custom string printing behavior (REF06)](#define-a-types-custom-string-printing-behavior-ref06)
	- [Functions](#functions)
		- [Pass each element of a slice as an individual argument to a variadic function (REF01)](#pass-each-element-of-a-slice-as-an-individual-argument-to-a-variadic-function-ref01)
		- [Define a function variable that can be called polymorphically](#define-a-function-variable-that-can-be-called-polymorphically)
	- [Types](#types)
		- [Define a type's custom string printing behavior (REF06)](#define-a-types-custom-string-printing-behavior-ref06-1)
	- [Data Structures](#data-structures)
		- [Floats](#floats)
			- [Compare two floating point numbers](#compare-two-floating-point-numbers)
		- [Slices](#slices-1)
			- [Pass each element of a slice as an individual argument to a variadic function (REF01)](#pass-each-element-of-a-slice-as-an-individual-argument-to-a-variadic-function-ref01-1)
			- [Remove an item at a specific index from a slice (REF02)](#remove-an-item-at-a-specific-index-from-a-slice-ref02)
	- [Data Serialization](#data-serialization)
		- [Convert an object to JSON and write it to a file (REF03)](#convert-an-object-to-json-and-write-it-to-a-file-ref03)
		- [Read from a JSON file and parse the contents (REF04)](#read-from-a-json-file-and-parse-the-contents-ref04)
	- [File I/O](#file-io)
		- [Reading from files](#reading-from-files)
			- [Read from a JSON file and parse the contents (REF04)](#read-from-a-json-file-and-parse-the-contents-ref04-1)
		- [Writing to files](#writing-to-files)
			- [Convert an object to JSON and write it to a file (REF03)](#convert-an-object-to-json-and-write-it-to-a-file-ref03-1)
			- [Create a temporary file (REF01)](#create-a-temporary-file-ref01)
			- [Generate a file that has fixed content, then inject dynamic data at runtime](#generate-a-file-that-has-fixed-content-then-inject-dynamic-data-at-runtime)
	- [Error Handling](#error-handling-1)
		- [Determine if an error is a specific type of error (REF05)](#determine-if-an-error-is-a-specific-type-of-error-ref05)
		- [Handle an error condition in a command-line tool (REF01)](#handle-an-error-condition-in-a-command-line-tool-ref01-1)
		- [Define a custom error type](#define-a-custom-error-type)
		- [Return an instance of type `error` including the inner error](#return-an-instance-of-type-error-including-the-inner-error)
	- [Logging](#logging)
		- [Logging](#logging-1)
	- [Testing](#testing)
		- [Specify setup or teardown logic to execute before and/or after your tests (REF02)](#specify-setup-or-teardown-logic-to-execute-before-andor-after-your-tests-ref02)
		- [Organize your test cases into subtests (REF05)](#organize-your-test-cases-into-subtests-ref05)
		- [Create acceptance tests for a CLI tool (REF01)](#create-acceptance-tests-for-a-cli-tool-ref01-1)
		- [Write an acceptance test for piping text into the input of a CLI tool (REF06)](#write-an-acceptance-test-for-piping-text-into-the-input-of-a-cli-tool-ref06-1)
		- [Write tests for the output of a tool that prints to STDOUT or to some other location](#write-tests-for-the-output-of-a-tool-that-prints-to-stdout-or-to-some-other-location-1)
		- [Write multiple test cases exercising the same function run under different conditions](#write-multiple-test-cases-exercising-the-same-function-run-under-different-conditions)
		- [Create a function to encapsulate repeated behavior across different tests](#create-a-function-to-encapsulate-repeated-behavior-across-different-tests)
	- [OS Access](#os-access)
		- [Determine the operating system running my code (REF03)](#determine-the-operating-system-running-my-code-ref03)

# What makes Go unique?

Go is a compiled language with very simple syntax. It's small, compiles fast, and comes packaged with great tooling, including things like testing and benchmarking built into the language. It was built for the web, and includes simple primitives for easy handling of concurrent code. It's a great language for web applications, CLIs, and even systems programming.

Syntactically, Go is like most programming languages, but these are the main things to be aware of that make it a little different:

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

- Interfaces are implicitly implemented by structs. There is no `implements` keyword, or anything similar. Also, Go interfaces only define behavior, not state, so they only define what a type should be able to do, not what data it should hold.

- When you call a function or method, the arguments are always copied. Passing by reference, or attempting to avoid copies of large data structures, requires explicit handling of pointers. Struct pointers are automatically de-referenced so if the type is `*Account` (a pointer to an account), you can access it like `myAccount.Balance()`, which is really a shortcut for `(*myAccount).Balance()`.

- When a function returns a pointer, you need to make sure you manually check for `nil`.

- Go lets you create simple new types as wrappers of other types, as in `type Money int`, akin to C's `typedef`. You can create instances of this wrapper type with the syntax `Money(225)`. The type is still represented as an integer, but you can declare methods on these types, which makes this feature surprisingly useful. For instance, you can implement the `Stringer` interface from `fmt` to define how your type is printed when used with the `%s` format string.

- Casting syntax is backwards; the value, not the type, goes in parenthesis: `int64(*myIntWrapper)`
  
- The only iteration construct in Go is `for`, and there is special syntax for iterating over a range.

- Many functions that operate on common types, like `append()` and `len()`, which would normally be methods in other languages, are built-in functions that operate without a receiver; that is, they are called by passing in the instance or value, rather than called by dereferencing the instance, i.e. `append(l, item)` rather than `l.append(item)`

[Contents](#contents)

# What do I need installed to write Go code

Simply [install Go](https://go.dev/doc/install) and run `$ go version`. All of the workflow tools you're likely to need are packaged with the language.

[Contents](#contents)

# I just want to write a quick script with Go

In a text file called `[my-script].go`, write

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

...and run `$ go run`.

You create an executable in Go by creating a package name `main` that defines a function called `main()`.  The `main()` function in Go is the entry point of executable applications. The `main` package is typically defined in a file called `main.go`, but this is a convention rather than a requirement.

_Notes on writing and then running a standalone script or small process here. Include the simplest way to retrieve command line arguments. Delete this block once the reference for this language is "complete"_

[Contents](#contents)

# I'm starting a new software project with Go

The top-level organizational unit in Go, the unit of _releasability_, is the **Module** (run `go help modules`). Dependencies are managed within a single module, which _typically makes up a single project or library_.

Modules are composed of **Packages** (run `go help packages`). A Go package is an organizational unit that represents a distinct concept and has its own namespace. It can be made up by a single file, or can be spread across multiple files. Technically, Go packages are just _directories_ of `.go` files that start with the `package` statement (by itself on a single line, followed by the import statements and then the rest of the code).    

To create a new Go module (project), navigate to the directory that you want to be the project root (you'll have to create a new directory with your project name and navigate into it first; use dashes `-` to separate words), and run `$ go mod init github.com/RockLikeAmadeus/my-project`. For more info, run `go help mod init`. For advice on structuring your new project, check out [this article](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project).

More info [here](./code-organization.md)

For specifics on creating a command-line (CLI) utility in particular, go [here](create-a-cli-utility.md).

[Contents](#contents)

# Syntax and Style

## Essential Syntax and Style

### Collections

The foundational collections are arrays, slices, and maps. Iterate over arrays and slices with `range`

```go
for _, number := range numbers {
	sum += number
}
```

[Contents](#contents)

#### Arrays

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

[Contents](#contents)

#### Slices

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

[Contents](#contents)

### Structs

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

[Contents](#contents)

#### Constructing Structs

```go
cash := new(Money) // zero value; the type of cash is *Money
cashVal := *(new(Money)) // the type of cashVal here is Money.
```

[Contents](#contents)

### Interfaces

```go
type Shape interface {
	Area() float64
}
```

[Contents](#contents)

### Error Handling

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

[Contents](#contents)

### Style

- Names should use `camelCase` or `PascalCase`, rather than `snake_case`

[Contents](#contents)

## Comprehensive Syntax and Style

[Learn Go in Y minutes](https://learnxinyminutes.com/go/)
<br>
[Go by Example](https://gobyexample.com/)
<br>
[Go Style Guide](https://google.github.io/styleguide/go/)
<br>
[Effective Go](https://go.dev/doc/effective_go)
<br>
[Full Language Specification](https://go.dev/ref/spec)
<br>
[Standard Library Source Code (for reference to best practices, expected style, and idioms)](https://go.dev/src/)

## Documenting <Language> Code

The convention is simple: to document a type, variable, constant, function, or even a package, write a regular comment directly preceding its declaration, with no intervening blank line. Godoc will then present that comment as text alongside the item it documents. For example, this is the documentation for the fmt package’s Fprint function:

```go
// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
```
Notice this comment is a complete sentence that begins with the name of the element it describes.

For more info including formatting concerns, as well as details for documenting packages, known issues, or deprecated functionality, go [here](https://go.dev/blog/godoc).

Documentation pages can be generated with the [`godoc` command](https://pkg.go.dev/golang.org/x/tools/cmd/godoc).

[Contents](#contents)

# Basics of Testing

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

## Subtests and Test Helpers

You can also define [subtests](https://pkg.go.dev/testing#hdr-Subtests_and_Sub_benchmarks). Additionally, it's a good idea to define [helper](https://pkg.go.dev/testing#B.Helper) methods to make your code simpler when possible (for things like setup and teardown, or other use cases, see below), which requires passing an instance of `testing.TB` and calling `t.Helper()`.

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
$ go test -v # verbose
$ go test -v ./... # recursive and verbose
```

[Contents](#contents)

## Testable Examples

For now, see [here](https://go.dev/blog/examples).

[Contents](#contents)

## Table Driven Tests

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

[Contents](#contents)

## Testing Functions That Require Files

Obviously, where it makes sense, file reading and writing should use interfaces so that your tests can provide test doubles and don't require actual file system access. However, for acceptance tests where you may actually want to read from sample input files and/or verify that application output matches "golden files" for simpler testing workflows, the correct way place to store these files is in a subdirectory called `testdata` under the project's directory. Go's build tool specifically ignores the `testdata` directory so that these testing artifacts don't end up in as part of your build.

## Executing only short-running tests

```bash
go test -short ./...
```

We can add the following to our _acceptance_ tests to see if the user wants to run our acceptance tests by inspecting the value of the flag

```go
if testing.Short() {
	t.Skip()
}
```

[Contents](#contents)

# Useful Library Imports

[Go Standard Library Documentation](https://pkg.go.dev/std)

```
import (
	"fmt"              // Formatted I/O, analogous to C's printf and scanf (https://pkg.go.dev/fmt)
	"io"               // Implements some I/O utility functions.
	m "math"           // Math library with local alias m.
	"net/http"         // Yes, a web server!
	_ "net/http/pprof" // Profiling library imported only for side effects
	"os"               // OS functions like working with the file system
	"strconv"          // String conversions.
	"flag"             // Command line flags!
)
```

[Contents](#contents)

# Essential Language Toolset Commands

```bash
$ go mod tidy # Clean up your Go mod file
$ go fmt
$ go build # Note: if multiple packages are specified, does not generate output,
# which is useful for checking for compilation errors without creating binaries
```

[Contents](#contents)

# Tools

[Awesome Go](https://awesome-go.com/) - like blessed.rs, but for Go.

[Wails](https://wails.io/) - like Tauri, but for Go.

[Bubble Tea](https://github.com/charmbracelet/bubbletea) - For TUIS; comes from [Charm](https://charm.sh/)

[Cobra](https://github.com/spf13/cobra) - for CLIs, widely used for tools written in Go.

[errcheck](https://github.com/kisielk/errcheck) - automatically exercise your error checking coverage.

[Contents](#contents)

# Utilities

[Funk](https://github.com/thoas/go-funk) - map, filter, contains, etc.

[Charm Log](https://github.com/charmbracelet/log) - logging, also from Charm

[Contents](#contents)

# How To

## Types of Applications

### CLI Utilities

#### [Create a CLI Utility](./create-a-cli-utility.md)
#### [Create acceptance tests for a CLI tool (REF01)](./examples/todo/cmd/todo/main_test.go)
#### [Handle an error condition in a command-line tool (REF01)](./examples/todo/cmd/todo/main.go)
#### [Parse command-line flags (REF02)](./examples/todo/cmd/todo/main.go)
#### [Make your CLI tool print usage information when -h or an invalid flag is passed (REF03)](./examples/todo/cmd/todo/main.go)
#### [Allow your CLI tool's users to pipe input from other tools (REF04)](./examples/todo/cmd/todo/main.go)
#### [Retrieve the current working directory for a CLI tool (REF04)](./examples/todo/cmd/todo/main_test.go)
#### [Write an acceptance test for piping text into the input of a CLI tool (REF06)](./examples/todo/cmd/todo/main_test.go)
### [Write tests for the output of a tool that prints to STDOUT or to some other location](./testing.md#write-tests-for-the-output-of-a-tool-that-prints-to-stdout-or-to-some-other-location)

## Strings

### [Define a type's custom string printing behavior (REF06)](./examples/todo/todo.go)

## Functions

### [Pass each element of a slice as an individual argument to a variadic function (REF01)](./examples/todo/todo.go)
### [Define a function variable that can be called polymorphically](./uncategorized-idioms-etc.md#polymorphic-function-use)

## Types

### [Define a type's custom string printing behavior (REF06)](./examples/todo/todo.go)

## Data Structures

### Floats
#### [Compare two floating point numbers](./uncategorized-idioms-etc.md#comparing-floats)

### Slices
#### [Pass each element of a slice as an individual argument to a variadic function (REF01)](./examples/todo/todo.go)
#### [Remove an item at a specific index from a slice (REF02)](./examples/todo/todo.go)


## Data Serialization

### [Convert an object to JSON and write it to a file (REF03)](./examples/todo/todo.go)
### [Read from a JSON file and parse the contents (REF04)](./examples/todo/todo.go)

## File I/O

### Reading from files

#### [Read from a JSON file and parse the contents (REF04)](./examples/todo/todo.go)

### Writing to files

#### [Convert an object to JSON and write it to a file (REF03)](./examples/todo/todo.go)
#### [Create a temporary file (REF01)](./examples/todo/todo_test.go)
#### [Generate a file that has fixed content, then inject dynamic data at runtime](./file-io.md#generate-a-file-that-has-fix-content-then-inject-dynamic-data-at-runtime)

## Error Handling

### [Determine if an error is a specific type of error (REF05)](./examples/todo/todo.go)
### [Handle an error condition in a command-line tool (REF01)](./examples/todo/cmd/todo/main.go)
### [Define a custom error type](./error-handling.md#defining-custom-error-types)
### [Return an instance of type `error` including the inner error](./error-handling.md#returning-an-instance-of-type-error-including-the-inner-error)

## Logging
### [Logging](./logging.md)

## Testing

### [Specify setup or teardown logic to execute before and/or after your tests (REF02)](./examples/todo/cmd/todo/main_test.go)
### [Organize your test cases into subtests (REF05)](./examples/todo/cmd/todo/main_test.go)
### [Create acceptance tests for a CLI tool (REF01)](./examples/todo/cmd/todo/main_test.go)
### [Write an acceptance test for piping text into the input of a CLI tool (REF06)](./examples/todo/cmd/todo/main_test.go)
### [Write tests for the output of a tool that prints to STDOUT or to some other location](./testing.md#write-tests-for-the-output-of-a-tool-that-prints-to-stdout-or-to-some-other-location)
### [Write multiple test cases exercising the same function run under different conditions](#table-driven-tests)
### [Create a function to encapsulate repeated behavior across different tests](#subtests-and-test-helpers)

## OS Access

### [Determine the operating system running my code (REF03)](./examples/todo/cmd/todo/main_test.go)













[Contents](#contents)