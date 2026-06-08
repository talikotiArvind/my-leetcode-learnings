package main
import "fmt"
import "sort"

func targetInts(nums [] int, target int) [] int {
    sort.Ints(nums)
    
    left := sort.SearchInts(nums, target)
    right := sort.SearchInts(nums, target+1)
    output := []int {}
    
    for i := left; i < right; i++ {
        output = append(output, i)
    }
    return output
}

func main() {
    nums := []int {1, 2, 5, 2, 3}
    pivot := 2
    fmt.Println(targetInts(nums, pivot))
}
