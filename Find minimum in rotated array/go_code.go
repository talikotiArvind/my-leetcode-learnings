package main
import "fmt"

func findMin(nums [] int) int {
    left, right := 0, len(nums) - 1
    for left < right {
        mid := (left + right) / 2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return nums[left]
}

func main () {
    nums1 := []int{3, 4, 5, 1, 2, 0}
	fmt.Printf("Array: %v -> Minimum: %d\n", nums1, findMin(nums1))
}
