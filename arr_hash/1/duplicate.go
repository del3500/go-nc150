package main

import "fmt"

// Given an integer array nums, return true if any value appears more than once in the array,
// otherwise, return false.
//
// Example 1:
// Input: nums := []int{1,2,3,3}
// Output: true
//
// Example 2:
// Input: nums := []int{1,2,3,4}
// Output: false

func main() {
	nums := []int{1, 1, 3}
	fmt.Println(HasDuplicate(nums))
}

func HasDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, exists := m[v]; exists {
			return true
		}
		m[v] = struct{}{}
	}
	return false
}
