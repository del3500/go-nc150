package main

import "fmt"

// Problem:
// given 2 strings s and t, return true if two strings are anagrams of each other, otherwise return false.

func main() {
	s1 := "abcdefg"
	s2 := "abcdefg"
	fmt.Println(Anagram(s1, s2))
}

func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[string]int)
	for _, v := range s1 {
		m[string(v)]++
	}
	for _, v := range s2 {
		m[string(v)]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
