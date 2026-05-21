package main

import "fmt"

func findErrorNums(nums []int) []int {
	n := len(nums)
	seen := make(map[int]bool)
	dup, actualSum := 0, 0

	for _, x := range nums {
		if seen[x] {
			dup = x
		}
		seen[x] = true
		actualSum += x
	}

	expectedSum := n * (n + 1) / 2
	missing := expectedSum - (actualSum - dup)
	return []int{dup, missing}
}

func main() {
	tests := [][]int{
		{1, 2, 2, 4},
		{1, 1},
		{3, 2, 3, 4, 5},
	}

	for _, nums := range tests {
		res := findErrorNums(nums)
		fmt.Printf("nums=%-18v → dup=%d, missing=%d\n", nums, res[0], res[1])
	}
}
