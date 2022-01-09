package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func IsSep(r rune) bool { return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) }

func Words(s string) []string {
	var w []string
	var tmp []rune
	wasSep := true
	for _, r := range s {
		if IsSep(r) {
			if !wasSep {
				w = append(w, string(tmp))
				tmp = nil
				wasSep = true
			}
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
	return -1
}

func wordOccurrences(s string) (words []string, occurrences []uint) {
	w := Words(s)
	for _, word := range w {
		if find(word, words) == -1 {
			words = append(words, word)
			occurrences = append(occurrences, 1)
		} else {
			i := find(word, words)
			occurrences[i]++
		}
	}
	return words, occurrences
}

func main() {
	fmt.Print("Enter a string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	words, occurrences := wordOccurrences(s)
	for i, word := range words {
		fmt.Println(word, ":", occurrences[i])
	}
}
