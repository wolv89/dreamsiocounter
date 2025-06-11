package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {

	filename := "./words.txt"

	// log.SetFlags(0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	wordCount := CountWordsInFile(file)
	fmt.Println(wordCount)

}

func CountWordsInFile(file *os.File) int {

	wordCount := 0
	isInsideWord := false

	reader := bufio.NewReader(file)

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		if !unicode.IsSpace(r) && !isInsideWord {
			wordCount++
		}

		isInsideWord = !unicode.IsSpace(r)
	}

	return wordCount

}

func CountWords(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}
