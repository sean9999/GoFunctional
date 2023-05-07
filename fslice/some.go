package fslice

type SomeFunction[T any] func(v T, i int, arr []T) bool

// Some takes a function that operates on as many elements of the slice
// as it takes to get a return value of true. It returns true in that case
// and false otherwise
func (fs Fslice[T]) Some(fn SomeFunction[T]) bool {
	r := true
	for i, v := range fs {
		r = fn(v, i, fs)
		if r {
			break
		}
	}
	return r
}
