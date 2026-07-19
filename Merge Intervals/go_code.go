package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// sort by start
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		cur := intervals[i]

		if cur[0] <= last[1] {
			// overlap — extend the end of the last merged interval
			if cur[1] > last[1] {
				last[1] = cur[1]
			}
		} else {
			// no overlap — start a new interval
			res = append(res, cur)
		}
	}

	return res
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

	fmt.Println("Input: ", intervals)
	fmt.Println("Merged:", merge(intervals))
	// Merged: [[1 6] [8 10] [15 18]]

	// edge case: fully nested + touching
	intervals2 := [][]int{{1, 4}, {4, 5}, {2, 3}}
	fmt.Println("\nInput: ", intervals2)
	fmt.Println("Merged:", merge(intervals2))
	// Merged: [[1 5]]
}
