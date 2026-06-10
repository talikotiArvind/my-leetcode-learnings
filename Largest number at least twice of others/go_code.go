package main

import "fmt"

func dominantIndex(nums []int) int {
	largest, second := -1, -1
	largestIdx := -1

	for i, num := range nums {
		if num > largest {
			second = largest
			largest = num
			largestIdx = i
		} else if num > second {
			second = num
		}
	}

	if largest >= 2*second {
		return largestIdx
	}
	return -1
}

func main() {
	fmt.Println(dominantIndex([]int{3, 6, 1, 0})) // 1
	fmt.Println(dominantIndex([]int{1, 2, 3, 4})) // -1
	fmt.Println(dominantIndex([]int{1}))          // 0
}
