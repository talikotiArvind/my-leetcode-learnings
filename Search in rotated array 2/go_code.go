package main

import "fmt"

func search(nums []int, target int) bool {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := (lo + hi) / 2

		if nums[mid] == target {
			return true
		}

		if nums[lo] == nums[mid] {
			lo++
			continue
		}

		if nums[lo] <= nums[mid] {
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return false
}

func main() {
	tests := []struct {
		nums   []int
		target int
		want   bool
	}{
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{1, 0, 1, 1, 1}, 0, true},
		{[]int{1, 1, 1, 1, 1}, 0, false},
		{[]int{1}, 1, true},
		{[]int{1}, 0, false},
	}

	for _, tc := range tests {
		got := search(tc.nums, tc.target)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("[%s] nums=%-25v target=%d => %v\n", status, tc.nums, tc.target, got)
	}
}
