package main

import (
	"fmt"
	"math"
)

func find132pattern(nums []int) bool {
	stack := []int{}
	third := math.MinInt64
	for i := len(nums) - 1; i >= 0; i-- {
		n := nums[i]
		if n < third {
			return true
		}
		for len(stack) > 0 && stack[len(stack)-1] < n {
			third = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, n)
	}
	return false
}

func main() {
	fmt.Println(find132pattern([]int{1, 2, 3, 4}))
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))
	fmt.Println(find132pattern([]int{-1, 3, 2, 0}))
}
