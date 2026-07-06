package main

import "fmt"

func maxNumberOfBalloons(text string) int {
	var count [26]int
	for _, ch := range text {
		count[ch-'a']++
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	res := count['b'-'a']
	res = min(res, count['a'-'a'])
	res = min(res, count['l'-'a']/2)
	res = min(res, count['o'-'a']/2)
	res = min(res, count['n'-'a'])

	return res
}

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))       // 1
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon")) // 2
	fmt.Println(maxNumberOfBalloons("leetcode"))         // 0
	fmt.Println(maxNumberOfBalloons("balon"))            // 0
}
