package main

import (
	"testing"
)

func TestwordOccurunces(t *testing.T) {
	for _, tc := range []struct {
		name    string
		input   string
		output1 []string
		output2 []uint
	}{
		{"empty string", "", nil, nil},
		{"string without repeated words", "He have scored a goal", []string{"He", "have", "scored", "a", "goal"}, []uint{1, 1}},
		{"string with repeated words", "I was interrupted so I needed to restart", []string{"I", "was", "interrupted", "so", "needed", "to", "restart"}, []uint{2, 1, 1, 1, 1, 1, 1}},
		{"string with repeated words 2.0", "He was angry so he left", []string{"He", "was", "angry", "so", "he", "left"}, []uint{1, 1, 1, 1, 1, 1}},
		{"string in unicode", "Плыли мы по морю", []string{"Плыли", "мы", "по", "морю"}, []uint{1, 1, 1, 1}},
		{"string where words are looking the same, but their letters are written in different languages", "English Water and Russian Wаter", []string{"English", "Water", "and", "Russian", "Wаter"}, []uint{1, 1, 1, 1, 1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, got1 := wordOccurrences(tc.input); got != tc.want || tc.want2 {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
