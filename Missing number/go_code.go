package main

import "fmt"

func missingNumber(nums []int) int {
	result := len(nums)
	for i, x := range nums {
		result ^= i ^ x
	}
	return result
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))                      // 2
	fmt.Println(missingNumber([]int{0, 1}))                         // 2
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))    // 8
}