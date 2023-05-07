package fslice

type FilterFunction[T any] func(v T, i int, arr []T) bool

// Filter takes a function that operates on every element of the slice
// and returns only those elements of the slice for which the function returned true
func (fs Fslice[T]) Filter(fn FilterFunction[T]) Fslice[T] {
	r := make([]T, 0, len(fs))
	for i, v := range fs {
		if fn(v, i, fs) {
			r = append(r, v)
		}
	}
	return Fslice[T](r)
}
