package main

import (
	"fmt"
	"os"
)

func main() {

	data, _ := os.ReadFile("./words.txt")
	_ = data

	wordCount := 1

	for _, b := range data {
		if b == ' ' {
			wordCount++
		}
	}

	fmt.Println(wordCount)

}
