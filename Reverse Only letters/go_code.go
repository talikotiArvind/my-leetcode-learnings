package main
import "fmt"

func reverseOnlyLetters(s string) string {
    b := []byte(s)
    l, r := 0, len(b) - 1
    
    for l < r {
        if !isLetter(b[l]) {
            l++
        } else if !isLetter(b[r]) {
            r--
        } else {
            b[l], b[r] = b[r], b[l]
            l++
            r--
        }
    }
    return string(b)
}

func isLetter(c byte) bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func main() {
  data := "a-bC-dEf-ghIj"
  fmt.Println(reverseOnlyLetters(data))
}
