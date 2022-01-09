package main

import (
	"fmt"
	"unicode"
)

func IsSep(r rune) bool { return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) }

func CountOccurrences(s string) (words []string, occurrences []uint) {
	var tmp []rune
	wasSep := true
	for _, r := range s {
		if IsSep(r) {
			if !wasSep {
				s = append(s, string(tmp))
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
	return
}

func main() {
	fmt.Println("Enter string")
	var s uint
	fmt.Scan(&s)
	fmt.Println(CountOccurrences(occurrences))

}
