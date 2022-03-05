package main

import (
	"reflect"
	"testing"
)

func TestWordOccurrences(t *testing.T) {
	for _, tc := range []struct {
		s           string
		words       []string
		occurrences []uint
	}{
		{"", []string{}, []uint{}},
		{"Hello", []string{"Hello"}, []uint{1}},
		{"Hello*world", []string{"Hello", "world"}, []uint{1, 1}},
		{"no no no no", []string{"no"}, []uint{4}},
		{"I am going to sleep", []string{"I", "am", "going", "to", "sleep"}, []uint{1, 1, 1, 1, 1}},
		{"Много снега - много хлеба, много воды - много травы.", []string{"Много", "снега", "много", "хлеба", "воды", "травы"}, []uint{1, 1, 3, 1, 1, 1}},
		{"I have a cat. My cat is grey. Her name is Carrie", []string{"I", "have", "a", "cat", "My", "is", "grey", "Her", "name", "Carrie"}, []uint{1, 1, 1, 2, 1, 2, 1, 1, 1, 1}},
	} {
		got1, got2 := wordOccurrences(tc.s)
		if !reflect.DeepEqual(got1, tc.words) || !reflect.DeepEqual(got2, tc.occurrences) {
			t.Errorf("got1 = %q, want = %q, got2 = %v, want = %v", got1, tc.words, got2, tc.occurrences)
		}
	}
}
