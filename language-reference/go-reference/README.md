
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

- Interfaces are implicitly implemented by structs. There is no `implements` keyword, or anything similar.

- When you call a function or method, the arguments are always copied. Passing by reference, or attempting to avoid copies of large data structures, requires explicit handling of pointers. Struct pointers are automatically de-referenced so if the type is `*Account` (a pointer to an account), you can access it like `myAccount.Balance()`, which is really a shortcut for `(*myAccount).Balance()`.

- When a function returns a pointer, you need to make sure you manually check for `nil`.

- Go lets you create simple new types as wrappers of other types, as in `type Money int`, akin to C's `typedef`. You can create instances of this wrapper type with the syntax `Money(225)`. The type is still represented as an integer, but you can declare methods on these types, which makes this feature surprisingly useful. For instance, you can implement the `Stringer` interface from `fmt` to define how your type is printed when used with the `%s` format string.

- Casting syntax is backwards; the value, not the type, goes in parenthesis: `int64(*myIntWrapper)`
  
- The only iteration construct in Go is `for`, and there is special syntax for iterating over a range.

[contents]

# What do I need installed to write Go code

Simply [install Go](https://go.dev/doc/install) and run `$ go version`. All of the workflow tools you're likely to need are packaged with the language.

[contents]

# I just want to write a quick script with Go

In a text file called `[my-script].go`, write

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

...and run `$ go run`. The `main()` function in Go is the entry point of executable applications.

_Notes on writing and then running a standalone script or small process here. Include the simplest way to retrieve command line arguments. Delete this block once the reference for this language is "complete"_

[contents]

# I'm starting a new software project with Go

The top-level organizational unit in Go, the unit of _releasability_, is the **Module** (run `go help modules`). Dependencies are managed within a single module, which _typically makes up a single project or library_.

Modules are composed of **Packages** (run `go help packages`). A Go package is an organizational unit that represents a distinct concept and has its own namespace. It can be made up by a single file, or can be spread across multiple files. Technically, Go packages are just _directories_ of `.go` files that start with the `package` statement (by itself on a single line, followed by the import statements and then the rest of the code).    

To create a new Go module (project), navigate to the directory that you want to be the project root (you'll have to create a new directory with your project name and navigate into it first; use dashes `-` to separate words), and run `$ go mod init github.com/RockLikeAmadeus/my-project`. For more info, run `go help mod init`. For advice on structuring your new project, check out [this article](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project).

More info [here](./code-organization.md)

[contents]

# Syntax and Style

## Essential Syntax and Style

### Collections

The foundational collections are arrays, slices, and maps. Iterate over arrays and slices with `range`

```go
for _, number := range numbers {
	sum += number
}
```

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

#### Constructing Structs

```go
cash := new(Money) // zero value; the type of cash is *Money
cashVal := *(new(Money)) // the type of cashVal here is Money. Is this bad style, though?
```

### Interfaces

```go
type Shape interface {
	Area() float64
}
```

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
### Style

- Names should use `camelCase` or `PascalCase`, rather than `snake_case`

[contents]

## Comprehensive Syntax and Style

[Learn Go in Y minutes](https://learnxinyminutes.com/go/)
<br>
[Go by Example](https://gobyexample.com/)
<br>
[Go Style Guide](https://google.github.io/styleguide/go/)

[contents]