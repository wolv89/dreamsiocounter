package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"text/tabwriter"
)

type DisplayOptions struct {
	Target                                      io.Writer
	ShowLines, ShowWords, ShowBytes, ShowHeader bool
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
	flag.BoolVar(&opts.ShowHeader, "header", false, "Used to toggle whether or not to show the column names as a header line")

	flag.Parse()

	log.SetFlags(0)

	wr := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)
	opts.Target = wr

	PrintHeader(opts)

	totals := Counts{}
	didError := false

	filenames := flag.Args()
	wg := sync.WaitGroup{}
	wg.Add(len(filenames))

	ch := make(chan CountsWithFilename)

	for _, filename := range filenames {

		go func() {

			defer wg.Done()

			counts, err := CountFile(filename)

			if err != nil {
				didError = true
				fmt.Fprintln(os.Stderr, "counter:", err)
				return
			}

			ch <- CountsWithFilename{
				counts,
				filename,
			}

		}()

	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		totals.Add(res.counts)
		res.counts.Print(opts, res.filename)
	}

	if len(filenames) == 0 {
		counts := GetCounts(os.Stdin)
		counts.Print(opts)
	} else if len(filenames) >= 2 {
		totals.Print(opts, "total")
	}

	fmt.Fprintln(opts.Target, "")

	wr.Flush()

	if didError {
		os.Exit(1)
	}

}
