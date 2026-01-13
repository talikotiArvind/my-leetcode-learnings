package main

import "fmt"

func twoSum(nums []int, target int) []int {
	output := make(map[int]int)
	for i, num := range nums {
		value := target - num
		if idx, found := output[value]; found {
			return []int{idx, i}
		}
		output[num] = i
	}
	return nil
}

func main() {
	nums := []int{3, 7, 11, 15}
	target := 10
	result := twoSum(nums, target)
	fmt.Println(result)
}
