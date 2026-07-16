package main

import (
	"fmt"
)

// maxSubarraySum calculates the maximum contiguous subarray sum
func maxSubarraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize tracking variables with the first element
	maxSoFar := nums[0]
	currentMax := nums[0]

	// Iterate through the slice starting from index 1
	for i := 1; i < len(nums); i++ {
		// Helper calculation to mimic math.Max for integers
		if nums[i] > currentMax+nums[i] {
			currentMax = nums[i]
		} else {
			currentMax = currentMax + nums[i]
		}

		// Update global maximum if current subarray sum is greater
		if currentMax > maxSoFar {
			maxSoFar = currentMax
		}
	}

	return maxSoFar
}

func main() {
	// Example array containing a mix of positive and negative numbers
	exampleArray := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	result := maxSubarraySum(exampleArray)
	fmt.Printf("Go Solution Result: %d\n", result) // Expected Output: 6 ([4, -1, 2, 1])
}
