package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	// `Buffer` type implements `Writer`
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
