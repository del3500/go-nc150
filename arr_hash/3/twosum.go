package main

import "fmt"

func main() {
	nums1 := []int{4, 2111, 5, 8, 0, 12382183}
	target := 7
	fmt.Println(twoSum(nums1, target))
}

func twoSum(nums []int, t int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := t - num
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}
