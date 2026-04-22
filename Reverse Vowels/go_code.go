package main
import "fmt"

func reverseVowels(s string) string {
    vowels := map[byte]bool{
        'a': true, 'e':true, 'i':true, 'o':true, 'u':true,
        'A': true, 'E':true, 'I':true, 'O':true, 'U':true,
    }
    
    result := []byte(s)
    left, right := 0, len(s) - 1
    
    for left < right {
        for left < right && !vowels[result[left]] {left++}
        for left < right && !vowels[result[right]] {right--}
        result[left], result[right] = result[right], result[left]
        left++
        right--
    }
    return string(result)
}

func main() {
    examples := []string{"hello", "leetcode", "aA", "race a car"}

    for _, s := range examples {
        fmt.Printf("%-12s => %s\n", s, reverseVowels(s))
    }
}
