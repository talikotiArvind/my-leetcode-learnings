package main

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	globalMax := nums[0]
	currentMax := nums[0]
	currentMin := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// Swap max and min if current element is negative
		if num < 0 {
			currentMax, currentMin = currentMin, currentMax
		}

		// Calculate new local boundaries
		prod1 := currentMax * num
		prod2 := currentMin * num

		currentMax = num
		if prod1 > currentMax {
			currentMax = prod1
		}

		currentMin = num
		if prod2 < currentMin {
			currentMin = prod2
		}

		if currentMax > globalMax {
			globalMax = currentMax
		}
	}

	return globalMax
}

func main() {
	example := []int{2, 3, -2, 4}
	fmt.Printf("Max Product (Go): %d\n", maxProduct(example)) // Output: 6
}
