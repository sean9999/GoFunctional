package fslice

func (fs Fslice[T]) Includes(x T) bool {

	for _, v := range fs {
		if v == x {
			return true
		}
	}

	return false
}
