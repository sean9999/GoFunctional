package fslice

type EveryFunction[T comparable] func(v T, i int, arr []T) bool

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
