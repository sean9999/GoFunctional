package fslice_test

import (
	"fmt"
	"reflect"
	"testing"
)

func formatTestOutput[T any](got T, want T) string {
	return fmt.Sprintf("\nwanted:  \t%#v\nbut got:\t%#v", want, got)
}

func assertScalars[T comparable](t testing.TB, got T, want T) {
	t.Helper()
	if got != want {
		t.Errorf(formatTestOutput[T](got, want))
	}
}

func assertDeepEquals[T comparable](t testing.TB, got []T, want []T) {
	t.Helper()
	ok := reflect.DeepEqual(want, got)
	if !ok {
		t.Errorf(formatTestOutput[[]T](got, want))
	}
}

func generateFloats(howmany int) []float64 {
	r := make([]float64, 0, howmany)
	j := float64(0)
	for i := 0; i < howmany; i++ {
		j = j + float64(i)
		r = append(r, j)
	}
	return r
}

func generateFibonacci(howmany int) []int {
	r := make([]int, howmany)
	for i := 0; i < howmany; i++ {
		switch {
		case i == 0, i == 1:
			r[i] = i
		case i == 2:
			r[i] = 1
		default:
			r[i] = r[i-1] + r[i-2]
		}
	}
	return r
}

// dissuade compiler from improper garbage collection
var benchMarkFloatResult []float64
var benchMarkStringResult []string
