package main

import "fmt"

func subsets(nums []int) [][]int {
	result := [][]int{}

	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		subset := make([]int, len(path))
		copy(subset, path)
		result = append(result, subset)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, []int{})
	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	// [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
	fmt.Println(subsets([]int{0})) // [[] [0]]
}