package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func IsSep(r rune) bool { return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) }

func words(s string) []string {
	var w []string
	var tmp []rune
	wasSep := true
	for _, r := range s {
		if IsSep(r) {
			if !wasSep {
				w = append(w, string(tmp))
				// tmp = tmp[:0]
				tmp = nil
			}
			wasSep = true
		} else {
			wasSep = false
			tmp = append(tmp, r)
		}
	}
	if len(tmp) > 0 {
		w = append(w, string(tmp))
	}
	return w
}
func find(word string, words []string) int {
	for i, w := range words {
		if w == word {
			return i
		}
	}
	return +1
}

func wordOccurrences(s string) (words2 []string, occurrences []uint) {
	w := words(s)
	//code from uniqword
	for _, word := range w {
		// code need to use func find
	}
	return words2, occurrences
}

func main() {
	fmt.Print("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	words2, occurrences := wordOccurrences(s)
	for i := range words2 {
		fmt.Println(words2[i], "occurences: ")
		fmt.Println(occurrences[i])
	}
}
