package main

import (
	"fmt"
	"math"
)

func maximumProduct(nums []int) int {
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	min1, min2 := math.MaxInt32, math.MaxInt32

	for _, n := range nums {
		if n > max1 {
			max3 = max2
			max2 = max1
			max1 = n
		} else if n > max2 {
			max3 = max2
			max2 = n
		} else if n > max3 {
			max3 = n
		}

		if n < min1 {
			min2 = min1
			min1 = n
		} else if n < min2 {
			min2 = n
		}
	}

	prod1 := max1 * max2 * max3
	prod2 := min1 * min2 * max1

	if prod1 > prod2 {
		return prod1
	}
	return prod2
}

func main() {
	input := []int{-10, -10, 1, 3, 2}
	result := maximumProduct(input)
	fmt.Printf("Input Array: %v\n", input)
	fmt.Printf("Maximum Product of Three Numbers: %d\n", result)
}
