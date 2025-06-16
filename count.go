package main

import (
	"bufio"
	"io"
	"os"
)

type Counts struct {
	Lines, Words, Bytes int
}

func CountWordsInFile(filename string) (Counts, error) {

	file, err := os.Open(filename)
	if err != nil {
		return Counts{}, err
	}
	defer file.Close()

	counts := Counts{}

	counts.Lines = CountLines(file)
	counts.Words = CountWords(file)
	counts.Bytes = CountBytes(file)

	return counts, nil

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

func CountBytes(file io.Reader) int {

	byteCount, _ := io.Copy(io.Discard, file)
	return int(byteCount)

}
