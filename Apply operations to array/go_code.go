package main

import "fmt"

func applyOperations(nums []int) []int {
	n := len(nums)

	// Step 1: apply the merge operation
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}

	// Step 2: shift non-zero elements to the front
	result := make([]int, n)
	pos := 0
	for _, num := range nums {
		if num != 0 {
			result[pos] = num
			pos++
		}
	}

	return result
}

func main() {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 2, 1, 1, 0}, []int{1, 4, 2, 0, 0, 0}},
		{[]int{0, 1}, []int{1, 0}},
		{[]int{1, 1, 0, 0, 1}, []int{2, 1, 0, 0, 0}},
		{[]int{2, 2, 2, 2}, []int{4, 4, 0, 0}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
	}

	for _, tc := range testCases {
		result := applyOperations(tc.nums)
		status := "PASS"
		if !equal(result, tc.expected) {
			status = "FAIL"
		}
		fmt.Printf("%s: -> %v (expected %v)\n", status, result, tc.expected)
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
