package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	rows := []map[rune]bool{}
	for _, row := range []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"} {
		m := make(map[rune]bool)
		for _, ch := range row {
			m[ch] = true
		}
		rows = append(rows, m)
	}

	res := []string{}
	for _, word := range words {
		lower := strings.ToLower(word)
		for _, row := range rows {
			inRow := true
			for _, ch := range lower {
				if !row[ch] {
					inRow = false
					break
				}
			}
			if inRow {
				res = append(res, word)
				break
			}
		}
	}

	return res
}

func main() {
	tests := [][]string{
		{"Hello", "Alaska", "Dad", "Peace"},
		{"omk"},
		{"adsf", "sfd"},
	}

	for _, words := range tests {
		fmt.Printf("input=%v  output=%v\n", words, findWords(words))
	}
}
