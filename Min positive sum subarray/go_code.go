package main

import "fmt"

func minimumSumSubarray(nums []int, l int, r int) int {
	n := len(nums)
	pre := make([]int, n+1)
	for i, x := range nums {
		pre[i+1] = pre[i] + x
	}

	best := -1
	for size := l; size <= r; size++ {
		for i := 0; i+size <= n; i++ {
			s := pre[i+size] - pre[i]
			if s > 0 && (best == -1 || s < best) {
				best = s
			}
		}
	}
	return best
}

func main() {
	cases := []struct {
		nums    []int
		l, r    int
		want    int
	}{
		{[]int{3, -2, 1, 4}, 2, 3, 1},   // [3,-2] sums to 1
		{[]int{-2, 2, -3, 1}, 2, 3, -1}, // nothing positive
		{[]int{1, 2, 3, 4}, 2, 4, 3},    // [1,2]
		{[]int{-1, -1, 5}, 1, 1, 5},     // only single elements allowed
	}

	for _, c := range cases {
		got := minimumSumSubarray(c.nums, c.l, c.r)
		status := "ok"
		if got != c.want {
			status = "FAIL"
		}
		fmt.Printf("%-16v l=%d r=%d -> %2d (want %2d) %s\n",
			c.nums, c.l, c.r, got, c.want, status)
	}
}
