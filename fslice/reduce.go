package fslice

// ReduceFunction operates on all elements, feeding the value of one iteration into the next
type ReduceFunction[T comparable] func(a T, b T) T

// Reduce takes a ReduceFunction and a seed, which is the initial value fed into the user-supplied ReduceFunction
//
//	// BUG(sean9999): Seed value should be optional and should not be strictly typed
//	// see: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/reduce#edge_cases
func (fs Fslice[T]) Reduce(fn ReduceFunction[T], seed T) T {
	r := seed
	for _, el := range fs {
		r = fn(r, el)
	}
	return r
}
