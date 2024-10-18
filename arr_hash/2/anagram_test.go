package main

import "testing"

func TestAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "valid anagram",
			s1:       "abcdefg",
			s2:       "abcdefg",
			expected: true,
		},
		{
			name:     "invalid anagram",
			s1:       "abcdd",
			s2:       "abcde",
			expected: false,
		},
		{
			name:     "empty string",
			s1:       "",
			s2:       "a",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Anagram(test.s1, test.s2)
			want := test.expected
			if got != want {
				t.Errorf("given: %s and %s, want %v, got %v", test.s1, test.s2, got, want)
			}
		})
	}
}
