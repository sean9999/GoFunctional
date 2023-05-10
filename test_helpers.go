package functional

import (
	"fmt"
	"os"
)

//	assistance for testing and benchmarks
//	use ./bin/generate-test-data to generate test data

type functionalTestSuite struct {
	LoremIpsumFilePath string
	LoremIpsumLengths  []int
}

// TestSuite provides convenience functions for tests and benchmarks
// see: ./bin/generate-test-data/ to generate test data
var TestSuite = functionalTestSuite{
	LoremIpsumFilePath: "fslice/testdata/lorem_ipsum_%d_words.txt",
	LoremIpsumLengths:  []int{10, 100, 1_000, 10_000, 100_000},
}

// LoadLoremIpsum loads lorem ipsum from fixtures.
// see: [bin/generate-test-data/main.go]
func (fts functionalTestSuite) LoadLoremIpsum(numWords int) (string, error) {
	var returnString string
	fileBytes, err := os.ReadFile(fmt.Sprintf(TestSuite.LoremIpsumFilePath, numWords))
	if err == nil {
		returnString = string(fileBytes)
	}
	return returnString, err
}
