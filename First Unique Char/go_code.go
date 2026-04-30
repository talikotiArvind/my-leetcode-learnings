package main
import "fmt"

func firstUniqChar (s string) int {
    output := [26]int{}
    
    for _, value := range s {
        output[value - 'a']++
    }
    
    for i, char := range s {
        if output[char - 'a'] == 1 {
            return i
        }
    }
    
    return -1
}

func main() {
  fmt.Println(firstUniqChar("leetcode"))  // 0 ('l')
  fmt.Println(firstUniqChar("aabb"))      // -1
}
