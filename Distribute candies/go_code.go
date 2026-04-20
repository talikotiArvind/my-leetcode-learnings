package main
import "fmt"

func distributeCandies(nums1 [] int) int {
    seen:= make(map[int]struct{})
    
    for _, c := range nums1 {
        seen[c] = struct{}{}
    }
    
    return min(len(seen), len(nums1)/2)
}

func main() {
    candies1 := []int{1, 1, 2, 2, 3, 3}
    fmt.Println(distributeCandies(candies1))
    
    candies2 := []int{1, 1, 2, 3}
    fmt.Println(distributeCandies(candies2))

    candies3 := []int{6, 6, 6, 6}
    fmt.Println(distributeCandies(candies3))
}
