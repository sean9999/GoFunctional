package fslice

// ToSlice returns the underlying slice unadorned
func (fs Fslice[T]) ToSlice() []T {
	return fs
}
