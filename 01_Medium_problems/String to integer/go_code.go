package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	// 1. Remove leading whitespaces
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	// 2. Check for optional sign
	sign := 1
	i := 0
	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	// 3. Read valid digit characters and handle overflow
	res := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '1' {
		digit := int(s[i] - '0')

		// Check for potential overflow before multiplying res by 10
		if res > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		res = res*10 + digit
		i++
	}

	return res * sign
}

func main() {
	// Test cases covering different edge cases
	testCases := []string{
		"42",
		"   -42",
		"1337c0d3",
		"0-1",
		"words and 987",
		"-91283472332", // Underflow test case
		"91283472332",  // Overflow test case
	}

	fmt.Printf("%-20s -> %s\n", "Input String", "Parsed Integer")
	fmt.Println(strings.Repeat("-", 40))

	for _, str := range testCases {
		result := myAtoi(str)
		fmt.Printf("%-20q -> %d\n", str, result)
	}
}
