package main
import "fmt"

func maxConsecutiveOnes (nums [] int) int {
    var count, maxCount int
    for _, val:= range nums {
        if val == 1 {
            count += 1
        } else {
            maxCount = max(maxCount, count)
            count = 0
        }
    }
    return max(maxCount, count)
}

func main() {
  fmt.Println("Try programiz.pro")
  nums1 := [] int {1,1,0,1,1,1,1}
  fmt.Println(maxConsecutiveOnes(nums1))
}
