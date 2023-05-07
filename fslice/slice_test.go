package fslice_test

import (
	"reflect"
	"testing"
)

func assertDeepEquals[T comparable](t testing.TB, got []T, want []T) {
	t.Helper()
	ok := reflect.DeepEqual(want, got)
	if !ok {
		t.Errorf("wanted %#v but got %#v", want, got)
	}
}

// @todo: make this less naive
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
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