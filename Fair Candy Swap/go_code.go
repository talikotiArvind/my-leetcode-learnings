package main

import "fmt"

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sumA, sumB := 0, 0
	for _, v := range aliceSizes { sumA += v }
	for _, v := range bobSizes { sumB += v }

	delta := (sumA - sumB) / 2
	bobSet := make(map[int]bool)
	for _, v := range bobSizes { bobSet[v] = true }

	for _, a := range aliceSizes {
		if b := a - delta; bobSet[b] {
			return []int{a, b}
		}
	}
	return nil
}

func main() {
	fmt.Println(fairCandySwap([]int{1, 1}, []int{2, 2}))       // [1 2]
	fmt.Println(fairCandySwap([]int{1, 2}, []int{2, 3}))       // [1 2]
	fmt.Println(fairCandySwap([]int{2}, []int{1, 3}))          // [2 3]
	fmt.Println(fairCandySwap([]int{1, 2, 5}, []int{2, 4}))    // [5 4]
}
