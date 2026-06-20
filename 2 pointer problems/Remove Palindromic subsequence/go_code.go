package main
import "fmt"

func removePalindromicSubsequence(s string) int {
    if len(s) == 0 {return 0}
    if isPalindrome(s) {return 1}
    return 2
}

func isPalindrome(s string) bool {
    low, high := 0, len(s) - 1
    for low < high {
        if s[low] != s[high] {
            return false
        }
        low ++
        high --
    }
    return true
}

func main() {
  fmt.Println(removePalindromicSubsequence("ababa")) // 1
  fmt.Println(removePalindromicSubsequence("abb"))    // 2
  fmt.Println(removePalindromicSubsequence(""))       // 0
}
