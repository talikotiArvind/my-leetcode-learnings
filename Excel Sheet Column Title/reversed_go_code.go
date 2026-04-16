package main

import "fmt"

func titleToNumber(columnNumber string) int {
    mapper := make(map[rune]int)
    for i := 1; i <= 26; i++ {
        mapper[rune(i+64)] = i
    }
    output := 0
    runes := []rune(columnNumber)
    for index, i := 0, len(runes)-1; i >= 0; i, index = i-1, index+1 {
        power := 1
        for j := 0; j < index; j++ {
            power *= 26
        }
        output += power * mapper[runes[i]]
    }   
    return output
}
func main() {
    fmt.Println(titleToNumber("A"))   // 1
    fmt.Println(titleToNumber("AB"))  // 28
    fmt.Println(titleToNumber("ZY"))  // 701
}
