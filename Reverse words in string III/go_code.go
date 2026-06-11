package main

import (
	"fmt"
)
func reverseWords(s string) string {
	bytes := []byte(s)
	left := 0

	for right := 0; right <= len(bytes); right++ {
		if right == len(bytes) || bytes[right] == ' ' {
			start := left
			end := right - 1
			for start < end {
				bytes[start], bytes[end] = bytes[end], bytes[start]
				start++
				end--
			}
			left = right + 1
		}
	}

	return string(bytes)
}

func main() {
	// Example 1: Standard LeetCode test case
	input1 := "Let's take LeetCode contest"
	output1 := reverseWords(input1)
	fmt.Printf("Input:  %q\n", input1)
	fmt.Printf("Output: %q\n\n", output1)

	// Example 2: Short words and punctuation
	input2 := "Mr Ding"
	output2 := reverseWords(input2)
	fmt.Printf("Input:  %q\n", input2)
	fmt.Printf("Output: %q\n", output2)
}
