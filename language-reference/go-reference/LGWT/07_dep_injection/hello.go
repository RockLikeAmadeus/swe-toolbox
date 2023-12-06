package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	// Fprintf is like Printf, but it it takes a `Writer` to write to
	fmt.Fprintf(writer, "Hello, %s\n", name)
}

// `http.ResponseWriter` implements `Writer`
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "Elodie")
	// Check at http://localhost:5001/
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
