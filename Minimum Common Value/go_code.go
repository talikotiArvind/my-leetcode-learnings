package main
import "fmt"

func minimumCommonValue(nums1 []int, nums2 []int) int {
    i, j := 0, 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] == nums2[j]{
            return nums1[i]
        } else if nums1[i] < nums2[j] {
            i++
        } else {
            j++
        }
    }
    return -1
}

func main() {
  nums1 := [] int {1,2,3,4}
  nums2 := [] int {10, 11}
  fmt.Println(minimumCommonValue(nums1, nums2))
}
