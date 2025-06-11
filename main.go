package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
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

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCount++
	}

	return wordCount

}

func CountWords(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}
