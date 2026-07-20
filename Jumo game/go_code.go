package main

import (
	"fmt"
)

// canJump determines if you can reach the last index of the array.
func canJump(nums []int) bool {
	farthestReachable := 0
	lastIndex := len(nums) - 1

	for currentPosition, jumpLength := range nums {
		// 1. Check if the current position is unreachable
		if currentPosition > farthestReachable {
			return false
		}

		// 2. Update the farthest position we can reach
		potentialReach := currentPosition + jumpLength
		if potentialReach > farthestReachable {
			farthestReachable = potentialReach
		}

		// 3. Check for early success
		if farthestReachable >= lastIndex {
			return true
		}
	}

	return farthestReachable >= lastIndex
}

func main() {
	// Example 1: True case (Can reach the end)
	trueCase := []int{2, 3, 1, 1, 4}
	fmt.Printf("Array: %v\n", trueCase)
	fmt.Printf("Can reach the end? %t\n\n", canJump(trueCase))

	// Example 2: False case (Cannot reach the end)
	falseCase := []int{3, 2, 1, 0, 4}
	fmt.Printf("Array: %v\n", falseCase)
	fmt.Printf("Can reach the end? %t\n", canJump(falseCase))
}
