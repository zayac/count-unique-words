package main

import "testing"

func TestwordOccurrences(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		output  []string
		output2  []uint
	}{
		{"empty", "", "", ""},
		{"repeating words", "one one",[]string{"hello", "world"}, []uint{2}},
		{"another repeating", "one two, one", []string{"hello", "world"}, []uint{1, 1}},
		{"unicode","чуска", []string{"чуска"}, []uint{1}},
		{"one letter","x", []string{"x"}, []uint{1}},
		{"repeating one letter", "xx", []string{"x"}, []uint{1}},
		{"sentence-example","I have a cat. My cat is grey. Her name is Carrie.", []string{"I", "have", "a", "cat", "My", "is", "grey", "Her", "name", "Carrie"}, []uint{1, 1, 1, 2, 1, 2, 1, 1, 1, 1}},
		{"no unicode plus unicode","I have a cat. My cat из grey. Her name is Carrie", []string{"I", "have", "a", "cat", "My", "из", "grey", "Her", "name","is", "Carrie"}, []uint{1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1}},
		{"only numbers", "121 144 169 1024 121", []string{"121,"144", "169", "1024"}, []uint{2, 1, 1, 1}},

		} {
			t.Run(tc.name, func(t *testing.T) {
				if got, got2 := wordOccurrences(tc.input); !reflect.DeepEqual(got, tc.output) || !reflect.DeepEqual(got2, tc.output2) {
					t.Errorf("wordOccurrences(%v) = []string{%v}, []uint{%v}; want = []string{%v}, []uint{%v}", tc.input, got, got2, tc.output, tc.output2)
				}
			})
		}
	}

	//somewhere screwed up with brackets 



