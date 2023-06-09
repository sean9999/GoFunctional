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
// performance characteristics, or view Go's cyclomatic complexity here: [https://goreportcard.com/report/github.com/sean9999/GoFunctional]
package functional
