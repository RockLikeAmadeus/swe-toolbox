
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

[contents]

# Syntax and Style

## Essential Syntax and Style

_Examples of the most important syntax and style to know without having to go all the way to external resources.
Also include the most notable syntax-related things to know about that make this language different than other languages. Delete this block once the reference for this language is "complete"_

[contents]

## Comprehensive Syntax and Style

_Links to relevant example syntax and style guides, like learnxinyminutes.com, or official style guides. Delete this block once the reference for this language is "complete"_

[contents]