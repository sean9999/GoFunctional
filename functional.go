// Package functional implements Functional Programming capabilities for common go data structures.
// The root package offers nothing except for some globals, some helper functions, and to act as a namespace for the sub-packages.
// Sub-packages include [pkg/github.com/sean9999/GoFunctional/fslice] . "fmap" is on the roadmap. "fstruct" is being considered.
//
// Many methods are chainable, and when possible return another fslice or fmap, rather than the underlying slice or map.
// This is what gives Go Functional is compasability, expressiveness, and lends it the signature Functional Programming style.
// enhancing composability and expressiveness,
// which is what gives functional programming it's signature style.
//
// While for most use-cases Go Functional is _essentially_ a zero-cost abstraction, it is not _fundamentally_ so.
// There are use cases where the trade-off between expressiveness and performance is not acceptable.
// Use the right tool for the right job. Run the included benchmarks for
// performance characteristics.
package functional

import (
	"github.com/sean9999/GoFunctional/fslice"
)

// Version is the version of the module Go Functional [github.com/sean9999/GoFunctional]
const Version = "v0.0.3"

// convenience for [fslice.From]
func FsliceFrom[T comparable](inputSlice []T) fslice.MethodSet[T] {
	fsliceWithData := fslice.From(inputSlice)
	return fsliceWithData
}

// convenience for [fslice.New]
func NewFslice[T comparable](length, capacity int) fslice.MethodSet[T] {
	fsliceZeroed := fslice.New[T](length, capacity)
	return fsliceZeroed
}
