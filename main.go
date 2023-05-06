package main

import "fmt"

// Flice method definitions
type FSlice[T comparable] []T

// method signatures
type MapFunction[T any] func(v T, i int, arr []T) T
type FilterFunction[T any] func(v T, i int, arr []T) bool
type SomeFunction[T any] func(v T, i int, arr []T) bool
type EveryFunction[T any] func(v T, i int, arr []T) bool

//type IncludesFunction[T any] func(v T) bool

// interface
type MethodSet[T comparable] interface {
	Map(MapFunction[T]) FSlice[T]
	Filter(FilterFunction[T]) FSlice[T]
	Some(SomeFunction[T]) bool
	Every(EveryFunction[T]) bool
	Includes(T) bool
	Raw() []T
}

func NewFSlice[T comparable](inputSlice []T) FSlice[T] {
	return FSlice[T](inputSlice)
}

// Map takes a function that operates on every element of the slice
// returning a new slice with the same number of elements
func (fs FSlice[T]) Map(fn MapFunction[T]) FSlice[T] {
	var r []T
	for i, v := range fs {
		r = append(r, fn(v, i, fs))
	}
	return FSlice[T](r)
}

// Filter takes a function that operates on every element of the slice
// and returns only those elements of the slice for which the function returned true
func (fs FSlice[T]) Filter(fn FilterFunction[T]) FSlice[T] {
	var r []T
	for i, v := range fs {
		if fn(v, i, fs) {
			r = append(r, v)
		}
	}
	return FSlice[T](r)
}

// Some takes a function that operates on as many elements of the slice
// as it takes to get a return value of true. It returns true in that case
// and false otherwise
func (fs FSlice[T]) Some(fn SomeFunction[T]) bool {
	var r bool
	for i, v := range fs {
		r = fn(v, i, fs)
		if r {
			break
		}
	}
	return r
}

func (fs FSlice[T]) Every(fn EveryFunction[T]) bool {
	var r bool
	for i, v := range fs {
		r = fn(v, i, fs)
		if !r {
			break
		}
	}
	return r
}

func (fs FSlice[T]) Includes(x T) bool {

	for _, v := range fs {
		if v == x {
			return true
		}
	}

	return false
}

// Raw returns the underlying slice unadorned
func (fs FSlice[T]) Raw() []T {
	return fs
}

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	var doubleIt MapFunction[int] = func(v int, _ int, _ []int) int {
		return v * 2
	}
	var everySecond = func(_ int, i int, _ []int) bool {
		return (i%2 == 0)
	}

	doubles := FSlice[int](slice1).Map(doubleIt).Raw()
	staggered := FSlice[int](slice1).Filter(everySecond).Map(doubleIt).Map(doubleIt)

	hasOdds := FSlice[int](slice1).Some(func(v int, _ int, _ []int) bool {
		return (v%2 == 1)
	})

	hasDoubleDigits := FSlice[int](slice1).Some(func(v int, _ int, _ []int) bool {
		return (v > 9)
	})

	hasSeven := FSlice[int](slice1).Includes(7)

	fmt.Printf("%#v\n", doubles)
	fmt.Printf("%#v\n", staggered)
	fmt.Printf("hasOdds = %#v\n", hasOdds)
	fmt.Printf("hasDoubleDigits = %#v\n", hasDoubleDigits)
	fmt.Printf("hasSeven = %#v\n", hasSeven)
}
