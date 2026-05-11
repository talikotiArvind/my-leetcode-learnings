package main

import (
	"fmt"
	"strings"
)

func detectCapitalUse(word string) bool {
	return word == strings.ToUpper(word) ||
		word == strings.ToLower(word) ||
		word == strings.ToTitle(string(word[0]))+strings.ToLower(word[1:])
}

func main() {
	fmt.Println(detectCapitalUse("USA"))      // true
	fmt.Println(detectCapitalUse("leetcode")) // true
	fmt.Println(detectCapitalUse("Google"))   // true
	fmt.Println(detectCapitalUse("FlaG"))     // false
}
