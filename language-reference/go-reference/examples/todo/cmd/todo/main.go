// Demonstrated in this file:
// Handle an error condition in a command-line tool (REF01)
// Parse command-line flags (REF02)
// Make your CLI tool print usage information when -h or an invalid flag is passed (REF03)
// Allow your CLI tool's users to pipe input from other tools (REF04)

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	// "strings"

	"github.com/RockLikeAmadeus/go-cli-projects/todo"
)

const todoFileName = ".todo.json"

func main() {

	// REF03 - NOTE: Calling the Usage() function is only necessary
	// if you intend to display custom usage output. By parsing flags
	// with the flags package, usage info will be displayed using the
	// -h flag or by passing unknown arguments to the cmd by default!
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s tool. Developed according to Powerful Command Line Applications in Go\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2020\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	// REF02
	//task := flag.String("task", "", "Task to be included in the ToDo list")
	add := flag.Bool("add", false, "Add task to the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item number to be completed")
	flag.Parse()

	l := &todo.List{}

	// Use the Get method to read ToDo items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err) // REF01
		os.Exit(1)
	}

	// Decide what to do based on the provided flags
	switch {
	case *list:
		// List current todo items
		fmt.Print(l)
	case *complete > 0:
		// Complete the given item
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		//Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	// case *task != "":
	// 	// Add the task
	// 	l.Add(*task)
	// 	// Save the new list
	// 	if err := l.Save(todoFileName); err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		os.Exit(1)
	// 	}
	case *add:
		// When any arguments (excluding flags) are provided,
		// they will be used as the new task
		t, err := getTask(os.Stdin, flag.Args()...) // REF04
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(t)
		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	// Concatenate all provided arguments with a space and
	// add to the list as an item
	// Note, the book says to print an error message here,
	// but why not just leave the same behavior as if we had
	// specified the -list flag?
	default:
		// Concatenate all arguments with a space
		item := strings.Join(os.Args[1:], " ")
		// Add the task
		l.Add(item)
		// Save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank.")
	}

	return s.Text(), nil
}
