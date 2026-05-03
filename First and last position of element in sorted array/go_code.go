package main
import "fmt"

func firstLastPosition (nums [] int, target int) [] int {
    binSearch := func(findFirst bool) int {
        l, r, result := 0, len(nums) - 1, -1
         for l <= r {
             mid := (l+r)/2
             if nums[mid] == target {
                 result = mid
                 if findFirst {
                     r = mid - 1
                 } else {
                     l = mid + 1
                 }
             }else if nums[mid] < target {
                 r = mid - 1
             }else{
                 l = mid + 1
             }
         }
         return result
    }
    return []int{binSearch(true), binSearch(false)}
}

func main() {
  nums:= []int{5,7,8,8,8,10}
  target := 6
  fmt.Println(firstLastPosition(nums, target))
}
