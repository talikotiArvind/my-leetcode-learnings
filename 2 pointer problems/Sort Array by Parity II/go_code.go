package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	even, odd, n := 0, 1, len(nums)
	for even < n {
		if nums[even]%2 == 0 {
			even += 2
		} else {
			for nums[odd]%2 == 1 {
				odd += 2
			}
			nums[even], nums[odd] = nums[odd], nums[even]
		}
	}
	return nums
}

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7})) // [4 5 2 7]
}
