package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func find(word string, words []string) int {
	for i, w := range words {
		if w == word {
			return i
		}
	}
	return -1
}

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

func wordOccurrences(s string) (string_of_words []string, occurrences []uint) {
	w := words(s)
	for _, word := range w {
		index := find(word, string_of_words)
		if index == -1 {
			string_of_words = append(string_of_words, word)
			occurrences = append(occurrences, 1)
		} else {
			occurrences[index]++
		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	string_of_words, occurrences := wordOccurrences(text)
	for k, v := range string_of_words {
		fmt.Printf("%v: %v\n", v, occurrences[k])
	}
}
