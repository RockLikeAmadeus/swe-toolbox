
# What makes Go unique?

Go is a compiled language with very simple syntax. It's small, compiles fast, and comes packaged with great tooling, including things like testing and benchmarking built into the language. It was built for the web, and includes simple primitives for easy handling of concurrent code. It's a great language for web applications, CLIs, and even systems programming.

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

The top-level organizational unit in Go, the unit of _releasability_, is the **Module** (run `go help modules`). Dependencies are managed within a single module, and _typically make up a single project or library_.

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