package fslice_test

import (
	"fmt"
	"testing"

	"github.com/sean9999/GoFunctional/fslice"
)

func ExampleFslice_Every() {

	inputSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	isEven := func(v int, _ int, _ []int) bool {
		return (v%2 == 0)
	}
	isTrueForEveryElement := fslice.New(inputSlice).Every(isEven)
	fmt.Println(isTrueForEveryElement)
	//	Output: false

}

func TestEvery(t *testing.T) {

	t.Run("all integers are even", func(t *testing.T) {
		inputSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		evenInteger := func(v int, _ int, _ []int) bool {
			return (v%2 == 0)
		}
		got := fslice.New(inputSlice).Every(evenInteger)
		want := false
		assertScalars(t, got, want)
	})

	t.Run("all integers are below 1000", func(t *testing.T) {
		inputSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		onlyEvenIntegers := func(v int, _ int, _ []int) bool {
			return (v < 1000)
		}
		got := fslice.New(inputSlice).Every(onlyEvenIntegers)
		want := true
		assertScalars(t, got, want)
	})

}
