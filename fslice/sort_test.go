package fslice_test

import (
	"testing"

	"github.com/sean9999/GoFunctional/fslice"
)

func TestSort(t *testing.T) {

	t.Run("ascending", func(t *testing.T) {
		ascending := func(a, b int32) int {
			if a < b {
				return 1
			}
			if a > b {
				return -1
			}
			//	this would be a terser way: return (b - a)
			return 0
		}
		inputSlize := []int32{101, 60, 1, 70, 2, 3, 4, 5, 97, 23, 50}
		want := []int32{1, 2, 3, 4, 5, 23, 50, 60, 70, 97, 101}
		got := fslice.From(inputSlize).Sort(ascending).ToSlice()
		assertDeepEquals(t, got, want)
	})

}
