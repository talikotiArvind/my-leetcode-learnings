package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)

	count := 0
	for i := range heights {
		if heights[i] != expected[i] {
			count++
		}
	}
	return count
}

func main() {
	tests := [][]int{
		{1, 1, 4, 2, 1, 3},
		{5, 1, 2, 3, 4},
		{1, 2, 3, 4, 5},
	}

	for _, heights := range tests {
		res := heightChecker(heights)
		fmt.Printf("heights=%-20v → mismatches=%d\n", heights, res)
	}
}
