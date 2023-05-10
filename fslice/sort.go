package fslice

import "sort"

// SortFunction compares two values from the slice and returns a negative value if (a > b), else positive or zero.
// Note that while Javascript makes a distinction between positive and zero values, Go does not.
// Go only asks if (a > b), and never asks if (a == b).
//
// see [https://pkg.go.dev/sort#IntSlice.Less]
type SortFunction[T comparable] func(a, b T) int

// Sort takes a SortFunction and returns a slice, sorted
func (fs Fslice[T]) Sort(fn SortFunction[T]) Fslice[T] {
	goSortFunction := func(i, j int) bool {
		//	return true if user-defined function returns a value greater than 0, else false
		return (fn(fs[i], fs[j]) > 0)
	}
	sort.Slice(fs, goSortFunction)
	return fs
}
