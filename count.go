package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Counts struct {
	Lines, Words, Bytes int
}

func (c *Counts) Add(d Counts) {
	c.Lines += d.Lines
	c.Words += d.Words
	c.Bytes += d.Bytes
}

func (c *Counts) Print(filenames ...string) {

	fmt.Printf("%d %d %d", c.Lines, c.Words, c.Bytes)

	for _, filename := range filenames {
		fmt.Printf(" %s", filename)
	}

	fmt.Print("\n")

}

func GetCounts(f io.ReadSeeker) Counts {

	counts := Counts{}

	counts.Lines = CountLines(f)
	f.Seek(0, io.SeekStart)
	counts.Words = CountWords(f)
	f.Seek(0, io.SeekStart)
	counts.Bytes = CountBytes(f)

	return counts
}

func CountFile(filename string) (Counts, error) {

	file, err := os.Open(filename)
	if err != nil {
		return Counts{}, err
	}
	defer file.Close()

	return GetCounts(file), nil

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
