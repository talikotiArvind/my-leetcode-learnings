package main
import "fmt"

func binarySearch(nums [] int, target int) int {
    left, right := 0, len(nums) - 1
    mid := 0
    for left <= right {
        mid = (left + right) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}


func main () {
    array := [] int {1,2,5,8,9,11,55,96}
    target := 99
    fmt.Println(binarySearch(array, target))
}
