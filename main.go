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
	ch, errCh := CountFiles(filenames)

loop:
	for {
		select {
		case res, open := <-ch:
			if !open {
				break loop
			}
			totals.Add(res.counts)
			res.counts.Print(opts, res.filename)
		case err, open := <-errCh:
			if !open {
				break loop
			}
			didError = true
			fmt.Fprintln(os.Stderr, "counter:", err)
		}
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

func CountFiles(filenames []string) (<-chan CountsWithFilename, <-chan error) {

	ch := make(chan CountsWithFilename)
	errCh := make(chan error)

	wg := sync.WaitGroup{}
	wg.Add(len(filenames))

	for _, filename := range filenames {

		go func() {

			defer wg.Done()

			counts, err := CountFile(filename)

			if err != nil {
				errCh <- err
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

	return ch, errCh

}
