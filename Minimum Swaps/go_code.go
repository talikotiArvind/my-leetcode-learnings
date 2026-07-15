package main

import "fmt"

func minSwaps(nums []int) int64 {
	k := 0
	for _, v := range nums {
		if v == 0 {
			k++
		}
	}
	region := len(nums) - k // final non-zero region = first n-k slots
	var swaps int64
	for i := 0; i < region; i++ {
		if nums[i] == 0 {
			swaps++ // a zero sitting where a non-zero must end up
		}
	}
	return swaps
}

func main() {
	tests := [][]int{
		{0, 1, 0, 3, 12}, // -> 2
		{0, 1, 0, 2},     // -> 1
		{0, 0, 1},        // -> 2
		{1, 2, 3},        // -> 0
	}
	for _, t := range tests {
		fmt.Printf("%v => %d\n", t, minSwaps(t))
	}
}
