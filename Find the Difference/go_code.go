package main
import "fmt"
func findDifference(s string, t string) string {
    result := byte(0)
    for i:= 0; i< len(s); i++ {
        result ^= s[i]
    }
    
    for i:= 0; i< len(t); i++ {
        result ^= t[i]
    }
    return string(result)
}


func main() {
    s := "abcd"
    t := "abcde"
    fmt.Println(findDifference(s,t))
}
