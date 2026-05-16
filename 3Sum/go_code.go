package main
import (
    "fmt"
    "sort"
    )


func threeSum(nums []int)[][] int {
     sort.Ints(nums)
     result := [][] int {}
     
     for i, a := range(nums) {
         if i > 0 && nums[i] == nums[i - 1]{
             continue
         }
         if a > 0 {break}
         left, right := i+1, len(nums) - 1
         
         for left < right {
             total := nums[i] + nums[left] + nums[right]
             if total < 0 {
                 left ++
             } else if total > 0 {
                 right --
             } else {
                 result = append(result, []int{a, nums[left], nums[right]})
                 left ++
                 for left < right && nums[left] == nums[left + 1] {
                    left ++
                 }
             }
         }
     }
     return result
 }
 
 func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1 -1 2] [-1 0 1]]
	fmt.Println(threeSum([]int{0, 0, 0}))              // [[0 0 0]]
	fmt.Println(threeSum([]int{0, 1, 1}))              // []
	fmt.Println(threeSum([]int{-1, -1, -1, 0, 1, 2})) // [[-1 -1 2] [-1 0 1]]
}
