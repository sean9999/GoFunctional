package fslice

// ReduceFunction operates on all elements, feeding the value of one iteration into the next
type ReduceFunction[T comparable] func(a T, b T) T

// Reduce takes a ReduceFunction and a seed, which is the intial value fed into the user-supplied ReduceFunction
func (fs Fslice[T]) Reduce(fn ReduceFunction[T], seed T) T {
	r := seed
	for _, el := range fs {
		r = fn(r, el)
	}
	return r
}
