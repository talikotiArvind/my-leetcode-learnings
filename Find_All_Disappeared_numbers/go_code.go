package main
import "fmt"

func findDisappearedNumbers(nums [] int) [] int {
    n := len(nums)
    var result [] int
    seen := make([]bool, n+1)
    for _, num := range(nums) {
        seen[num] = true
    }
    
    for i:= 1; i<= n; i++ {
        if !seen[i] {
            result = append(result, i)
        }
    }
    return result
}

func main() {
    input1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result1 := findDisappearedNumbers(input1)
	fmt.Printf("Input: %v\nMissing: %v\n\n", input1, result1)
}
