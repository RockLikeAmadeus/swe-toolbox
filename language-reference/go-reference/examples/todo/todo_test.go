// Demonstrated in this file:
// Create a temporary file (REF01)

package todo_test

import (
	"os"
	"testing"

	"github.com/RockLikeAmadeus/go-cli-projects/todo"
)

// TestAdd tests the Add method of the List type
func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, received %q", taskName, l[0].Task)
	}
}

// TestComplete tests the Complete method of the List type
func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, received %q", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("Task is already completed at creation: %q", taskName)
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("Completing task failed: %q", taskName)
	}
}

// TestDelete tests the Delete method of the List type
func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, received %q", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected list length %d, received %d", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, received %q", tasks[2], l[1].Task)
	}
}

// TestSaveGet tests the Save and Get methods of the List type
func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, received %q", taskName, l1[0].Task)
	}

	// Not a fan of this in a unit test, but this is how the book does it
	tf, err := os.CreateTemp("", "") // REF01

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task", l1[0].Task, l2[0].Task)
	}
}
