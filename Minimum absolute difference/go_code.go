package main

import (
	"fmt"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	minDiff := 1 << 31 // large sentinel; constraints keep diffs well below this
	res := [][]int{}

	for i := 0; i < len(arr)-1; i++ {
		d := arr[i+1] - arr[i]
		if d < minDiff {
			minDiff = d
			res = [][]int{{arr[i], arr[i+1]}} // new minimum: reset
		} else if d == minDiff {
			res = append(res, []int{arr[i], arr[i + 1]}) // tie: append
		}
	}
	return res
}

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))       // [[1 2] [2 3] [3 4]]
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))  // [[1 3]]
	fmt.Println(minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27})) // [[-14 -10] [19 23] [23 27]]
}
