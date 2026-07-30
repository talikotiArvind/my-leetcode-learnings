package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	neg := num < 0
	if neg {
		num = -num
	}
	digits := []byte{}
	for num > 0 {
		digits = append(digits, byte('0'+num%7)) // remainder is the next digit
		num /= 7
	}
	if neg {
		digits = append(digits, '-')
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 { // reverse in place
		digits[i], digits[j] = digits[j], digits[i]
	}
	return string(digits)
}

func main() {
	examples := []int{100, -7, 0, 7}
	for _, num := range examples {
		fmt.Printf("num = %d -> base7 = %s\n", num, convertToBase7(num))
	}
}
