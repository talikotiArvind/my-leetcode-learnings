package main
import "fmt"

func containsDuplicate(nums [] int, k int) bool {
    seen := make(map[int]int)
    
    for i, num := range nums {
      if j, ok := seen[num]; ok &&i-j <= k {
          return true
      }
      seen[num] = i
    }
    return false
}

func main () {
    nums := []int {1,2,3}
    k := 3
    fmt.Println(containsDuplicate(nums, k)) //false
}
