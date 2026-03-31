package main

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// Arrange
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{"reverses order", "cba", "abc"},
		{"already sorted", "abc", "abc"},
		{"all same chars", "aaaaa", "aaaaa"},
		{"empty string", "", ""},
		{"single char", "a", "a"},
		{"with duplicates", "bcaacb", "aabbcc"},
		{"long string", "zyxwvutsrqponmlkjihgfedcba", "abcdefghijklmnopqrstuvwxyz"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			res := bubbleSort(tc.input)
			// Assert
			if res != tc.want {
				t.Errorf("bubbleSort(%q) = %q, want %q", tc.input, res, tc.want)
			}
		})
	}
}
