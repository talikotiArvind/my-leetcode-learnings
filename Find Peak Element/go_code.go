package main
import "fmt"

func findPeak(nums [] int) int {
    left, right := 0, len(nums) - 1
    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] < nums[mid + 1] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

func main () {
    nums1 := []int{3, 4, 5, 1, 2, 100, 22}
	fmt.Printf("Array: %v -> Peak Position: %d\n", nums1, findPeak(nums1))
}
