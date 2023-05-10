// Package functional implements Functional Programming capabilities for common go data structures.
// The root package offers nothing except for some globals, some helper functions, and a way to group the subpackages topgether
// Sub-packages include flsice and fmap, which operate on slice and maps respectively
//
// Many methods are chainable, enhancing composability and expressiveness, which is what gives functional programming it's signature style.
//
// Go Functional provides a near-zero cost abstraction for a great many use-cases, but will
// not do so for all. Use the right tool for the right job. Run the included benchmarks for
// performance characteristics.
package functional

import (
	"github.com/sean9999/GoFunctional/fslice"
)

const Version = "v0.0.1"

// convenience for fslice.From()
func FsliceFrom[T comparable](inputSlice []T) fslice.MethodSet[T] {
	fsliceWithData := fslice.From(inputSlice)
	return fsliceWithData
}

// convenience for fslice.New()
func NewFslice[T comparable](length, capacity int) fslice.MethodSet[T] {
	fsliceZeroed := fslice.New[T](length, capacity)
	return fsliceZeroed
}
