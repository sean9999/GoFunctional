package fslice_test

import (
	"fmt"
	"reflect"
	"strings"

	"testing"

	"github.com/sean9999/GoFunctional/fslice"
)

func assertMapDeepEquals[T comparable](t testing.TB, inputSlice []T, want []T, mapFn fslice.MapFunction[T]) {
	t.Helper()
	got := fslice.New(inputSlice).Map(mapFn).ToSlice()
	slicesMatch := reflect.DeepEqual(want, got)
	if !slicesMatch {
		t.Errorf("wanted %#v but got %#v", want, got)
	}
}

func TestMap(t *testing.T) {

	t.Run("Five Integers Doubled", func(t *testing.T) {
		doubleThem := fslice.MapFunction[int](func(v int, _ int, _ []int) int { return v * 2 })
		fiveIntegers := []int{1, 2, 3, 4, 5}
		exptectedResult := []int{2, 4, 6, 8, 10}
		assertMapDeepEquals(t, fiveIntegers, exptectedResult, doubleThem)
	})

	t.Run("Identity case against a furnished slice", func(t *testing.T) {
		var passThrough fslice.MapFunction[int] = func(v int, _ int, _ []int) int {
			return v
		}
		assertMapDeepEquals(t, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, passThrough)
	})

	t.Run("Identity case against an empty slice", func(t *testing.T) {
		var passThrough fslice.MapFunction[float64] = func(v float64, _ int, _ []float64) float64 {
			return v
		}
		assertMapDeepEquals(t, []float64{}, []float64{}, passThrough)
	})

	t.Run("Capitalize Hello World", func(t *testing.T) {
		capitalize := func(word string, _ int, _ []string) string {
			return strings.ToUpper(word)
		}
		assertMapDeepEquals(t, []string{"Hello", "world"}, []string{"HELLO", "WORLD"}, capitalize)
	})

}

// prevent compiler from throughing away mem between results
var benchMarkResult []float64

func BenchmarkMap(b *testing.B) {

	lengths := []int{10, 100, 1_000, 10_000, 100_000}

	b.Run("Identity", func(b *testing.B) {

		var passThrough fslice.MapFunction[float64] = func(v float64, _ int, _ []float64) float64 {
			return v
		}

		for _, thisLength := range lengths {

			inputSlice := generateFloats(thisLength)

			b.Run(fmt.Sprintf("Functional_%d", thisLength), func(b *testing.B) {
				thisBenchMarkResult := make([]float64, 0, thisLength)

				for i := 0; i < b.N; i++ {
					thisBenchMarkResult = fslice.New(inputSlice).Map(passThrough).ToSlice()
				}
				benchMarkResult = thisBenchMarkResult
			})

			b.Run(fmt.Sprintf("Bare_%d", thisLength), func(b *testing.B) {
				thisBenchMarkResult := make([]float64, 0, thisLength)
				for i := 0; i < b.N; i++ {
					returnSlice := make([]float64, 0, thisLength)
					for _, v := range inputSlice {
						returnSlice = append(returnSlice, v)
					}
					thisBenchMarkResult = returnSlice
				}
				benchMarkResult = thisBenchMarkResult
			})

		}

	})

	b.Run("Fibonacci", func(b *testing.B) {

		for _, thisLength := range lengths {

			inputSlice := generateFloats(thisLength)
			var fib fslice.MapFunction[float64] = func(v float64, i int, arr []float64) float64 {
				r := float64(0)
				switch {
				case i == 0, i == 1:
					r = float64(i)
				default:
					r = arr[i-1] + arr[i-2]
				}
				return r
			}

			b.Run(fmt.Sprintf("Functional_%d", thisLength), func(b *testing.B) {
				thisBenchMarkResult := make([]float64, 0, thisLength)

				for i := 0; i < b.N; i++ {
					thisBenchMarkResult = fslice.New(inputSlice).Map(fib).ToSlice()
				}
				benchMarkResult = thisBenchMarkResult
			})

			b.Run(fmt.Sprintf("Bare_%d", thisLength), func(b *testing.B) {
				thisBenchMarkResult := make([]float64, 0, thisLength)
				for i := 0; i < b.N; i++ {
					returnSlice := make([]float64, 0, thisLength)
					for i, v := range inputSlice {
						switch {
						case i == 0, i == 1:
							returnSlice = append(returnSlice, float64(i))
						default:
							returnSlice = append(returnSlice, inputSlice[i-1]+inputSlice[i-2])
						}
						returnSlice = append(returnSlice, v)
					}
					thisBenchMarkResult = returnSlice
				}
				benchMarkResult = thisBenchMarkResult
			})

		}

	})

	b.Run("Lorem Ipsum to Uppercase", func(b *testing.B) {

	})

}

