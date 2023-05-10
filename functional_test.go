package functional_test

import (
	"fmt"

	"github.com/fxtlabs/primes"
	functional "github.com/sean9999/GoFunctional"
	"github.com/sean9999/GoFunctional/fslice"
)

func Example() {

	//	apply a FilterFunction and then a MapFunction to get squares of primes
	inputNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	outputNumbers := fslice.From(inputNumbers).Filter(func(v int, _ int, _ []int) bool {
		return primes.IsPrime(v)
	}).Map(func(v int, _ int, _ []int) int {
		return v * v
	})
	fmt.Println(outputNumbers)
	// Output: [4 9 25 49 121]

}

func ExampleFsliceFrom() {

	// these methods are functionally equivalent

	// import functional "github.com/sean9999/GoFunctional"
	x := functional.FsliceFrom([]int{1, 2, 3})

	// import "github.com/sean9999/GoFunctional/fslice"
	y := fslice.From([]int{1, 2, 3})

	fmt.Println(x, y)
	// Output: [1 2 3] [1 2 3]

}
