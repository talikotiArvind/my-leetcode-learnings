package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ngeMap := make(map[int]int)
	var stack []int

	for _, num := range nums2 {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ngeMap[top] = num
		}
		stack = append(stack, num)
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ngeMap[top] = -1
	}

	result := make([]int, len(nums1))
	for i, num := range nums1 {
		result[i] = ngeMap[num]
	}

	return result
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	result := nextGreaterElement(nums1, nums2)
	fmt.Printf("Result: %v\n", result)
}
