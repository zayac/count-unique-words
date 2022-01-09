package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordOccurrences(s string) (words []string, occurrences []uint) {

	return words, occurrences
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(wordOccurrences(s))
}
