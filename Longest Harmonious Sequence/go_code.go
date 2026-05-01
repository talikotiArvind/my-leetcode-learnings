package main
import "fmt"

func longestHarmoniousSequence(nums [] int) int {
    freq := make(map[int]int)
    for _, value := range(nums) {
        freq[value] ++
    }
    best := 0
    
    for num, count := range freq {
        if val, ok := freq[num + 1]; ok {
            if count+val > best {
                best = count + val
            }
        }
    }
    return best
}


func main () {
    //nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
    nums1 := []int{1,1,1,1}
    fmt.Println(longestHarmoniousSequence(nums1))
}
