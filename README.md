# Go Functional

Go Functional provides Functional Programming capabilities to Go.

[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/sean9999/GoFunctional/graphs/commit-activity)

[![Go Reference](https://pkg.go.dev/badge/github.com/sean9999/GoFunctional.svg)](https://pkg.go.dev/github.com/sean9999/GoFunctional)

[![Go Report Card](https://goreportcard.com/badge/github.com/sean9999/GoFunctional)](https://goreportcard.com/report/github.com/sean9999/GoFunctional)

[![Go version](https://img.shields.io/github/go-mod/go-version/sean9999/GoFunctional.svg)](https://github.com/sean9999/GoFunctional)

![Go Functional](chariot.jpg)

Read the [blog post here](https://www.seanmacdonald.ca/posts/go-functional): 

## Root package

The root package _functional_ acts as a namespace for sub-packages that target specific data structures, like slices and maps. It also exposes some common helpers useful for testing and benchmarks. You will rarely need to import it; rather, you will want to import a sub-package, such as Fslice.

## Sub Packages

### Fslice

Fslice provides a set of methods for operating on slices. Methods that could return a slice return an Fslice, making the system chainable. Among them are common functional programming methods, such as:

- Map() - takes a function that maps over every element, returning a transformed slice (chainable)
- Filter() - takes a function that filters a slice to a smaller subset (chainable)
- Includes() - returns true if a value was found in the slice
- Some() - takes a function that returns a boolean. Returns true early if the function returned true once
- Every() - takes a function that returns a boolean. Returns false early if the function returned false once
- ToSlice() - returns the underlying (non-functional) slice
- Reduce() - takes a function that behaves like an accumulator, returning a single value
- Sort() - takes a function that sorts a slice, and returns it (chainable)

Calling it is easy. It works like this:

```go
import (
	"fmt"
	"github.com/sean9999/GoFunctional/fslice"
)

//  Input
inputNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

//  a FilterFunction that returns true if the number is prime
isPrime := func(val int, _ int, _ []int) bool {
    return primes.IsPrime(val)
}

//  a MapFunction that squares a number
square := func(val int, _ int, _ []int) int {
    return val * val
}

//  a slice containing only the primes, squared 
outputNums := fslice.From(inputNums).Filter(isPrime).Map(square)

fmt.Println(outputNums)
// Output: [4 9 25 49 121]
```
Or you could put your functions inline, which is common in Functional Programming:

```go
outputNums := fslice.From(inputNums).Filter(func(n int, _ int, _ []int) bool {
    return primes.IsPrime(val)
}).Map(func(n int, _ int, _ []int) int {
    return n * n
})

fmt.Println(outputNums)
// Output: [4 9 25 49 121]
```

## Design Goals

- near-zero cost abstraction, with allowances made for readability and flexibility
- respectful of both Functional Programming and Go idioms
- well defined behaviour
- well tested
- well benchmarked, so that time and space complexity can be known and documented

## What's Next?

A similar set of methods for maps, which will be under a package called Fmap. There might be an opportunity to do something for structs as well. We'll see.

While generic, Go Functional is not fully polymorphic. I'd like to find an elegant way to provide this, using the concept of a Functor.

There is no reason Go Functional could not provide an elegant and idiomatic [set-builder](https://en.wikipedia.org/wiki/Set-builder_notation#In_programming_languages).

Benchmarks should be plotted so that performance characteristics can be seen visually, and time and space complexity metrics can be extrapolated.

Some system to measure the rate at which cyclomatic complexity increases as the chain of methods increases, would be useful.

Pull requests are most welcome.
