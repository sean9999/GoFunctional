package fslice_test

import (
	"strings"
	"testing"

	"github.com/sean9999/GoFunctional/fslice"
)

var scrabbleValueMap map[string]int = map[string]int{
	"a": 1,
	"b": 3,
	"c": 3,
	"d": 2,
	"e": 1,
	"f": 4,
	"g": 2,
	"h": 4,
	"i": 1,
	"j": 8,
	"k": 5,
	"l": 1,
	"m": 3,
	"n": 1,
	"o": 1,
	"p": 3,
	"q": 10,
	"r": 1,
	"s": 1,
	"t": 1,
	"u": 1,
	"v": 4,
	"w": 4,
	"x": 8,
	"y": 4,
	"z": 10,
}

func scrabbleValue(word string) int {
	score := 0
	for _, letter := range strings.Split(strings.ToLower(word), "") {
		score += scrabbleValueMap[letter]
	}
	return score
}

func TestReduce(t *testing.T) {

	t.Run("sum integers", func(t *testing.T) {
		sum := func(a, b int32) int32 {
			return a + b
		}
		inputSlize := []int32{1000, 100, 10, 1}
		want := int32(1111)
		got := fslice.From(inputSlize).Reduce(sum, 0)
		assertScalars(t, got, want)
	})

	t.Run("highest scrabble-value word", func(t *testing.T) {
		highestScrabbleValue := func(a, b string) string {
			score_a := scrabbleValue(a)
			score_b := scrabbleValue(b)
			if score_a > score_b {
				return a
			} else {
				return b
			}
		}
		inputSlize := []string{"all", "your", "base", "are", "belonging", "to", "us"}
		want := "belonging"
		got := fslice.From(inputSlize).Reduce(highestScrabbleValue, "")
		assertScalars(t, got, want)
	})

}
