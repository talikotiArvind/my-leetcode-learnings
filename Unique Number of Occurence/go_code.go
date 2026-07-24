package main
import "fmt"

func uniqueOccurrences(nums [] int) bool {
    freq:= make(map[int]int)
    for _, v:= range nums{
        freq[v]++
    }
    
    seen := make(map[int]struct{}, len(freq))
    for _, c := range freq {
        if _, dup := seen[c]; dup {
            return false
        }
        seen[c] = struct{}{}
    }
    return true
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))                     // true
	fmt.Println(uniqueOccurrences([]int{1, 2}))                                 // false
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))     // true
}
