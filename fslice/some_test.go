package fslice_test

import (
	"fmt"
	"strings"

	"testing"

	"github.com/sean9999/GoFunctional/fslice"
)

func ExampleFslice_Some() {

	// at least one element is a multiple of 3
	inputSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	isMultipleOf3 := func(v int, _ int, _ []int) bool {
		return (v%3 == 0 && v >= 3)
	}
	result := fslice.From(inputSlice).Some(isMultipleOf3)
	fmt.Println(result)
	// Output: true

}

func TestSome(t *testing.T) {

	t.Run("at least one element is a multiple of 3", func(t *testing.T) {
		inputSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		isMultipleOf3 := func(v int, _ int, _ []int) bool {
			return (v%3 == 0 && v >= 3)
		}
		got := fslice.From(inputSlice).Some(isMultipleOf3)
		want := true
		if got != want {
			t.Errorf("isMultipleOf3 should have returned %v for at least one element in %v", want, inputSlice)
		}
	})

	t.Run("Filter out every 3rd element", func(t *testing.T) {
		inputSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var every3rd fslice.FilterFunction[int] = func(_ int, i int, _ []int) bool {
			return !(i%3 == 2 && i >= 2)
		}
		got := fslice.From(inputSlice).Filter(every3rd).ToSlice()
		want := []int{1, 2, 4, 5, 7, 8, 10}
		assertDeepEquals(t, got, want)
	})

	t.Run("contains at least one CAPITALISED word", func(t *testing.T) {
		inputSlice := []string{"all", "your", "BASE", "are", "belong", "to", "US"}
		isAllCaps := func(word string, _ int, _ []string) bool {
			upperCaseWord := strings.ToUpper(word)
			if word == upperCaseWord {
				return true
			}
			return false
		}
		want := true
		got := fslice.From(inputSlice).Some(isAllCaps)
		if got != want {
			t.Errorf("isAllCaps should have returned %v for at least one element in %v", want, inputSlice)
		}
	})

	t.Run("Contains one or more prime numbers", func(t *testing.T) {
		isPrime := func(n int, _ int, _ []int) bool {
			return IsPrime(n)
		}
		inputSlice := []int{0, 1, 2, 11, 13, 17, 23, 29, 31, 37, 43, 53, 61, 79, 87, 91, 101, 103, 107, 113, 433, 761, 25519, 65531}
		got := fslice.From(inputSlice).Some(isPrime)
		want := true
		if got != want {
			t.Errorf("inputSlice definitely contains a prime number")
		}
	})

	t.Run("Vacuous case with function that returns false", func(t *testing.T) {
		cull := func(_ int, _ int, _ []int) bool {
			return false
		}
		emptySlice := []int{}
		got := fslice.From(emptySlice).Some(cull)
		want := true
		if got != want {
			t.Errorf("Vacous case should be true")
		}
	})
	t.Run("Vacuous case with function that returns true", func(t *testing.T) {
		passThrough := func(_ float32, _ int, _ []float32) bool {
			return true
		}
		emptySlice := []float32{}
		got := fslice.From(emptySlice).Some(passThrough)
		want := true
		if got != want {
			t.Errorf("Vacous case should be true")
		}
	})

}
