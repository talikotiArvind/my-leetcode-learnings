package main

import "fmt"

func rob(nums []int) int {
	prev, curr := 0, 0
	for _, x := range nums {
		take := prev + x
		if curr > take {
			take = curr
		}
		prev, curr = curr, take
	}
	return curr
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))    // 4
	fmt.Println(rob([]int{2, 7, 9, 3, 1})) // 12
	fmt.Println(rob([]int{5}))             // 5
}