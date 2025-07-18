package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type DisplayOptions struct {
	ShowLines, ShowWords, ShowBytes bool
}

func (opts DisplayOptions) ShouldShowLines() bool {
	if !opts.ShowBytes && !opts.ShowLines && !opts.ShowWords {
		return true
	}
	return opts.ShowLines
}

func (opts DisplayOptions) ShouldShowWords() bool {
	if !opts.ShowBytes && !opts.ShowLines && !opts.ShowWords {
		return true
	}
	return opts.ShowWords
}

func (opts DisplayOptions) ShouldShowBytes() bool {
	if !opts.ShowBytes && !opts.ShowLines && !opts.ShowWords {
		return true
	}
	return opts.ShowBytes
}

func main() {

	opts := DisplayOptions{}

	flag.BoolVar(&opts.ShowLines, "l", false, "Used to toggle whether or not to show the line count")
	flag.BoolVar(&opts.ShowWords, "w", false, "Used to toggle whether or not to show the word count")
	flag.BoolVar(&opts.ShowBytes, "c", false, "Used to toggle whether or not to show the byte count")

	flag.Parse()

	log.SetFlags(0)

	totals := Counts{}
	didError := false

	filenames := flag.Args()

	for _, filename := range filenames {

		counts, err := CountFile(filename)

		if err != nil {
			didError = true
			fmt.Fprintln(os.Stderr, "counter:", err)
			continue
		}

		counts.Print(opts, filename)
		totals.Add(counts)

	}

	if len(filenames) == 0 {
		counts := GetCounts(os.Stdin)
		counts.Print(opts)
	} else if len(filenames) >= 2 {
		totals.Print(opts, "total")
	}

	fmt.Println("")

	if didError {
		os.Exit(1)
	}

}
