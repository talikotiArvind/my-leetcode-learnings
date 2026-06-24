package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	i, j, n := 0, 1, len(nums)
	for i < n {
		if nums[i]%2 == 0 {
			i += 2 // even value correctly at even index
		} else {
			for nums[j]%2 == 1 { // find an even value sitting at an odd index
				j += 2
			}
			nums[i], nums[j] = nums[j], nums[i]
			i += 2
		}
	}
	return nums
}

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7})) // [4 5 2 7]
	fmt.Println(sortArrayByParityII([]int{2, 3}))       // [2 3]
	fmt.Println(sortArrayByParityII([]int{3, 1, 4, 2})) // [4 1 2 3]
	fmt.Println(sortArrayByParityII([]int{8, 7}))       // [8 7]
}
