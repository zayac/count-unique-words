package main

import (
	"reflect"
	"testing"
)

func TestwordOccurrences(t *testing.T) {
	for _, tc := range []struct {
		name    string
		input   string
		output1 []string
		output2 []uint
	}{
		{"empty", "", nil, nil},
		{"repeat", "hi, hi", []string{"hi"}, []uint{2}},
		{"no repeat", "hi, qq", []string{"hi", "qq"}, []uint{1, 1}},
		{"repeat and no repeat", "hi hi, qq", []string{"hi", "q"}, []uint{2, 1}},
		{"sentence from example1", "Where there is a will, there is a way.", []string{"Where", "there", "is", "a", "will", "way"}, []uint{1, 2, 2, 2, 1, 1}},
		{"sentence from example2", "Много снега - много хлеба, много воды - много травы.", []string{"Много", "снега", "много", "хлеба", "воды", "травы"}, []uint{1, 1, 3, 1, 1, 1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got1, got2 := wordOccurrences(tc.input); !reflect.DeepEqual(got1, tc.output1) || !reflect.DeepEqual(got2, tc.output2) {
				t.Errorf("wordOccurrences(%v) = []string{%v}, []uint{%v}; want = []string{%v}, []uint{%v}", tc.input, go1, got2, tc.output1, tc.output2)
			}
		})
	}
}
