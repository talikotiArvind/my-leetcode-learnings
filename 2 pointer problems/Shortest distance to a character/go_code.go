package main
import "fmt"

func shortDistanceToChar(s string, c byte) [] int {
    n := len(s)
    ans := make([]int,n)
    
    
    prev := -n  
    for i := 0; i < n; i++ {
        if s[i] == c {
            prev = i
        }
        ans[i] = i - prev
    }
    
    prev = 2 * n
    for i := n -1; i >= 0; i-- {
        if s[i] == c{
            prev = i
        }
        ans[i] = min(ans[i], prev - i)
    }
    return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
  fmt.Println(shortDistanceToChar("loveleetcode", 'e'))
}
