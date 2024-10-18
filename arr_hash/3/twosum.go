package main

import "fmt"

func main() {
	nums1 := []int{4, 2111, 5, 8, 0, 12382183}
	target := 7
	fmt.Println(twoSum(nums1, target))
}

func twoSum(nums []int, t int) []int {
	ans := make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == t {
				ans = append(ans, i)
				ans = append(ans, j)
				return ans
			}
		}
	}
	return ans
}
