package main

import (
	"fmt"
)

func longestPalindrome(s string) int {
	counts := make(map[rune]int)
	for _, ch := range s {
		counts[ch]++
	}

	length := 0
	oddFound := false

	for _, freq := range counts {
		length += (freq / 2) * 2
		if freq%2 == 1 {
			oddFound = true
		}
	}

	if oddFound {
		length++
	}

	return length
}

func main() {
	tests := []string{
		"abccccdd", // 7
		"a",        // 1
		"aabb",     // 4
		"ccc",      // 3
	}

	for _, s := range tests {
		fmt.Printf("s=%q  longest=%d\n", s, longestPalindrome(s))
	}
}
