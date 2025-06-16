package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	log.SetFlags(0)

	totals := Counts{}
	didError := false

	for i := 1; i < len(os.Args); i++ {

		filename := os.Args[i]
		counts, err := CountFile(filename)

		if err != nil {
			didError = true
			fmt.Fprintln(os.Stderr, "counter:", err)
			continue
		}

		fmt.Printf("%d %d %d %s\n", counts.Lines, counts.Words, counts.Bytes, filename)

		totals.Lines += counts.Lines
		totals.Words += counts.Words
		totals.Bytes += counts.Bytes

	}

	if len(os.Args) == 1 {
		counts := GetCounts(os.Stdin)
		fmt.Printf("%d %d %d\n", counts.Lines, counts.Words, counts.Bytes)
	}

	if len(os.Args) > 2 {
		fmt.Printf("%d %d %d total\n", totals.Lines, totals.Words, totals.Bytes)
	}

	fmt.Println("")

	if didError {
		os.Exit(1)
	}

}
