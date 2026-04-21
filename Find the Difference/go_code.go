package main
import "fmt"

// XOR (fastest)
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

// sum of ASCII values
func findDifferencesASCIIChar(s string, t string) string {
    sum := 0
    for _, c := range t { sum += int(c)}
    for _, c := range s { sum -= int(c)}
    return string(sum)
}

func main() {
    s := "abcd"
    t := "abcde"
    fmt.Println(findDifference(s,t))
}
