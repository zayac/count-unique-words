package main

import (
	"fmt"
	"strings"
)
func wordOccurrences(s string) (words []string, occurrences []uint){
	o := make(occurrences []uint)
	for x, y := range s {
		fmt.Println(x,y)
		o[y]=x
	}
	return o
}
func main() {
	s := []string{"I have a cat."," My cat is grey."," Her name is Carrie."}
	fmt.Println("s:", s)
	o := wordOccurrences(s)
	fmt.Println(o)
}