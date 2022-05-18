package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count the number of lines")
	fBytes := flag.Bool("b", false, "Count the number of bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *fBytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	wordCount := 0

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}
