package fslice_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	lorem "github.com/drhodes/golorem"
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

func generateLoremIpsum(numwords int) []string {
	text := lorem.Sentence(numwords, numwords)
	return strings.Split(text, " ")
}

// dissuade compiler from improper garbage collecting
var benchMarkFloatResult []float64
var benchMarkStringResult []string
