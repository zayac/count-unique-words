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
	for s, w := range words {
		if w == word {
			return s
		}
	}
	return -1
}
func wordOccurrences(s string) (stringwords []string, occurrences []uint) {
	w := words(s)
	for _, word := range w {
		if find(word, stringwords) == -1 {
			stringwords = append(stringwords, word)
			occurrences = append(occurrences, 1)
		} else {
			occurrences[find(word, stringwords)] += 1
		}
	}
	return stringwords, occurrences
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(wordOccurrences(s))
}
