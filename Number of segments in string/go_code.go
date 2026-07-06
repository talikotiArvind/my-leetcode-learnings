package main

import "fmt"

func countSegments(s string) int {
	count := 0
	inSegment := false

	for _, ch := range s {
		if ch != ' ' {
			if !inSegment {
				count++
				inSegment = true
			}
		} else {
			inSegment = false
		}
	}

	return count
}

func main() {
	fmt.Println(countSegments("Hello, my name is John")) // 5
	fmt.Println(countSegments(""))                        // 0
	fmt.Println(countSegments("   "))                     // 0
	fmt.Println(countSegments("a b  c   d"))              // 4
	fmt.Println(countSegments("love live! mu'sic forever")) // 4
}
