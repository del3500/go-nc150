package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "test1",
			nums:     []int{1, 2, 3, 4, 5},
			target:   9,
			expected: []int{3, 4},
		},
		{
			name:     "test2",
			nums:     []int{5, 5, 2},
			target:   50,
			expected: nil,
		},
		{
			name:     "test3",
			nums:     []int{5, 5, 2},
			target:   10,
			expected: []int{0, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := twoSum(test.nums, test.target)
			want := test.expected
			if !reflect.DeepEqual(got, want) {
				t.Errorf("given: %v, got %v, want %v", test.nums, got, want)
			}
		})
	}
}
