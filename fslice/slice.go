package fslice

// Fslice is a slice with Functional Programming methods, as defined in the MethodSet interface
type Fslice[T comparable] []T

// implements
type MethodSet[T comparable] interface {
	Map(MapFunction[T]) Fslice[T]
	Filter(FilterFunction[T]) Fslice[T]
	Some(SomeFunction[T]) bool
	Every(EveryFunction[T]) bool
	Includes(T) bool
	ToSlice() []T
}

// constructor
func New[T comparable](inputSlice []T) MethodSet[T] {
	functionalSlice := Fslice[T](inputSlice)
	return functionalSlice
}

// return the underlying slice
func (fs Fslice[T]) ToSlice() []T {
	return fs
}
