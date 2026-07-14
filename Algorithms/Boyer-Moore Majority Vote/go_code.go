package main

import "fmt"

// majorityElement returns the element appearing > n/2 times (assumes one exists).
func majorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, x := range nums {
		if count == 0 {
			candidate = x
		}
		if x == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func main() {
	votes := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Printf("array    : %v\n", votes)
	fmt.Printf("majority : %d\n", majorityElement(votes))
}
