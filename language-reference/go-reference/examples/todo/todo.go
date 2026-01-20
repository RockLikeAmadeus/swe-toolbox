// Demonstrated in this file:
// Pass each element of a slice as an individual argument to a variadic function (REF01)
// Remove an item at a specific index from a slice (REF02)
// Convert an object to JSON and write it to a file (REF03)
// Read from a JSON file and parse the contents (REF04)
// Determine if an error is a specific type of error (REF05)
// Define a type's custom string printing behavior (REF06)

package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// Item struct represents a ToDo item.
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of ToDo items.
type List []item

// Add creates a new todo item and appends it to the list.
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// Complete method marks a ToDo item as completed by
// setting Done = true and CompletedAt to the current time
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	// Adjusting for 0-based index (caller will expect 1-based)
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Delete method deletes a ToDo item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	// The ... syntax here unpacks the second slice's elements,
	// as append is a variadic function
	// REF01, REF02
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Save method encodes the List as JSON and saves it
// using the provided file name
// REF03
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name, decodes
// the JSON data and parses it into a List
// REF04
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) { // REF05
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

// String prints out a formatted list
// Implements the fmt.Stringer interface
// REF06 - NOTE, the String() method works better when using a value
// receiver, hence why we're not using a pointer receiver for this one method
func (l List) String() string {
	formatted := ""

	for k, t := range l {
		prefix := " "
		if t.Done {
			prefix = "X"
		}

		// Adjust the item number k to print numbers starting from 1
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}

	return formatted
}
