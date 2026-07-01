package main

import "fmt"

func checkValidString(s string) bool {
	low, high := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			low++
			high++
		case ')':
			low--
			high--
		default: // '*'
			low-- // could be ')'
			high++ // could be '('
		}
		if high < 0 { // unmatchable ')'
			return false
		}
		if low < 0 { // clamp: count can't go below 0
			low = 0
		}
	}
	return low == 0
}

func main() {
	fmt.Println(checkValidString("()"))    // true
	fmt.Println(checkValidString("(*)"))   // true
	fmt.Println(checkValidString("(*))"))  // true
	fmt.Println(checkValidString(")("))    // false
	fmt.Println(checkValidString("(((**")) // false
}
