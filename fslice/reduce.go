package fslice

type ReduceFunction[T comparable] func(a T, b T) T

func (fs Fslice[T]) Reduce(fn ReduceFunction[T], seed T) T {

	return fs[0]
}
