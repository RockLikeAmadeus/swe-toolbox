# Haven't written Go code lately? Start here.

Go **Modules** are composed of **Packages** (run `go help modules`). Go **Packages** are just directories of `.go` files that start with the `package` statement (run `go help packages`).

To create a new Go module (project), navigate to the directory that you want to be the project root, and run `$ go mod init github.com/RockLikeAmadeus/my-project`. For more info, run `go help mod init`. For advice on structuring your new project, check out [this article](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project).

The `main()` function in Go is the entry point of executable applications.

Try `$ go help mod`

For basic syntax, go [here](https://learnxinyminutes.com/docs/go/).

A few oddities to be aware of, though:

- There are no private and public keywords in Go. To make a field (or a `struct` itself) public, you just need to capitalize the first letter of its name. For example, Laptop instead of laptop and Cpu instead of cpu.

- You can define _named return values_, which create a new variable in your function initialized with its zero value. This will provide some automatic documentation functionality, but it should generally be used only when the meaning of the output isn't clear from the context.

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

# Different ways of instantiating structs

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

---

Tools: 

[Awesome Go](https://awesome-go.com/) - like blessed.rs, but for Go.

[Wails](https://wails.io/) - like Tauri, but for Go.

[Cobra](https://github.com/spf13/cobra) - for CLIs, widely used for tools written in Go.