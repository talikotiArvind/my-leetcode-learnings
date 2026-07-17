package main

import (
	"fmt"
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	s := strconv.Itoa(num)
	count := 0
	for i := 0; i+k <= len(s); i++ {
		w, _ := strconv.Atoi(s[i : i+k])
		if w != 0 && num%w == 0 {
			count++
		}
	}
	return count
}

func main() {
	cases := []struct {
		num, k, want int
	}{
		{240, 2, 2},     // 24 and 40 both divide 240
		{430043, 2, 2},  // 43 appears twice
		{1, 1, 1},       // single digit divides itself
		{20, 1, 1},      // "2" divides, "0" is skipped
	}

	for _, c := range cases {
		got := divisorSubstrings(c.num, c.k)
		status := "ok"
		if got != c.want {
			status = "FAIL"
		}
		fmt.Printf("num=%-7d k=%d -> %d (want %d) %s\n", c.num, c.k, got, c.want, status)
	}
}
