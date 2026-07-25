package main

import (
	"fmt"
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	res := 0
	n := len(heaters)
	for _, h := range houses {
		i := sort.SearchInts(heaters, h)
		left := math.MaxInt64
		right := math.MaxInt64
		if i > 0 {
			left = h - heaters[i-1]
		}
		if i < n {
			right = heaters[i] - h
		}
		d := left
		if right < d {
			d = right
		}
		if d > res {
			res = d
		}
	}
	return res
}

func main() {
	fmt.Println(findRadius([]int{1, 2, 3}, []int{2}))
	fmt.Println(findRadius([]int{1, 2, 3, 4}, []int{1, 4}))
	fmt.Println(findRadius([]int{1, 5}, []int{2}))
}
