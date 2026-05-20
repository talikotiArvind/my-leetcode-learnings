package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0

	expand := func(left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return right - left - 1
	}

	for i := 0; i < len(s); i++ {
		len1 := expand(i, i)
		len2 := expand(i, i+1)

		maxLen := len1
		if len2 > len1 {
			maxLen = len2
		}

		if maxLen > (end - start) {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func main() {
	input := "babad"
	output := longestPalindrome(input)
	fmt.Printf("Input: %s\nOutput: %s\n", input, output)
}
