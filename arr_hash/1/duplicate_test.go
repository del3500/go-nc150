package main

import "testing"

func TestHasDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "no duplicates",
			nums:     []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "has duplicates",
			nums:     []int{1, 2, 1},
			expected: true},
		{
			name:     "no element",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HasDuplicate(test.nums)
			if result != test.expected {
				t.Errorf("For %v, expected %v, but got %v", test.nums, test.expected, result)
			}
		})
	}
}
