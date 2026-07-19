package main

import "fmt"

// Approach 1: used boolean array
func permute(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	path := []int{}

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			// copy path — don't append the slice header directly
			comb := make([]int, len(path))
			copy(comb, path)
			res = append(res, comb)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack()
			path = path[:len(path)-1] // undo
			used[i] = false
		}
	}

	backtrack()
	return res
}

// Approach 2: in-place swapping (no used array)
func permuteSwap(nums []int) [][]int {
	res := [][]int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(nums) {
			comb := make([]int, len(nums))
			copy(comb, nums)
			res = append(res, comb)
			return
		}
		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			backtrack(start + 1)
			nums[start], nums[i] = nums[i], nums[start] // swap back
		}
	}

	backtrack(0)
	return res
}

func main() {
	nums := []int{1, 2, 3}

	fmt.Println("Approach 1 (used array):")
	for _, p := range permute(nums) {
		fmt.Println(p)
	}

	fmt.Println("\nApproach 2 (swapping):")
	for _, p := range permuteSwap(nums) {
		fmt.Println(p)
	}
}
