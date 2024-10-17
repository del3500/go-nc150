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
	if len(nums) <= 0 {
		return false
	}
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for _, v := range m {
		if v > 1 {
			return true
		}
	}
	return false
}
