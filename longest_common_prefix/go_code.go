package main

import "fmt"

func main() {
	strs := []string{"flower", "dog", "flight"}
	minLen := len(strs[0])
	for _, s := range strs {
		if len(s) < minLen {
			minLen = len(s)
		}
	}

	prefix := ""

	for i := 0; i < minLen; i++ {
		ch := strs[0][i]
		same := true
		for _, word := range strs {
			if word[i] != ch {
				same = false
				break
			}
		}
		if same {
			prefix += string(ch)
		} else {
			break
		}

	}
	fmt.Printf("Common prefix: '%s'", prefix)
}
