package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		bi, bj := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
		if bi == bj {
			return arr[i] < arr[j] // tie-break by value
		}
		return bi < bj // primary: by popcount
	})
	return arr
}

func main() {
	tests := [][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8},
		{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
		{2, 3, 5, 7, 11, 13, 17, 19},
	}
	for _, arr := range tests {
		fmt.Printf("%v\n  -> %v\n", arr, sortByBits(arr))
	}
}
