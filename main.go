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
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)
	wordCount := 0

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}
