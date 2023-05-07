package fslice_test

import (
	"strings"

	"testing"

	"github.com/sean9999/GoFunctional/fslice"
)

func TestFilter(t *testing.T) {

	t.Run("only even integers", func(t *testing.T) {
		inputSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		onlyEvenIntegers := func(v int, _ int, _ []int) bool {
			return (v%2 == 0 && v > 1)
		}
		got := fslice.New(inputSlice).Filter(onlyEvenIntegers).ToSlice()
		want := []int{2, 4, 6, 8, 10, 12}
		assertDeepEquals(t, got, want)
	})

	t.Run("Filter out every 3rd element", func(t *testing.T) {
		inputSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var every3rd fslice.FilterFunction[int] = func(_ int, i int, _ []int) bool {
			return !(i%3 == 2 && i >= 2)
		}
		got := fslice.New(inputSlice).Filter(every3rd).ToSlice()
		want := []int{1, 2, 4, 5, 7, 8, 10}
		assertDeepEquals(t, got, want)
	})

	t.Run("Omit capitalized words", func(t *testing.T) {
		inputSlice := []string{"all", "your", "BASE", "are", "belong", "to", "US"}
		noCaps := func(word string, _ int, _ []string) bool {
			upperCaseWord := strings.ToUpper(word)
			if word == upperCaseWord {
				return false
			}
			return true
		}
		want := []string{"all", "your", "are", "belong", "to"}
		got := fslice.New(inputSlice).Filter(noCaps).ToSlice()
		assertDeepEquals(t, got, want)
	})

	t.Run("Omit prime numbers", func(t *testing.T) {
		outPrimes := func(n int, _ int, _ []int) bool {
			return !IsPrime(n)
		}
		inputSlice := []int{0, 1, 2, 11, 13, 17, 23, 29, 31, 37, 43, 53, 61, 79, 87, 91, 101, 103, 107, 113, 433, 761, 25519, 65531}
		got := fslice.New(inputSlice).Filter(outPrimes).ToSlice()
		want := []int{0, 1, 87, 91, 25519, 65531}
		assertDeepEquals(t, got, want)
	})

	t.Run("Vacuous case with false return", func(t *testing.T) {
		cull := func(_ int, _ int, _ []int) bool {
			return false
		}
		emptySlice := []int{}
		got := fslice.New(emptySlice).Filter(cull).ToSlice()
		want := emptySlice
		assertDeepEquals(t, got, want)
	})
	t.Run("Vacuous case with true return", func(t *testing.T) {
		passThrough := func(_ float32, _ int, _ []float32) bool {
			return true
		}
		emptySlice := []float32{}
		got := fslice.New(emptySlice).Filter(passThrough).ToSlice()
		want := emptySlice
		assertDeepEquals(t, got, want)
	})

}
