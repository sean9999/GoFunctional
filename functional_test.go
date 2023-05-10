package functional_test

import (
	"fmt"

	"github.com/sean9999/GoFunctional/fslice"
)

func ExampleFslice_Filter() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	onlyEven := func(n int, i int, arr []int) bool {
		return (n%2 == 0)
	}
	filteredNumbers := fslice.From(numbers).Filter(onlyEven).ToSlice()
	fmt.Println(filteredNumbers)
	// Output: [2 4 6 8]
}
