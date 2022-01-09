package main

import (
	"reflect"
	"testing"
)

func TestWordOccurrences(t *testing.T) {
	for _, tc := range []struct {
		s               string
		wantWords       []string
		wantOccurrences []uint
	}{
		{"", nil, nil},
		{"I have a cat. My cat is grey. Her name is Carrie.", []string{"I", "have", "a", "cat", "My", "is", "grey", "Her", "name", "Carrie"}, []uint{1, 1, 1, 2, 1, 2, 1, 1, 1, 1}},
		{"Where there is a will, there is a way.", []string{"Where", "there", "is", "a", "will", "way"}, []uint{1, 2, 2, 2, 1, 1}},
		{"Много снега - много хлеба, много воды - много травы.", []string{"Много", "снега", "много", "хлеба", "воды", "травы"}, []uint{1, 1, 3, 1, 1, 1}},
		{"Example with with duplication.", []string{"Example", "with", "duplication"}, []uint{1, 2, 1}},
		{"Пример с с с повторением, повторением.", []string{"Пример", "с", "повторением"}, []uint{1, 3, 2}},
	} {
		got1, got2 := wordOccurrences(tc.s)
		if !reflect.DeepEqual(got1, tc.wantWords) || !reflect.DeepEqual(got2, tc.wantOccurrences) {
			t.Errorf("got1 = %q, want = %q, got2 = %v, want = %v", got1, tc.wantWords, got2, tc.wantOccurrences)
		}
	}
}
