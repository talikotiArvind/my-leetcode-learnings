package main

import (
	"fmt"
	"strings"
)

func reversePrefix(word string, ch byte) string {
	idx := strings.IndexByte(word, ch)
	if idx == -1 {
		return word
	}

	b := []byte(word)
	for i, j := 0, idx; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	testCases := []struct {
		word     string
		ch       byte
		expected string
	}{
		{"abcdefd", 'd', "dcbaefd"},
		{"xyxzxe", 'z', "zxyxxe"},
		{"abcd", 'z', "abcd"},
		{"a", 'a', "a"},
	}

	for _, tc := range testCases {
		result := reversePrefix(tc.word, tc.ch)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s: word=%q, ch=%q -> %q (expected %q)\n",
			status, tc.word, tc.ch, result, tc.expected)
	}
}
