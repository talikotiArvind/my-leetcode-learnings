package main

import (
	"fmt"
	"math"
	"slices"
)

func countElements(arr1 []int, arr2 []int, d int) int {
	// Sort both slices in place
	slices.Sort(arr1)
	slices.Sort(arr2)

	i, j := 0, 0
	count := 0

	// Two-pointer traversal
	for i < len(arr1) && j < len(arr2) {
		// math.Abs works on float64, so we cast to float64 or do manual diff
		if int(math.Abs(float64(arr1[i]-arr2[j]))) <= d {
			i++
		} else if arr1[i] > arr2[j] {
			j++
		} else {
			i++
			count++
		}
	}

	// Add the remaining elements of arr1 if any
	if i < len(arr1) {
		count += len(arr1) - i
	}

	return count
}

func main() {
	// Example test case
	arr1 := []int{4, 5, 8}
	arr2 := []int{10, 9, 1, 8}
	d := 2

	result := countElements(arr1, arr2, d)
	fmt.Printf("Resulting count: %d\n", result)
}
