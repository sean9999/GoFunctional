package fslice_test

import (
	"testing"

	"github.com/sean9999/GoFunctional/fslice"
)

func TestReduce(t *testing.T) {

	t.Run("sum integers", func(t *testing.T) {
		sum := func(a, b int32) int32 {
			return a + b
		}
		inputSlize := []int32{1000, 100, 10, 1}
		want := int32(1113)
		got := fslice.From(inputSlize).Reduce(sum, 0)
		assertScalars(t, got, want)
	})

}
