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

var TestSuite = functionalTestSuite{
	LoremIpsumFilePath: "fslice/testdata/lorem_ipsum_%d_words.txt",
	LoremIpsumLengths:  []int{10, 100, 1_000, 10_000, 100_000},
}

func (fts functionalTestSuite) LoadLoremIpsum(numWords int) (string, error) {
	var returnString string
	fileBytes, err := os.ReadFile(fmt.Sprintf(TestSuite.LoremIpsumFilePath, numWords))
	if err == nil {
		returnString = string(fileBytes)
	}
	return returnString, err
}
