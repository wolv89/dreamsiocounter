package main

import (
	"bufio"
	"io"
	"os"
)

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

func CountLines(file io.Reader) int {

	lineCount := 0

	reader := bufio.NewReader(file)

	for {

		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		if r == '\n' {
			lineCount++
		}

	}

	return lineCount

}
