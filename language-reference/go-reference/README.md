# Haven't written Go code lately? Start here.

## Organzing Go Code

Go **Modules** are composed of **Packages** (run `go help modules`). Go **Packages** are just directories of `.go` files that start with the `package` statement (run `go help packages`).

To create a new Go module (project), navigate to the directory that you want to be the project root, and run `$ go mod init github.com/RockLikeAmadeus/my-project`. For more info, run `go help mod init`. For advice on structuring your new project, check out [this article](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project).

The `main()` function in Go is the entry point of executable applications.

Try `$ go help mod`

## Basic Syntax

For basic syntax, go [here](https://learnxinyminutes.com/docs/go/).

## Surprising Features

Go is like most programming languages, but these are the main things to be aware of that make it a little different:

- There are no private and public keywords in Go. To make a field (or a `struct` itself) public, you just need to capitalize the first letter of its name. For example, Laptop instead of laptop and Cpu instead of cpu.

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

TODO

# Writing tests

Go unit tests live in files alongside (at the same level of) the files that they test, and must have the same name as the file suffixed with `_test`. In their most basic form they look like this:

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

## Testable Examples

For now, see [here](https://go.dev/blog/examples).

## Tools: 

[Awesome Go](https://awesome-go.com/) - like blessed.rs, but for Go.

[Wails](https://wails.io/) - like Tauri, but for Go.

[Cobra](https://github.com/spf13/cobra) - for CLIs, widely used for tools written in Go.