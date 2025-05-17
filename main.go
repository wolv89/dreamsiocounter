package main

import (
	"fmt"
	"os"
)

func main() {

	data, _ := os.ReadFile("./words.txt")

	fmt.Println(countWords(data))

}

func countWords(data []byte) int {

	wordCount := 1

	for _, b := range data {
		if b == ' ' {
			wordCount++
		}
	}

	return wordCount

}
