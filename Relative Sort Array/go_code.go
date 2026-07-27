package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	frequencies := make([]int, 1001)
	for _, num := range arr1 {
		frequencies[num]++
	}

	result := make([]int, 0, len(arr1))

	for _, num := range arr2 {
		for frequencies[num] > 0 {
			result = append(result, num)
			frequencies[num]--
		}
	}

	for num := 0; num <= 1000; num++ {
		for frequencies[num] > 0 {
			result = append(result, num)
			frequencies[num]--
		}
	}

	return result
}

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}

	fmt.Println(relativeSortArray(arr1, arr2))
}
