/*
A Go test must be in a file called xxx_test.go

The test function
- must start with the word `Test`
- must take exactly one argument of type `*testing.T`

Execute your tests with `$ go test`, or use the VS Code extension.

The TDD Cycle:
- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor
*/

package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Alec", "")
		want := "Hello, Alec"
		// Refactoring to remove duplication is *not* just for production code, but tests as well
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Claude", "French")
		want := "Bonjour, Claude"
		assertCorrectMessage(t, got, want)
	})
}

// For helper functions, accepting testing.TB is preferred
// testing.TB is an interface satisfied by both Go's test and benchmarking tools
func assertCorrectMessage(t testing.TB, got, want string) {
	// Always use this for helper functions to improve error messages
	t.Helper()

	if got != want {
		// `%q` is very useful for test printing,
		// as it wraps values in double quotes
		t.Errorf("got %q want %q", got, want)
	}
}
