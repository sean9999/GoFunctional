package fslice_test

import (
	"fmt"
	"reflect"
	"strings"

	"regexp"
	"testing"

	"github.com/sean9999/GoFunctional/fslice"
)

func assertMapDeepEquals[T comparable](t testing.TB, inputSlice []T, want []T, mapFn fslice.MapFunction[T]) {
	t.Helper()
	got := fslice.From(inputSlice).Map(mapFn).ToSlice()
	slicesMatch := reflect.DeepEqual(want, got)
	if !slicesMatch {
		t.Errorf("wanted %#v but got %#v", want, got)
	}
}

func ExampleFslice_Map() {

	// square the numbers
	squared := func(n int, _ int, _ []int) int {
		return n * n
	}
	squares := fslice.From([]int{1, 2, 3, 4, 5}).Map(squared).ToSlice()

	// convert every other word to SHOUTCASE
	lowerBase := []string{"all", "your", "base", "are", "belong", "to", "us"}
	shoutCaseEveryOther := func(word string, i int, _ []string) string {
		if i%2 == 1 {
			word = strings.ToUpper(word)
		}
		return word
	}
	shoutyBase := fslice.From(lowerBase).Map(shoutCaseEveryOther).ToSlice()

	fmt.Println(squares)
	fmt.Println(shoutyBase)

	// Output:
	// [1 4 9 16 25]
	// [all YOUR base ARE belong TO us]

}

func TestMap(t *testing.T) {

	t.Run("Five Integers Doubled", func(t *testing.T) {
		doubleThem := fslice.MapFunction[int](func(v int, _ int, _ []int) int { return v * 2 })
		fiveIntegers := []int{1, 2, 3, 4, 5}
		exptectedResult := []int{2, 4, 6, 8, 10, 11}
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

func BenchmarkMap(b *testing.B) {

	lengths := []int{10, 100, 1_000, 10_000, 100_000}

	for _, thisLength := range lengths {

		inputFloats := generateFloats(thisLength)
		inputStrings := generateLoremIpsum(thisLength)

		b.Run("Identity", func(b *testing.B) {

			var passThrough fslice.MapFunction[float64] = func(v float64, _ int, _ []float64) float64 {
				return v
			}

			b.Run(fmt.Sprintf("Functional_%d", thisLength), func(b *testing.B) {
				thisBenchMarkResult := make([]float64, 0, thisLength)

				for i := 0; i < b.N; i++ {
					thisBenchMarkResult = fslice.From(inputFloats).Map(passThrough).ToSlice()
				}
				benchMarkFloatResult = thisBenchMarkResult
			})

			b.Run(fmt.Sprintf("Bare_%d", thisLength), func(b *testing.B) {
				thisBenchMarkResult := make([]float64, 0, thisLength)
				for i := 0; i < b.N; i++ {
					returnSlice := make([]float64, 0, thisLength)
					for _, v := range inputFloats {
						returnSlice = append(returnSlice, v)
					}
					thisBenchMarkResult = returnSlice
				}
				benchMarkFloatResult = thisBenchMarkResult
			})

		})

		b.Run("Fibonacci", func(b *testing.B) {

			fib := func(_ float64, i int, arr []float64) float64 {
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
					thisBenchMarkResult = fslice.From(inputFloats).Map(fib).ToSlice()
				}
				benchMarkFloatResult = thisBenchMarkResult
			})

			b.Run(fmt.Sprintf("Bare_%d", thisLength), func(b *testing.B) {
				thisBenchMarkResult := make([]float64, 0, thisLength)
				for i := 0; i < b.N; i++ {
					returnSlice := make([]float64, 0, thisLength)
					for j := range inputFloats {
						switch {
						case j == 0, j == 1:
							returnSlice = append(returnSlice, float64(i))
						default:
							returnSlice = append(returnSlice, inputFloats[j-1]+inputFloats[j-2])
						}
					}
					thisBenchMarkResult = returnSlice
				}
				benchMarkFloatResult = thisBenchMarkResult
			})

		})

		b.Run("Convert some words to Uppercase", func(b *testing.B) {

			var vowelsToUpper fslice.MapFunction[string] = func(word string, i int, arr []string) string {
				re := regexp.MustCompile(`^[aeiouAEIOU]`)
				if re.MatchString(word) {
					word = strings.ToUpper(word)
				}
				return word
			}

			b.Run(fmt.Sprintf("Functional_%d", thisLength), func(b *testing.B) {
				thisBenchMarkResult := make([]string, 0, thisLength)
				for i := 0; i < b.N; i++ {
					thisBenchMarkResult = fslice.From(inputStrings).Map(vowelsToUpper).ToSlice()
				}
				benchMarkStringResult = thisBenchMarkResult
			})

			b.Run(fmt.Sprintf("Bare_%d", thisLength), func(b *testing.B) {
				thisBenchMarkResult := make([]string, 0, thisLength)
				for i := 0; i < b.N; i++ {
					thisBenchMarkResult := make([]string, 0, thisLength)
					for _, word := range inputStrings {
						re := regexp.MustCompile(`^[aeiouAEIOU]`)
						if re.MatchString(word) {
							thisBenchMarkResult = append(thisBenchMarkResult, strings.ToUpper(word))
						} else {
							thisBenchMarkResult = append(thisBenchMarkResult, word)
						}
					}
				}
				benchMarkStringResult = thisBenchMarkResult
			})

		})

	}

}
