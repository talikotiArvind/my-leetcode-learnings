package main

import "fmt"

func findIndices(nums []int, indexDiff int, valueDiff int) []int {
	minIdx, maxIdx := 0, 0 // indices of min/max in the eligible prefix
	for j := indexDiff; j < len(nums); j++ {
		i := j - indexDiff // newly eligible index
		if nums[i] < nums[minIdx] {
			minIdx = i
		}
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
		if nums[maxIdx]-nums[j] >= valueDiff {
			return []int{maxIdx, j}
		}
		if nums[j]-nums[minIdx] >= valueDiff {
			return []int{minIdx, j}
		}
	}
	return []int{-1, -1}
}

func main() {
	tests := []struct {
		nums               []int
		indexDiff, valueDiff int
	}{
		{[]int{5, 1, 4, 1}, 2, 4},   // -> [0 3]
		{[]int{2, 1}, 0, 0},         // -> [0 0]
		{[]int{1, 2, 3}, 2, 4},      // -> [-1 -1]
	}
	for _, t := range tests {
		fmt.Printf("nums=%v indexDiff=%d valueDiff=%d => %v\n",
			t.nums, t.indexDiff, t.valueDiff, findIndices(t.nums, t.indexDiff, t.valueDiff))
	}
}
