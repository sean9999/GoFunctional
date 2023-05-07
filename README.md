# Go Functional

Go Functional provides a Go package called Functional, which provides Functional Programming capabilities to Go's data types

## Design Goals

- near-zero cost abstraction, with allowances made for readability and flexibility
- respectful of both Functional Programming and Go idioms
- well defined behaviour
- well tested
- well benchmarked, so that time and space complexity can be known and documented

## Functional.Slice

Provides certain methods common to Functional Programming for operating on slices. The methods are chainable

- Slice.map() - takes a function that maps over every element, returning a transformed slice
- Slice.filter() - takes a function that filters a slice to a smaller subset
- Slice.includes() - returns true if a value was found in the slice
- Slice.some() - takes a function that returns a boolean. Returns true early if the function returned true once
- Slice.every() - takes a function that returns a boolean. Returns false early if the function returned false once
- Slice.Underlying() - returns the underlying slice


