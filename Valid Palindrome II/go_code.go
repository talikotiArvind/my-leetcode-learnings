package main

import (
	"fmt"
)

// validPalindrome determines if a string can be a palindrome after deleting at most one character.
func validPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			// If a mismatch occurs, check both deletion options:
			// 1. Skip the left character: check range [left+1, right]
			// 2. Skip the right character: check range [left, right-1]
			return isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
		}
		left++
		right--
	}

	return true
}

func isPalindrome(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	input1 := "aba"
	fmt.Printf("Input: %q -> Valid Palindrome II: %v\n", input1, validPalindrome(input1))

	input2 := "abca"
	fmt.Printf("Input: %q -> Valid Palindrome II: %v\n", input2, validPalindrome(input2))

	input3 := "abcde"
	fmt.Printf("Input: %q -> Valid Palindrome II: %v\n", input3, validPalindrome(input3))
}
