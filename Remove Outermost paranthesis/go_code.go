package main

import (
	"fmt"
	"strings"
)

func removeOuterParentheses(s string) string {
	var sb strings.Builder
	depth := 0
	for _, ch := range s {
		if ch == '(' {
			if depth > 0 {
				sb.WriteRune(ch)
			}
			depth++
		} else {
			depth--
			if depth > 0 {
				sb.WriteRune(ch)
			}
		}
	}
	return sb.String()
}

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))         // ()()()
	fmt.Println(removeOuterParentheses("(()())(())(()(()))")) // ()()()()(())
	fmt.Println(removeOuterParentheses("()()"))               // (empty string)
}
