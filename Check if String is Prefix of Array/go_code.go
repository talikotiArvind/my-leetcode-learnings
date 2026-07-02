package main

import (
	"fmt"
	"strings"
)

func isPrefixString(s string, words []string) bool {
	var sb strings.Builder

	for _, word := range words {
		sb.WriteString(word)
		concat := sb.String()

		if concat == s {
			return true
		}
		if len(concat) > len(s) {
			return false
		}
	}

	return false
}

func main() {
	testCases := []struct {
		s        string
		words    []string
		expected bool
	}{
		{"iloveleetcode", []string{"i", "love", "leetcode", "apples"}, true},
		{"iloveleetcode", []string{"apples", "i", "love", "leetcode"}, false},
		{"a", []string{"a", "a"}, true},
		{"aaa", []string{"aa", "a"}, true},
		{"iloveleetcode", []string{"i", "love", "leet", "code"}, false},
	}

	for _, tc := range testCases {
		result := isPrefixString(tc.s, tc.words)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("%s: s=%q, words=%v -> %v (expected %v)\n",
			status, tc.s, tc.words, result, tc.expected)
	}
}
