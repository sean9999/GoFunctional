package fslice_test

import (
	"fmt"
	"unicode/utf8"

	"strings"

	"testing"

	"github.com/sean9999/GoFunctional/fslice"
)

func ExampleFslice_Filter() {

	inputSlice := []string{"all", "your", "BASE", "are", "belong", "to", "US"}
	noShouting := func(word string, _ int, _ []string) bool {
		upperCaseWord := strings.ToUpper(word)
		if word == upperCaseWord {
			return false
		}
		return true
	}
	outputSlice := fslice.New(inputSlice).Filter(noShouting).ToSlice()
	fmt.Println(outputSlice)
	//	Output: [all your are belong to]

}

func TestFilter(t *testing.T) {

	t.Run("only even integers", func(t *testing.T) {
		inputSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		onlyEvenIntegers := func(v int, _ int, _ []int) bool {
			return (v%2 == 0 && v > 1)
		}
		got := fslice.New(inputSlice).Filter(onlyEvenIntegers).ToSlice()
		want := []int{2, 4, 6, 8, 10, 12}
		assertDeepEquals(t, got, want)
	})

	t.Run("Filter out every 3rd element", func(t *testing.T) {
		inputSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var outEvery3rd fslice.FilterFunction[int] = func(_ int, i int, _ []int) bool {
			return !(i%3 == 2 && i >= 2)
		}
		got := fslice.New(inputSlice).Filter(outEvery3rd).ToSlice()
		want := []int{1, 2, 4, 5, 7, 8, 10}
		assertDeepEquals(t, got, want)
	})

	t.Run("Omit capitalized words", func(t *testing.T) {
		inputSlice := []string{"all", "your", "BASE", "are", "belong", "to", "US"}
		noCaps := func(word string, _ int, _ []string) bool {
			upperCaseWord := strings.ToUpper(word)
			if word == upperCaseWord {
				return false
			}
			return true
		}
		want := []string{"all", "your", "are", "belong", "to"}
		got := fslice.New(inputSlice).Filter(noCaps).ToSlice()
		assertDeepEquals(t, got, want)
	})

	t.Run("Omit prime numbers", func(t *testing.T) {
		outPrimes := func(n int, _ int, _ []int) bool {
			return !IsPrime(n)
		}
		inputSlice := []int{0, 1, 2, 11, 13, 17, 23, 29, 31, 37, 43, 53, 61, 79, 87, 91, 101, 103, 107, 113, 433, 761, 25519, 65531}
		got := fslice.New(inputSlice).Filter(outPrimes).ToSlice()
		want := []int{0, 1, 87, 91, 25519, 65531}
		assertDeepEquals(t, got, want)
	})

	t.Run("Vacuous case with false return", func(t *testing.T) {
		cull := func(_ int, _ int, _ []int) bool {
			return false
		}
		emptySlice := []int{}
		got := fslice.New(emptySlice).Filter(cull).ToSlice()
		want := emptySlice
		assertDeepEquals(t, got, want)
	})
	t.Run("Vacuous case with true return", func(t *testing.T) {
		passThrough := func(_ float32, _ int, _ []float32) bool {
			return true
		}
		emptySlice := []float32{}
		got := fslice.New(emptySlice).Filter(passThrough).ToSlice()
		want := emptySlice
		assertDeepEquals(t, got, want)
	})

}

func BenchmarkFilter(b *testing.B) {

	lengths := []int{10, 100, 1_000, 10_000, 100_000}

	for _, thisLength := range lengths {

		inputFloats := generateFloats(thisLength)
		inputStrings := generateLoremIpsum(thisLength)

		b.Run("Identity", func(b *testing.B) {
			var passThrough fslice.FilterFunction[float64] = func(v float64, _ int, _ []float64) bool {
				return true
			}
			b.Run(fmt.Sprintf("Functional_%d", thisLength), func(b *testing.B) {
				thisBenchMarkResult := make([]float64, 0, thisLength)
				for i := 0; i < b.N; i++ {
					thisBenchMarkResult = fslice.New(inputFloats).Filter(passThrough)
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

		b.Run("Remove multiples of 5", func(b *testing.B) {
			var outFives fslice.FilterFunction[float64] = func(v float64, _ int, _ []float64) bool {
				return (int64(v)%5 == 0)
			}
			b.Run(fmt.Sprintf("Functional_%d", thisLength), func(b *testing.B) {
				thisBenchMarkResult := make([]float64, 0, thisLength)
				for i := 0; i < b.N; i++ {
					thisBenchMarkResult = fslice.New(inputFloats).Filter(outFives)
				}
				benchMarkFloatResult = thisBenchMarkResult
			})
			b.Run(fmt.Sprintf("Bare_%d", thisLength), func(b *testing.B) {
				thisBenchMarkResult := make([]float64, 0, thisLength)
				for i := 0; i < b.N; i++ {
					thisBenchMarkResult = []float64{}
					for _, n := range inputFloats {
						if int64(n)%5 != 0 {
							thisBenchMarkResult = append(thisBenchMarkResult, n)
						}
					}
				}
				benchMarkFloatResult = thisBenchMarkResult
			})
		})

		b.Run("Omit 4-letter words", func(b *testing.B) {
			no4LetterWords := func(word string, _ int, _ []string) bool {
				return utf8.RuneCountInString(word) != 4
			}
			b.Run(fmt.Sprintf("Functional_%d", thisLength), func(t *testing.B) {
				thisBenchMarkResult := make([]string, 0, thisLength)
				for i := 0; i < b.N; i++ {
					thisBenchMarkResult = fslice.New(inputStrings).Filter(no4LetterWords).ToSlice()
				}
				benchMarkStringResult = thisBenchMarkResult
			})
		})

		b.Run("remove duplicate words", func(b *testing.B) {
			noDuplicate := func(thisWord string, i int, arr []string) bool {
				ok := true
				for _, thatWord := range arr[:i] {
					if thisWord == thatWord {
						ok = false
						break
					}
				}
				return ok
			}
			b.Run(fmt.Sprintf("Functional_%d", thisLength), func(t *testing.B) {
				thisBenchMarkResult := make([]string, 0, thisLength)
				for i := 0; i < b.N; i++ {
					thisBenchMarkResult = fslice.New(inputStrings).Filter(noDuplicate).ToSlice()
				}
				benchMarkStringResult = thisBenchMarkResult
			})
		})
	}

}
