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

	total := 0
	didError := false

	for i := 1; i < len(os.Args); i++ {

		filename := os.Args[i]
		wordCount, err := CountWordsInFile(filename)

		if err != nil {
			didError = true
			fmt.Fprintln(os.Stderr, "counter:", err)
			continue
		}

		fmt.Printf("%d %s\n", wordCount, filename)
		total += wordCount

	}

	if len(os.Args) == 1 {
		wordCount := CountWords(os.Stdin)
		fmt.Println(wordCount)
	}

	if len(os.Args) > 2 {
		fmt.Printf("%d total\n", total)
	}

	fmt.Println("")

	if didError {
		os.Exit(1)
	}

}

func CountWordsInFile(filename string) (int, error) {

	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return CountWords(file), nil

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
