package main

import "fmt"

func computeLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0
	i := 1

	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func kmpSearch(text string, pattern string) []int {
	var matches []int
	lps := computeLPS(pattern)
	i, j := 0, 0

	for i < len(text) {
		if pattern[j] == text[i] {
			i++
			j++
		}
		if j == len(pattern) {
			matches = append(matches, i-j)
			j = lps[j-1]
		} else if i < len(text) && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return matches
}

func main() {
	txt := "ABABDABACDABABCABAB"
	pat := "ABABCABAB"
	fmt.Printf("KMP Match Indices (Go): %v\n", kmpSearch(txt, pat)) // Output: [10]
}
