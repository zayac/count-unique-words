package main

import "testing"

func TestwordOccurrences(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		output  []string
		outputt  []uint
	}{
		{"empty", "", "", ""},
		{"one letter", "a", "a", "1"},
		{"repeating letters", "bebe", []string{"bebe"}, []uint{1}},
		{"sentence", "I have a cat. My cat is grey. Her name is Carrie.", []string{"I", "have", "a", "cat", "My", "is", "grey", "Her", "name", "Carrie"}, []uint{1, 1, 1, 2, 1, 2, 1, 1, 1, 1}},
		{"two same words", "kot kot", []string{"kot"}, []uint{2}},
		{"special letters", "čelovek ne čelovek", []string{"čelovek", "ne"}, uint[]{2, 1}},
		{"numbers", "123 123 000009", []string{"123", "09"}, []uint{2, 1}},
		{"separators", "jelka\n\r\sneg", []string{"jelka", "sneg"}, []uint{1, 1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, got2 := wordOccurrences(tc.input); !reflect.DeepEqual(got, tc.output) || !reflect.DeepEqual(got2, tc.output2) {
				t.Errorf("wordOccurrences(%v) = []string{%v}, []uint{%v}; want = []string{%v}, []uint{%v}", tc.input, got, got2, tc.output, tc.output2)
			}
		})
	}
}
