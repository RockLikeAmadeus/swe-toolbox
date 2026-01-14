You can run this utility by piping in any shell expression that returns a string, like this:

`$ cat main.go | ./wc -l`

```go
// Try building this and running `$ cat main.go | ./wc -l`

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
)

func main() {
	// Defining a boolean flag -l to count lines instead of words. The flag package provides utilities for command line flags.
	lines := flag.Bool("l", false, "Count lines")

	// Parsing the flags provided by the user
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	// A scanner is used to read text from a Reader (such as files) delimited by spaces or new lines
	scanner := bufio.NewScanner(r)

	// If the count lines flag is not set, we want to count words so we define
	// the scanner split type to words (default is split by lines)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	// Defining a counter
	wordCount := 0

	// For every word scanned, increment the counter (but in real code, check for errors!)
	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	expected := 4
	result := count(b, false)
	if result != expected {
		t.Errorf("TestCountWords: Expected %d, received %d instead.\n", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline2 word1")
	expected := 3
	result := count(b, true)
	if result != expected {
		t.Errorf("TestCountLines: Expected %d, received %d instead\n", expected, result)
	}
}
```