package main

import (
	"reflect"
	"testing"
)

func TestwordOccurrences(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  []string
		want2 []uint
	}{
		{"empty", "", nil, nil},
		{"simple string", "hello world", []string{"hello", "world"}, []uint{1, 1}},
		{"string and repeat", "hello hello, world", []string{"hello", "world"}, []uint{2, 1}},
		{"string from README", "Where there is a will, there is a way.", []string{"Where", "there", "is", "a", "will", "way"}, []uint{1, 2, 2, 2, 1, 1}},
		{"string with unicode", "привет привет, мир"}, []string{"привет", "мир", []uint{2, 1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, got1 := wordOccurrences(tc.input); got != tc.want || tc.want2 {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

// test doesnt't work
