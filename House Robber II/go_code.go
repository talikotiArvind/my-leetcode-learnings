package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	robLinear := func(houses []int) int {
		prev, curr := 0, 0
		for _, money := range houses {
			prev, curr = curr, maxInt(curr, prev+money)
		}
		return curr
	}
	return maxInt(robLinear(nums[:len(nums)-1]), robLinear(nums[1:]))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	check := func(got, want int) {
		if got != want {
			panic(fmt.Sprintf("got %d, want %d", got, want))
		}
	}
	check(rob([]int{2, 3, 2}), 3)
	check(rob([]int{1, 2, 3, 1}), 4)
	check(rob([]int{1, 1, 1, 1}), 2)
	check(rob([]int{5}), 5)
	check(rob([]int{0, 0}), 0)
	fmt.Println("All tests passed.")
}
