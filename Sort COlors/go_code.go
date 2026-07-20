package main

import (
	"fmt"
)

// sortColors sorts an array of 0s, 1s, and 2s in-place.
func sortColors(nums []int) {
	low := 0
	mid := 0
	high := len(nums) - 1

	for mid <= high {
		switch nums[mid] {
		case 0:
			// Swap mid with low, move both pointers forward
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			// Element is in the correct place, just move mid forward
			mid++
		case 2:
			// Swap mid with high, move high pointer backward
			// Do not move mid forward yet, as the swapped element needs checking
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}

func main() {
	// Example Input: Red (0), White (1), and Blue (2)
	colors := []int{2, 0, 2, 1, 1, 0}
	
	fmt.Printf("Original array: %v\n", colors)
	
	sortColors(colors)
	
	fmt.Printf("Sorted array:   %v\n", colors)
}

