package fslice_test

import (
	"fmt"
	"strings"

	"github.com/fxtlabs/primes"
	"github.com/sean9999/GoFunctional/fslice"
)

func Example() {

	fib := generateFibonacci(25)

	// of the first 25 fibonacci numbers, which are not primes?
	fibNonPrimes := fslice.From(fib).Filter(func(n int, _ int, _ []int) bool {
		r := false
		if n > 1 {
			r = !primes.IsPrime(n)
		}
		return r
	})

	// among those, which are co-prime with all others?
	coPrimes := fibNonPrimes.Filter(func(n int, i int, nonPrimes []int) bool {

		return fslice.From(nonPrimes).Every(func(m int, j int, _ []int) bool {
			return (i == j) || primes.Coprime(n, m)
		})

	})

	// Fslice is generic, accepting any comparable type
	lowerBase := []string{"all", "your", "base", "are", "belong", "to", "us"}
	shoutCaseEveryOther := func(word string, i int, _ []string) string {
		if i%2 == 1 {
			word = strings.ToUpper(word)
		}
		return word
	}
	shoutyBase := fslice.From(lowerBase).Map(shoutCaseEveryOther)

	fmt.Println(fibNonPrimes)
	fmt.Println(coPrimes)
	fmt.Println(shoutyBase)

	// Output:
	// [8 21 34 55 144 377 610 987 2584 4181 6765 10946 17711 46368]
	// [4181 17711]
	// [all YOUR base ARE belong TO us]

}
