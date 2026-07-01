package main

import "fmt"

func removeDuplicateLetters(s string) string {
	var last [26]int
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i // last occurrence index
	}

	var inStack [26]bool
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if inStack[c-'a'] {
			continue // already placed, skip
		}
		// pop a larger top if it reappears later than i
		for len(stack) > 0 && stack[len(stack)-1] > c && last[stack[len(stack)-1]-'a'] > i {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			inStack[top-'a'] = false
		}
		stack = append(stack, c)
		inStack[c-'a'] = true
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicateLetters("bcabc"))    // abc
	fmt.Println(removeDuplicateLetters("cbacdcbc")) // acdb
	fmt.Println(removeDuplicateLetters("abacb"))    // abc
}
