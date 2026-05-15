package main
import "fmt"

func lengthOfLongestSubstring(data string) int {
    left, best := 0, 0
    seen := make(map[byte]bool)
    
    for right := 0; right < len(data); right ++ {
        for seen[data[right]] {
            delete (seen, data[left])
            left ++
        }
        seen[data[right]] = true
        if right - left + 1 > best {
            best = right - left + 1
        }
    }
    return best
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3 → "abc"
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1 → "b"
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3 → "wke"
	fmt.Println(lengthOfLongestSubstring(""))         // 0
}
