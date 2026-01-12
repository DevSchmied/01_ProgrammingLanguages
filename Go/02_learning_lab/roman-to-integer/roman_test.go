package main

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"simple", "III", 3},
		{"mixed", "LVIII", 58},
		{"subtractive", "IV", 4},
		{"complex", "MCMXCIV", 1994},
		{"single", "M", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.input)
			if got != tt.want {
				t.Errorf("romanToInt(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestRomanToIntOptimized(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"simple", "III", 3},
		{"mixed", "LVIII", 58},
		{"subtractive", "IX", 9},
		{"complex", "MCMXCIV", 1994},
		{"max", "MMMCMXCIX", 3999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToIntOptimized(tt.input)
			if got != tt.want {
				t.Errorf("romanToIntOptimized(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
