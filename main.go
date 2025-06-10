package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	filename := "./words.txt"

	// log.SetFlags(0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	PrintFileContents(file)
	// fmt.Println(CountWords(data))

}

func PrintFileContents(file *os.File) {

	const bufferSize = 8192
	buffer := make([]byte, bufferSize)

	for {
		size, err := file.Read(buffer)
		if err != nil {
			break
		}

		subBuffer := buffer[:size]
		fmt.Println(string(subBuffer))
	}

}

func CountWords(data []byte) int {
	words := bytes.Fields(data)
	return len(words)
}
