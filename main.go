package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	filename := "./words.txt"

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	fmt.Println(CountWords(data))

}

func CountWords(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}
