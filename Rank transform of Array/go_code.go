package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr) // don't mutate the caller's slice
	sort.Ints(sorted)

	rank := make(map[int]int, len(sorted))
	next := 1
	for _, v := range sorted {
		if _, seen := rank[v]; !seen {
			rank[v] = next
			next++
		}
	}

	out := make([]int, len(arr))
	for i, v := range arr {
		out[i] = rank[v]
	}
	return out
}

func main() {
	fmt.Println(arrayRankTransform([]int{40, 10, 20, 30}))                    // [4 1 2 3]
	fmt.Println(arrayRankTransform([]int{100, 100, 100}))                     // [1 1 1]
	fmt.Println(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12})) // [5 3 4 2 8 6 7 1 3]
	fmt.Println(arrayRankTransform([]int{}))                                  // []
}
