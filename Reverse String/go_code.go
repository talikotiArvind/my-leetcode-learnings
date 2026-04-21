package main
import "fmt"

func ReverseString(s []byte) {
    left, right := 0, len(s) - 1
    for left < right {
        s[left], s[right] = s[right], s[left]
        left ++
        right --
    }
}

func main() {
    s := []byte{'h', 'e', 'l', 'l', 'o'}
    ReverseString(s)
    fmt.Println(string(s))
}
