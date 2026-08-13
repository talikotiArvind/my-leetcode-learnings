package main

import "fmt"

func generateParenthesis(n int) []string {
	result := []string{}

	var backtrack func(current string, open, close int)
	backtrack = func(current string, open, close int) {
		if len(current) == 2*n {
			result = append(result, current)
			return
		}
		if open < n {
			backtrack(current+"(", open+1, close)
		}
		if close < open {
			backtrack(current+")", open, close+1)
		}
	}

	backtrack("", 0, 0)
	return result
}

func main() {
	fmt.Println(generateParenthesis(3))
	// [((())) (()()) (())() ()(()) ()()()]
	fmt.Println(generateParenthesis(1)) // [()]
}