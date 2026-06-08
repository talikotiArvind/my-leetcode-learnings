package main
import "fmt"

func sortWithPivot(nums [] int, pivot int) [] int {
    less := [] int {}
    pivot_data := []int {}
    more := []int {}
    
    for _, value := range nums {
        if value < pivot {
            less = append(less, value)
        } else if value == pivot {
            pivot_data = append(pivot_data, value)
        } else {
            more = append(more, value)
        }
    }
    less =  append(append(less, pivot_data...), more...)
    return less
}

func main() {
    nums := []int {9,12,5,10,14,3,10}
    pivot := 10
    fmt.Println(sortWithPivot(nums, pivot))
}
