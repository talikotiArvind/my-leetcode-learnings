package main

import "fmt"

func maxSubArray(nums []int) int {
	best, current := nums[0], nums[0]
	for _, x := range nums[1:] {
		if current+x > x {
			current = current + x
		} else {
			current = x
		}
		if current > best {
			best = current
		}
	}
	return best
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(maxSubArray([]int{1}))                             // 1
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))                // 23
}