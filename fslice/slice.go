package fslice

// Fslice is a slice with helpful methods in the Functional Programming style
type Fslice[T comparable] []T

// MethodSet is the interface Fslice implements
type MethodSet[T comparable] interface {
	Map(MapFunction[T]) Fslice[T]
	Filter(FilterFunction[T]) Fslice[T]
	Some(SomeFunction[T]) bool
	Every(EveryFunction[T]) bool
	Includes(T) bool
	Reduce(ReduceFunction[T], T) T
	Sort(SortFunction[T]) Fslice[T]
	ToSlice() []T
}

// From is a constructor. It wraps an existing slice,
// returning an Fslice[T] satisfying MethodSet[T]
func From[T comparable](inputSlice []T) MethodSet[T] {
	functionalSlice := Fslice[T](inputSlice)
	return functionalSlice
}

// New is an initializer. It pre-allocates a slice with size and capacity set, and wraps it,
// returning an Fslice[T] satisfying MethodSet[T]
func New[T comparable](size, capacity int) MethodSet[T] {
	// pre-allocate
	return Fslice[T](make([]T, size, capacity))
}

// ToSlice returns the underlying slice
func (fs Fslice[T]) ToSlice() []T {
	return fs
}
