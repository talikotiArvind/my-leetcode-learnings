package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	prev, curr := 1, 1
	for i := 1; i < len(s); i++ {
		temp := 0
		if s[i] != '0' {
			temp += curr
		}
		two, _ := strconv.Atoi(s[i-1 : i+1])
		if two >= 10 && two <= 26 {
			temp += prev
		}
		prev, curr = curr, temp
	}
	return curr
}

func main() {
	fmt.Println(numDecodings("12"))  // 2
	fmt.Println(numDecodings("226")) // 3
	fmt.Println(numDecodings("06"))  // 0
}