package main

import (
	"reflect"
	"testing"
)

func TestTrimLeft(t *testing.T) {
	for _, tc := range []struct {
		name    string
		input   string
		output  []string
		output1 []uint
	}{
		{"empty", "", nil, nil},
		{"No repeat", "Hello, world", []string{"Hello", "world"}, []uint{1, 1}},
		{"Siple repeat", "Hi Hi", []string{"Hi"}, []uint{2}},
		{"Repeat with words between", "Hi Hello Hi", []string{"Hi", "Hello"}, []uint{2, 1}},
		{"Unicode characters instead of spaces", "\n\r\tHello\n\r\tHi\n\t\rHello\n\r\t", []string{"Hello", "Hi"}, []uint{2, 1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, got1 := wordOccurrences(tc.input); !reflect.DeepEqual(got, tc.output) || !reflect.DeepEqual(got1, tc.output1) {
				t.Errorf("wordOccurrences(%v) = []string{%v}, []uint{%v}; want = []string{%v}, []uint{%v}", tc.input, got, got1, tc.output, tc.output1)
			}
		})
	}
}
