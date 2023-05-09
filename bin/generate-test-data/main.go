package main

import (
	"fmt"
	"os"
	"strings"

	lorem "github.com/drhodes/golorem"
)

const loremIpsumFilePath = "fslice/testdata/lorem_ipsum_%d_words.txt"

var loremIpsumLengths = []int{10, 100, 1_000, 10_000, 100_000}

func loremIpsumSlice(numwords int) []string {
	text := lorem.Sentence(numwords, numwords)
	return strings.Split(text, " ")
}

func generateLoremIpsum() {
	for _, thisLength := range loremIpsumLengths {
		thisFileName := fmt.Sprintf(loremIpsumFilePath, thisLength)
		thisHandle, err := os.Create(thisFileName)
		if err == nil {
			inputStrings := loremIpsumSlice(thisLength)
			for i, word := range inputStrings {
				if i == 0 {
					thisHandle.WriteString(word)
				} else {
					thisHandle.WriteString(fmt.Sprintf(", %s", word))
				}
			}
			err = thisHandle.Close()
			if err == nil {
				fmt.Printf("wrote %d\twords to %s\n", thisLength, thisFileName)
			} else {
				fmt.Printf("could not release handle for %s:\t%v\n", thisFileName, err)
			}
		} else {
			fmt.Printf("could not get handle for %s:\t%v\n", thisFileName, err)
		}
	}
}

func main() {
	generateLoremIpsum()
}
