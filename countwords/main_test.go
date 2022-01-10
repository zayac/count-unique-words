package main

import (
	"testing"
)

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func equaluints(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestWordOccurrences(t *testing.T) {
	for _, tc := range []struct {
		s           string
		words       []string
		occurrences []uint
	}{
		{"Hello", []string{"Hello"}, []uint{1}},
		{"Hello Hello Hello Hello", []string{"Hello"}, []uint{4}},
		{"плыли, плыли и приплыли", []string{"плыли", "и", "приплыли"}, []uint{2, 1, 1}},
		{"I am going to sleep", []string{"I", "am", "going", "to", "sleep"}, []uint{1, 1, 1, 1, 1}},
		{"You have come very fast. I have left very fast.", []string{"You", "have", "come", "very", "fast", "I", "left"}, []uint{1, 2, 1, 2, 2, 1, 1}},
		{"Gmail!google", []string{"Gmail", "google"}, []uint{1, 1}},
		{"", []string{}, []uint{}},
	} {
		got1, got2 := wordOccurrences(tc.s)
		if !equal(got1, tc.words) || !equaluints(got2, tc.occurrences) {
			t.Errorf("got1 = %q, want = %q, got2 = %v, want = %v", got1, tc.words, got2, tc.occurrences)
		}
	}
}
