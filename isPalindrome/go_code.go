package main

import "fmt"

func is_palindrome(x int) bool {
	// Time complexity: O(log(n))
	// Space complexity: O(1)
	if x < 0 {
		return false
	}
	original := x
	reversed := 0
	for x != 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}
	return original == reversed
}

func reverseString(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func is_palindrome_as_string(x int) bool {
	// Time complexity: O(n)
	// Space complexity: O(n)
	if x < 0 {
		return false
	}
	reversed := reverseString(string(rune(x)))
	fmt.Println(reversed)
	original := string(rune(x))
	return original == reversed
}

func main() {
	x := 12
	result := is_palindrome_as_string(x)
	fmt.Println(result) // Output: true
}
