package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	left := 0
	windowSum := 0
	best := len(nums) + 1 // sentinel: bigger than any real answer

	for right := 0; right < len(nums); right++ {
		windowSum += nums[right] // expand the window to the right

		// shrink from the left while the window is still valid
		for windowSum >= target {
			length := right - left + 1
			if length < best {
				best = length
			}
			windowSum -= nums[left]
			left++
		}
	}

	if best == len(nums)+1 {
		return 0 // no valid window found
	}
	return best
}

func main() {
	// Example 1: subarray [4,3] has sum 7, length 2 -> answer 2
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})) // 2

	// Example 2: single element 4 >= 4, length 1 -> answer 1
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4})) // 1

	// Example 3: total sum is 1+1+1+1+1+1+1 = 7 < 11, no window -> answer 0
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1})) // 0
}
