package main
import "fmt"

func check(s string) int {
    n := len(s)
    count := 0
    
    for i := 0; i<n-2;i++ {
        if s[i] != s[i+1] && s[i+1] != s[i+2] && s[i+2] != s[i] {
            count += 1
        }
    }
    return count
}

func main() {
  fmt.Println(check("aababcabc"))
}
