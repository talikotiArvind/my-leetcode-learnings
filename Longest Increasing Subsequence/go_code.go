package main

import (
	"fmt"
	"sort"
)

func lengthOfLIS(nums []int) int {
	tails := []int{}
	for _, num := range nums {
		i := sort.SearchInts(tails, num)
		if i == len(tails) {
			tails = append(tails, num)
		} else {
			tails[i] = num
		}
	}
	return len(tails)
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))           // 4
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))        // 1
}
