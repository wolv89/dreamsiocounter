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

type CountsWithFilename struct {
	counts   Counts
	filename string
}

func (c *Counts) Add(d Counts) {
	c.Lines += d.Lines
	c.Words += d.Words
	c.Bytes += d.Bytes
}

func (c *Counts) Print(opts DisplayOptions, filenames ...string) {

	stats := []string{}

	if opts.ShouldShowLines() {
		stats = append(stats, strconv.Itoa(c.Lines))
	}
	if opts.ShouldShowWords() {
		stats = append(stats, strconv.Itoa(c.Words))
	}
	if opts.ShouldShowBytes() {
		stats = append(stats, strconv.Itoa(c.Bytes))
	}

	fmt.Fprint(opts.Target, strings.Join(stats, "\t")+"\t")

	for _, fname := range filenames {
		fmt.Fprint(opts.Target, " "+fname)
	}

	fmt.Fprint(opts.Target, "\n")

}

func PrintHeader(opts DisplayOptions) {

	if !opts.ShowHeader {
		return
	}

	cols := []string{}

	if opts.ShouldShowLines() {
		cols = append(cols, "lines")
	}
	if opts.ShouldShowWords() {
		cols = append(cols, "words")
	}
	if opts.ShouldShowBytes() {
		cols = append(cols, "bytes")
	}

	fmt.Fprintln(opts.Target, strings.Join(cols, "\t")+"\t")

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
