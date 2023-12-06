package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("Expected %d but received % d, given %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// Careful: reflect.DeepEqual() is not type safe
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %d but received %d", want, got)
	}
}

func TestSumAllTails(t *testing.T) {
	// We can assign functions to variables, like JavaScript
	// Doing so here limits the scope of this helper and hides it, which we like
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %d but received %d", want, got)
		}
	}

	t.Run("Sums of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})
	t.Run("Sums of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
