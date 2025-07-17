package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Counts struct {
	Lines, Words, Bytes int
}

func (c *Counts) Add(d Counts) {
	c.Lines += d.Lines
	c.Words += d.Words
	c.Bytes += d.Bytes
}

func (c *Counts) Print(opts DisplayOptions, filenames ...string) {

	xs := []string{}

	if opts.ShouldShowLines() {
		xs = append(xs, strconv.Itoa(c.Lines))
	}
	if opts.ShouldShowWords() {
		xs = append(xs, strconv.Itoa(c.Words))
	}
	if opts.ShouldShowBytes() {
		xs = append(xs, strconv.Itoa(c.Bytes))
	}

	xs = append(xs, filenames...)

	fmt.Print(strings.Join(xs, " "), "\n")

}

func GetCounts(f io.Reader) Counts {

	counts := Counts{}

	inWord := false
	reader := bufio.NewReader(f)

	for {

		r, size, err := reader.ReadRune()
		if err != nil {
			break
		}

		counts.Bytes += size

		if r == '\n' {
			counts.Lines++
		}

		isSpace := unicode.IsSpace(r)
		if !isSpace && !inWord {
			counts.Words++
		}

		inWord = !isSpace

	}

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
