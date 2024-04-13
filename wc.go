package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func countWords(scanner *bufio.Scanner) int {
	scanner.Split(bufio.ScanWords)
	fmt.Println("hi")

	wc := 0
	for scanner.Scan() {
		wc += 1
	}
	return wc
}

func countLines(scanner *bufio.Scanner) int {
	scanner.Split(bufio.ScanLines)
	wc := 0
	for scanner.Scan() {
		wc += 1
	}
	return wc
}

func main() {
	shouldCountLines := flag.Bool("l", true, "count lines")
	shouldCountWords := flag.Bool("w", true, "count lines")

	flag.Parse()

	output := ""

	scanner := bufio.NewScanner(os.Stdin)

	if *shouldCountLines {
		output += fmt.Sprintf("\t%d", countLines(scanner))
	}
	if *shouldCountWords {
		output += fmt.Sprintf("\t%d", countWords(scanner))
	}
	fmt.Println(output)
}
