package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func IsSep(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r)
}

func Words(s string) []string {
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
	return 0
}

func wordOccurrences(s string) (words []string, occurrences []uint) {
	w := Words(s)
	for _, word := range w {
		if len(words) == len(occurrences) { // specified in the task
			count := find(word, words)
			if count == 0 {
				words = append(words, word)
				occurrences = append(occurrences, 1)
			} else {
				occurrences[count]++
			}
		}
	}
	return words, occurrences
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println("Occurrences:")
	words, occurrences := wordOccurrences(s)
	for i, r := range words {
		fmt.Printf("%v: %v\n", r, occurrences[i])
	}
}
