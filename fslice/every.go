package fslice

// EveryFunction operates on an element of a slice and returns true or false
type EveryFunction[T comparable] func(v T, i int, arr []T) bool

// Every returns true if its EveryFunction returns true for every element
func (fs Fslice[T]) Every(fn EveryFunction[T]) bool {
	r := true
	for i, v := range fs {
		r = fn(v, i, fs)
		if !r {
			break
		}
	}
	return r
}
