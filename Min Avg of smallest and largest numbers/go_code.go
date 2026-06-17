package main
import "fmt"
import "sort"

func findMinAverage(nums [] int) float64 {
    sort.Ints(nums)
    n := len(nums)
    minAvg := float64(nums[n-1])
    for i := 0; i < n/2; i++ {
        curAvg := float64((nums[i] + nums[n-1-i])) / 2.0
        if curAvg < minAvg {minAvg = curAvg}
    }
    return minAvg
}

func main() {
  nums := []int{7, 8, 3, 4, 15, 13, 4, 1}
  fmt.Println(findMinAverage(nums))
}
