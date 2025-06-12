package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	log.SetFlags(0)

	if len(os.Args) < 2 {
		log.Fatalln("error: no filename provided")
	}

	total := 0

	for i := 1; i < len(os.Args); i++ {

		filename := os.Args[i]
		wordCount := CountWordsInFile(filename)

		fmt.Printf("%d %s\n", wordCount, filename)
		total += wordCount

	}

	if len(os.Args) > 2 {
		fmt.Printf("%d total\n", total)
	}

	fmt.Println("")

}

func CountWordsInFile(filename string) int {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	return CountWords(file)

}

func CountWords(file io.Reader) int {

	wordCount := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCount++
	}

	return wordCount

}
