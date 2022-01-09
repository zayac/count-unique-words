package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func IsSep(r rune) bool { return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) }

func wordOccurrences(s string) (words []string, occurrences []uint) {
	var w []string
	var tmp []rune
	wasSep := true
	for _, r := range s {
		if IsSep(r) {
			if !wasSep {
				w = append(w, string(tmp))
				tmp = nil
			}
			wasSep = true
		} else {
			wasSep = false
			tmp = append(tmp, r)
		}
	}
	for i, j := range w {
		if j == s {
			occurrences[i]++
		} else {
			words = append(w, string(tmp))

		}
	}
	return words, occurrences
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(wordOccurrences(s))
}
