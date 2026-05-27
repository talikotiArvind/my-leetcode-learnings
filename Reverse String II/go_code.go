package main
import "fmt"

func reverseStringsII(s string, k int) string {
    b := []byte(s)
    for i:= 0; i < len(b); i+=2*k {
        left, right := i, min(i+k-1,len(b)-1)
        for left < right {
            b[left], b[right] = b[right], b[left]
            left++
            right--
        }
    }
    return string(b)
}

func min(a, b int) int {
    if a < b {return a}
    return b
}

func main() {
  data := "abcdefg"
  fmt.Println(reverseStringsII(data, 2))
}
