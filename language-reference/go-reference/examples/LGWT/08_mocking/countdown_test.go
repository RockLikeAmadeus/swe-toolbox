/*
From https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking
We don't want to spend a long time with code that will theoretically work after some hacking
because that's often how developers fall down rabbit holes. It's an important skill to be able
to slice up requirements as small as you can so you can have working software.


From the same source:

If your mocking code is becoming complicated or you are having to mock out lots of things to test
something, you should listen to that bad feeling and think about your code. Usually it is a sign of
- The thing you are testing is having to do too many things (because it has too many dependencies
	to mock)
	- Break the module apart so it does less
- Its dependencies are too fine-grained
	- Think about how you can consolidate some of these dependencies into one meaningful module
- Your test is too concerned with implementation details
	- Favour testing expected behaviour rather than the implementation

Normally a lot of mocking points to bad abstraction in your code.

What people see here is a weakness in TDD but it is actually a strength, more often than not poor
test code is a result of bad design or put more nicely, well-designed code is easy to test.


From the same source:

- I feel like if a test is working with more than 3 mocks then it is a red flag - time for a
	rethink on the design
- Use spies with caution. Spies let you see the insides of the algorithm you are writing which
	can be very useful but that means a tighter coupling between your test code and the
	implementation. Be sure you actually care about these details if you're going to spy on them


From the same source:

Mocking requires no magic and is relatively simple; using a framework can make mocking seem more
complicated than it is.
*/
package main

import (
	"bytes"
	"testing"
	"reflect"
	"time"
)

// *Spies* are a kind of *mock* which can record how a dependency is used, such as by examining
// the passed-in arguments, or the number of times it has been called, etc.
/*type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}*/

// We could just use the above implementation, but this doesn't allow us to test everything about
// the behavior of our code, such as the order of operations. In this case, we can create a single
// spy object for both operations.
type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

// Another Spy to test our configurable sleeper
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// The Tests:
func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		// Any time we're printing or writing to anything, this is a good pre-made mock for us
		buffer := &bytes.Buffer{} // struct literal
		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		// Backtick syntax lets us define a raw string
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	// Here when we instantiate ConfigurableSleeper, we pass it the Sleep() method on our spy so
	// that the ConfigurableSleeper's `sleep` property points to this method
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}