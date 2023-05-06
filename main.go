package main

import "fmt"

// Flice method definitions
type FSlice[T any] []T

// method signatures
type MapFunction[T any] func(v T, i int, arr []T) T
type FilterFunction[T any] func(v T, i int, arr []T) bool
type SomeFunction[T any] func(v T, i int, arr []T) bool
type EveryFunction[T any] func(v int, i int, arr []T) bool

// interface
type MethodSet[T any] interface {
	Map(MapFunction[T]) []T
	Filter(FilterFunction[T]) bool
	Some(SomeFunction[T]) bool
	Every(EveryFunction[T]) bool
}

func NewFSlice[T any](inputSlice []T) FSlice[T] {
	return FSlice[T](inputSlice)
}

func (fs FSlice[T]) Map(fn func(v T, i int, self []T) T) FSlice[T] {
	var r []T
	for i, v := range fs {
		r = append(r, fn(v, i, fs))
	}
	return FSlice[T](r)
}

func (fs FSlice[T]) Filter(fn func(v T, i int, self []T) bool) FSlice[T] {
	var r FSlice[T]

	for i, v := range fs {
		if fn(v, i, fs) {
			r = append(r, v)
		}
	}

	return r
}

func main() {

	slice1 := []int{1, 2, 3, 4, 5}
	var doubleIt MapFunction[int] = func(v int, _ int, _ []int) int {
		return v * 2
	}
	var everySecond FilterFunction[int] = func(_ int, i int, _ []int) bool {
		return (i%2 == 0)
	}

	doubles := FSlice[int](slice1).Map(doubleIt)
	staggered := FSlice[int](slice1).Filter(everySecond).Map(doubleIt).Map(doubleIt)

	fmt.Println(doubles)
	fmt.Println(staggered)
}
