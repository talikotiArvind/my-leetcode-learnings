package main
import "fmt"

func countBinaryStrings(s string) int {
    prev, cur := 0, 1
    ans := 0
    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            cur++
        } else {
            ans += min(prev, cur)
            prev = cur
            cur = 1
        }
    }
    ans += min(prev, cur)
    return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
  fmt.Println(countBinaryStrings("00110011"))
}
