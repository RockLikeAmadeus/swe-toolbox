// Demonstrated in this file:
// Create acceptance tests for a CLI tool (REF01)
// Specify setup or teardown logic to execute before and/or after your tests (REF02)
// Determine the operating system running my code (REF03)
// Retrieve the current working directory for a CLI tool (REF04)
// Organize your test cases into subtests (REF05)
// Write an acceptance test for piping text into the input of a CLI tool (REF06)

package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// REF01

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) { // REF02
	setup()

	fmt.Println("Running tests...")
	result := m.Run()

	teardown()

	os.Exit(result)
}

func setup() {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" { // REF03
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	// Remove the list if it exists
	os.Remove(fileName)
}

func teardown() {
	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)
}

func TestTodoCLI(t *testing.T) {
	task := "test task number 1"

	dir, err := os.Getwd() // REF04
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	// Note: these subtests are not actually independent, should fix
	t.Run("AddNewTaskWithoutFlags", func(t *testing.T) { // REF05
		cmd := exec.Command(cmdPath, strings.Split(task, " ")...)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		os.Remove(fileName)
	})

	t.Run("AddNewTaskWithFlag", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	task2 := "test task number 2"
	t.Run("AddNewTaskFromSTDIN", func(t *testing.T) { // REF06
		cmd := exec.Command(cmdPath, "-add")
		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		io.WriteString(cmdStdIn, task2)
		cmdStdIn.Close()

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf(" 1: %s\n 2: %s\n", task, task2)

		if expected != string(out) {
			t.Errorf("Expected %q, received %q\n", expected, string(out))
		}
	})
}
