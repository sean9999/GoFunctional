package fslice

// MapFunction operates on every element of a Slice and returns an element of the same type.
// It is passed into [Fslice.Map]
type MapFunction[T comparable] func(v T, i int, arr []T) T

// Map takes a MapFunction
// and returns a new Slice with the same number of elements,
// those elements being of the same type
func (fs Fslice[T]) Map(fn MapFunction[T]) Fslice[T] {
	r := make([]T, 0, len(fs))
	for i, v := range fs {
		r = append(r, fn(v, i, fs))
	}
	return Fslice[T](r)
}
