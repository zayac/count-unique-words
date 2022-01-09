package main

import (
	"fmt"
	"unicode"
)

func IsSep(r rune) bool { return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) }

func wordOccurrences(s string) (words []string, occurrences []uint) {
	var tmp []rune
	wasSep := true
	for _, r := range s {
		if IsSep(r) {
			if !wasSep {
				s = append(s, string(tmp))
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
		s = append(s, string(tmp))
	}
	//надо создать срез с отформатированными значениями
	return
}

func main() {
	var s uint
	fmt.Println("Enter a string: ")
	fmt.Scan(&s)
	fmt.Println("Occurrences: ", wordOccurrences(occurrences))
	return
}
