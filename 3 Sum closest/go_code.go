package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	res := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

		minSum := nums[i] + nums[i+1] + nums[i+2]
		if minSum >= target {
			if abs(minSum-target) < abs(res-target) {
				res = minSum
			}
			break
		}

		maxSum := nums[i] + nums[n-1] + nums[n-2]
		if maxSum <= target {
			if abs(maxSum-target) < abs(res-target) {
				res = maxSum
			}
			continue
		}

		for left < right {
			s := nums[i] + nums[left] + nums[right]
			if abs(s-target) < abs(res-target) {
				res = s
			}
			if s < target {
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else if s > target {
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else {
				return target
			}
		}
	}

	return res
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	tests := []struct {
		nums   []int
		target int
	}{
		{[]int{-1, 2, 1, -4}, 1},
		{[]int{0, 0, 0}, 1},
		{[]int{1, 1, 1, 0}, 100},
	}

	for _, t := range tests {
		fmt.Printf("nums=%v  target=%d  closest=%d\n",
			t.nums, t.target, threeSumClosest(t.nums, t.target))
	}
}
