
package main
import "fmt"

func findPivotIndex(nums [] int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    
    left := 0
    
    for i, n := range nums {
        if left == total-left-n {
            return i
        }
        left += n
    }
    return -1
}

func main() {
  data := []int{1,7,3,6,5,6}
  fmt.Println(findPivotIndex(data))
}
