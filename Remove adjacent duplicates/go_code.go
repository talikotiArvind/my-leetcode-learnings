package main

import "fmt"

func removeDuplicates(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1] // pop the match
		} else {
			stack = append(stack, s[i]) // push
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicates("abbaca")) // ca
	fmt.Println(removeDuplicates("azxxzy")) // ay
	fmt.Println(removeDuplicates("aaaaaa")) // (empty string)
}
