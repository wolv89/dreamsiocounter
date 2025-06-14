package main

import (
	"fmt"
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
