package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	sort.Ints(nums)

	ans := math.MaxInt
	for i := 0; i+k-1 < len(nums); i++ {
		if d := nums[i+k-1] - nums[i]; d < ans {
			ans = d
		}
	}
	return ans
}

func main() {
	cases := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{90}, 1, 0},                    // single student
		{[]int{9, 4, 1, 7}, 2, 2},            // pick 7 and 9
		{[]int{4, 1, 7, 9, 8}, 3, 2},         // window 7,8,9
		{[]int{1, 100, 101, 102, 200}, 3, 2}, // tight cluster in the middle
	}

	for _, c := range cases {
		got := minimumDifference(c.nums, c.k)
		status := "ok"
		if got != c.want {
			status = "FAIL"
		}
		fmt.Printf("%-22v k=%d -> %d (want %d) %s\n", c.nums, c.k, got, c.want, status)
	}
}
