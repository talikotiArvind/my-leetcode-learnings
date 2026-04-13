package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	charToWord := make(map[rune]string)
	wordToChar := make(map[string]rune)

	for i, c := range pattern {
		w := words[i]
		mappedWord, charExists := charToWord[c]
		mappedChar, wordExists := wordToChar[w]

		if charExists != wordExists {
			return false
		}
		if charExists && (mappedWord != w || mappedChar != c) {
			return false
		}

		charToWord[c] = w
		wordToChar[w] = c
	}

	return true
}

func main() {
	testCases := []struct {
		pattern  string
		s        string
		expected bool
	}{
		{"aabb", "dog dog cat cat", true},   // normal match
		{"abba", "dog cat cat dog", true},   // classic example
		{"ab",   "dog dog",         false},  // same word, different chars
		{"aa",   "dog cat",         false},  // same char, different words
		{"aab",  "dog dog dog",     false},  // a→dog, b→dog conflict
		{"abcd", "dog cat",         false},  // length mismatch
	}

	for _, tc := range testCases {
		result := wordPattern(tc.pattern, tc.s)
		status := "PASS"
		if result != tc.expected {
			status = "❌ FAIL"
		}
		fmt.Printf("%s | pattern=%-6s s=%-20s expected=%-5v got=%v\n",
			status, tc.pattern, tc.s, tc.expected, result)
	}
}