/*
func BenchmarkMapIdentityFunctional_10(b *testing.B) {
	floats, _ := generate10Floats()
	var passThrough fslice.MapFunction[float64] = func(v float64, _ int, _ []float64) float64 {
		return v
	}
	for n := 0; n < b.N; n++ {
		benchMarkResult = fslice.New(floats).Map(square).UnderlyingSlice()
	}
}

func BenchmarkMapIdentityBare_10(b *testing.B) {
	const sliceLength = 10
	floats, _ := generate10Floats()
	for n := 0; n < b.N; n++ {
		r := make([]float64, 0, sliceLength)
		for _, val := range floats {
			r = append(r, val)
		}
		benchMarkResult = r
	}
}

func BenchmarkMapIdentityFunctional_100(b *testing.B) {
	floats, _ := generate100Floats()
	var passThrough fslice.MapFunction[float64] = func(v float64, _ int, _ []float64) float64 {
		return v
	}
	for n := 0; n < b.N; n++ {
		benchMarkResult = fslice.New(floats).Map(square).UnderlyingSlice()
	}
}

func BenchmarkMapIdentityBare_100(b *testing.B) {
	const sliceLength = 100
	floats, _ := generate10Floats()
	for n := 0; n < b.N; n++ {
		r := make([]float64, 0, sliceLength)
		for _, val := range floats {
			r = append(r, val)
		}
		benchMarkResult = r
	}
}

func BenchmarkMapIdentityFunctional_1000(b *testing.B) {
	floats, _ := generate1000Floats()
	var passThrough fslice.MapFunction[float64] = func(v float64, _ int, _ []float64) float64 {
		return v
	}
	for n := 0; n < b.N; n++ {
		benchMarkResult = fslice.New(floats).Map(square).UnderlyingSlice()
	}
}

func BenchmarkMapIdentityBare_1000(b *testing.B) {
	const sliceLength = 1000
	floats, _ := generate10Floats()
	for n := 0; n < b.N; n++ {
		r := make([]float64, 0, sliceLength)
		for _, val := range floats {
			r = append(r, val)
		}
		benchMarkResult = r
	}
}

func BenchmarkMapIdentityFunctional_10000(b *testing.B) {
	floats, _ := generate10000Floats()
	var passThrough fslice.MapFunction[float64] = func(v float64, _ int, _ []float64) float64 {
		return v
	}
	for n := 0; n < b.N; n++ {
		benchMarkResult = fslice.New(floats).Map(square).UnderlyingSlice()
	}
}

func BenchmarkMapIdentityBare_10000(b *testing.B) {
	const sliceLength = 10_000
	floats, _ := generate10Floats()
	for n := 0; n < b.N; n++ {
		r := make([]float64, 0, sliceLength)
		for _, val := range floats {
			r = append(r, val)
		}
		benchMarkResult = r
	}
}

func BenchmarkMapIncrementorFunctional_10(b *testing.B) {
	floats, _ := generate10Floats()
	var incrementor fslice.MapFunction[float64] = func(v float64, i int, arr []float64) float64 {
		lastValue := float64(0)
		if i > 0 {
			lastValue = arr[i-1]
		}
		return lastValue + float64(i) + v
	}
	for n := 0; n < b.N; n++ {
		benchMarkResult = fslice.New(floats).Map(incrementor).UnderlyingSlice()
	}
}

func BenchmarkMapIncrementorFunctional_100(b *testing.B) {
	floats, _ := generate100Floats()
	var incrementor fslice.MapFunction[float64] = func(v float64, i int, arr []float64) float64 {
		lastValue := float64(0)
		if i > 0 {
			lastValue = arr[i-1]
		}
		return lastValue + float64(i) + v
	}
	for n := 0; n < b.N; n++ {
		benchMarkResult = fslice.New(floats).Map(incrementor).UnderlyingSlice()
	}
}

func BenchmarkMapIncrementorFunctional_1000(b *testing.B) {
	floats, _ := generate1000Floats()
	var incrementor fslice.MapFunction[float64] = func(v float64, i int, arr []float64) float64 {
		lastValue := float64(0)
		if i > 0 {
			lastValue = arr[i-1]
		}
		return lastValue + float64(i) + v
	}
	for n := 0; n < b.N; n++ {
		benchMarkResult = fslice.New(floats).Map(incrementor).UnderlyingSlice()
	}
}

func BenchmarkMapIncrementorFunctional_10000(b *testing.B) {
	floats, _ := generate10000Floats()
	var incrementor fslice.MapFunction[float64] = func(v float64, i int, arr []float64) float64 {
		lastValue := float64(0)
		if i > 0 {
			lastValue = arr[i-1]
		}
		return lastValue + float64(i) + v
	}
	for n := 0; n < b.N; n++ {
		benchMarkResult = fslice.New(floats).Map(incrementor).UnderlyingSlice()
	}
}

func BenchmarkMapIncrementorBare_10(b *testing.B) {
	const sliceLength = 10
	floats, _ := generate10Floats()
	for n := 0; n < b.N; n++ {
		r := make([]float64, 0, sliceLength)
		lastValue := float64(0)
		for i, v := range floats {
			if i > 0 {
				lastValue = floats[i-1]
			}
			r = append(r, v+lastValue+float64(i))
		}
		benchMarkResult = r
	}
}

func BenchmarkMapIncrementorBare_100(b *testing.B) {
	const sliceLength = 100
	floats, _ := generate100Floats()
	for n := 0; n < b.N; n++ {
		r := make([]float64, 0, sliceLength)
		lastValue := float64(0)
		for i, v := range floats {
			if i > 0 {
				lastValue = floats[i-1]
			}
			r = append(r, v+lastValue+float64(i))
		}
		benchMarkResult = r
	}
}

func BenchmarkMapIncrementorBare_1000(b *testing.B) {
	const sliceLength = 1000
	floats, _ := generate1000Floats()
	for n := 0; n < b.N; n++ {
		r := make([]float64, 0, sliceLength)
		lastValue := float64(0)
		for i, v := range floats {
			if i > 0 {
				lastValue = floats[i-1]
			}
			r = append(r, v+lastValue+float64(i))
		}
		benchMarkResult = r
	}
}

func BenchmarkMapIncrementorBare_10000(b *testing.B) {
	const sliceLength = 10_000
	floats, _ := generate10000Floats()
	for n := 0; n < b.N; n++ {
		r := make([]float64, 0, sliceLength)
		lastValue := float64(0)
		for i, v := range floats {
			if i > 0 {
				lastValue = floats[i-1]
			}
			r = append(r, v+lastValue+float64(i))
		}
		benchMarkResult = r
	}
}
*/
