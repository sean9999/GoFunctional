package fslice_test

import (
	"fmt"
	"testing"

	"github.com/sean9999/GoFunctional/fslice"
)

func ExampleFslice_Every() {

	//	is every element of this slice even?
	inputSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	isEven := func(v int, _ int, _ []int) bool {
		return (v%2 == 0)
	}
	answer := fslice.From(inputSlice).Every(isEven)

	fmt.Println(answer)
	//	Output: false

}

func TestEvery(t *testing.T) {

	t.Run("all integers are even", func(t *testing.T) {
		inputSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		evenInteger := func(v int, _ int, _ []int) bool {
			return (v%2 == 0)
		}
		got := fslice.From(inputSlice).Every(evenInteger)
		want := false
		assertScalars(t, got, want)
	})

	t.Run("all integers are below 1000", func(t *testing.T) {
		inputSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		onlyEvenIntegers := func(v int, _ int, _ []int) bool {
			return (v < 1000)
		}
		got := fslice.From(inputSlice).Every(onlyEvenIntegers)
		want := true
		assertScalars(t, got, want)
	})

}
