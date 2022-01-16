package main

import (
	"reflect"
	"testing"
)

func TestWordOccurrences(t *testing.T) {
	for _, tc := range []struct {
		name       string
		input      string
		wantstring []string
		wantuint   []uint
	}{
		{"empty", "", nil, nil},
		{"just a single letter", "a", []string{"a"}, []uint{1}},
		{"numbers", "111 2212 6901112764 311 2212", []string{"111", "2212", "6901112764", "311"}, []uint{1, 2, 1, 1}},
		{"example 1", "Where there is a will, there is a way.", []string{"Where", "there", "is", "a", "will", "way"}, []uint{1, 2, 2, 2, 1, 1}},
		{"example 2", "I have a cat. My cat is grey. Her name is Carrie.", []string{"I", "have", "a", "cat", "My", "is", "grey", "Her", "name", "Carrie"}, []uint{1, 1, 1, 2, 1, 2, 1, 1, 1, 1}},
		{"my own example", "I really love pizza, he loves it, we love it...", []string{"I", "really", "love", "pizza", "he", "loves", "it", "we"}, []uint{1, 1, 2, 1, 1, 1, 2, 1}},
		{"all unique words", "A journey of thousand steps begins with a single step.", []string{"A", "journey", "of", "thousand", "steps", "begins", "with", "a", "single", "step"}, []uint{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
		{"multiple occurrences of the same word", "Hello, friend! my friend!! friend!!! old friend,, friend??  friend... ", []string{"Hello", "friend", "my", "old"}, []uint{1, 6, 1, 1}},
		{"all the same words", "hello,hello,  hello!!!  hello. hello??  hello*++ ", []string{"hello"}, []uint{6}},
		{"a non-space separator", "hELLO!TherE?MY.old/fRIENd&Buddy;hELLO^", []string{"hELLO", "TherE", "MY", "old", "fRIENd", "Buddy"}, []uint{2, 1, 1, 1, 1, 1}},
		{"unicode example", "Много снега - много хлеба, много воды - много травы.", []string{"Много", "снега", "много", "хлеба", "воды", "травы"}, []uint{1, 1, 3, 1, 1, 1}},
		{"mixed unicode", " Leтом не pripaсешь, зиmoj  ne priнесешь! (Leтом horoшая поgoda кстаti ) ", []string{"Leтом", "не", "pripaсешь", "зиmoj", "ne", "priнесешь", "horoшая", "поgoda", "кстаti"}, []uint{2, 1, 1, 1, 1, 1, 1, 1, 1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, got1 := wordOccurrences(tc.input); !reflect.DeepEqual(got, tc.wantstring) || !reflect.DeepEqual(got1, tc.wantuint) {
				t.Errorf("got = %v, %v, want = %v, %v", got, got1, tc.wantstring, tc.wantuint)
			}
		})
	}
}
