package main

import "fmt"

func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for nums[i] >= 1 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			targetIdx := nums[i] - 1
			nums[i], nums[targetIdx] = nums[targetIdx], nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

func main() {
	case1 := []int{3, 4, -1, 1}
	fmt.Printf("Input: %v -> First Missing Positive: %d\n", case1, firstMissingPositive(case1))

	case2 := []int{1, 2, 3}
	fmt.Printf("Input: %v -> First Missing Positive: %d\n", case2, firstMissingPositive(case2))

	case3 := []int{7, 8, 9, 11, 12}
	fmt.Printf("Input: %v -> First Missing Positive: %d\n", case3, firstMissingPositive(case3))

	case4 := []int{1, 1, 2, 2}
	fmt.Printf("Input: %v -> First Missing Positive: %d\n", case4, firstMissingPositive(case4))
}
