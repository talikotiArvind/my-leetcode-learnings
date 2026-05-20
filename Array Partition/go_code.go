package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	// 1. Sort the slice in ascending order
	sort.Ints(nums)
	
	maxSum := 0
	
	// 2. Loop through even indices and add to sum
	for i := 0; i < len(nums); i += 2 {
		maxSum += nums[i]
	}
	
	return maxSum
}

func main() {
	// Example Input
	input := []int{6, 2, 6, 5, 1, 2}
	
	// Call the function
	result := arrayPairSum(input)
	fmt.Printf("Sorted array: %v\n", input) // [1 2 2 5 6 6]
	fmt.Printf("Maximum pair sum: %d\n", result) // 9
}
