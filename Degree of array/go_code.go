package main

import "fmt"

func findShortestSubArray(nums []int) int {
	first := make(map[int]int)
	count := make(map[int]int)
	degree := 0
	result := 0

	for i, x := range nums {
		if _, seen := first[x]; !seen {
			first[x] = i
		}
		count[x]++

		if count[x] > degree {
			degree = count[x]
			result = i - first[x] + 1
		} else if count[x] == degree {
			result = min(result, i-first[x]+1)
		}
	}

	return result
}

func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))          // 2
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))    // 6
	fmt.Println(findShortestSubArray([]int{1}))                      // 1
}
