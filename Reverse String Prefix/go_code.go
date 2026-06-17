package main

import (
	"fmt"
)
func ReversePrefixByCount(s string, k int) string {
	// Guard against out-of-bound errors
	if k <= 1 || k > len(s) {
		return s
	}

	// 1. Extract prefix slice and convert to bytes
	prefix := []byte(s[:k])
	
	// 2. Reverse the prefix slice
	for i, j := 0, len(prefix)-1; i < j; i, j = i+1, j-1 {
		prefix[i], prefix[j] = prefix[j], prefix[i]
	}

	// 3. Merge the reversed prefix with the remaining suffix
	return string(prefix) + s[k:]
}

func main() {
	str := "abcdefg"
	k := 4

	result := ReversePrefixByCount(str, k)
	fmt.Printf("Original: %s, Count: %d\n", str, k)
	fmt.Printf("Result:   %s\n", result) // Output: dcbaefg
}
