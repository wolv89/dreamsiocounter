package main

import (
	"fmt"
	"os"
)

func main() {

	data, _ := os.ReadFile("./words.txt")

	fmt.Println(CountWords(data))

}

func CountWords(data []byte) int {

	if len(data) == 0 {
		return 0
	}

	wordCount := 1

	for _, b := range data {
		if b == ' ' {
			wordCount++
		}
	}

	return wordCount

}
